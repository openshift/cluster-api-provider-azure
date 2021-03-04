package termination

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// azureTerminationEndpointURL see the following link for more details about the endpoint
	// https://docs.microsoft.com/en-us/azure/virtual-machines/windows/scheduled-events#endpoint-discovery
	azureTerminationEndpointURL                          = "http://169.254.169.254/metadata/scheduledevents?api-version=2019-08-01"
	terminatingConditionType    corev1.NodeConditionType = "Terminating"
	terminationRequestedReason                           = "TerminationRequested"
)

// Handler represents a handler that will run to check the termination
// notice endpoint and mark node for deletion if the instance termination notice is fulfilled.
type Handler interface {
	Run(stop <-chan struct{}) error
}

// NewHandler constructs a new Handler
func NewHandler(logger logr.Logger, cfg *rest.Config, pollInterval time.Duration, namespace, nodeName string) (Handler, error) {
	c, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	pollURL, err := url.Parse(azureTerminationEndpointURL)
	if err != nil {
		// This should never happen
		panic(err)
	}

	logger = logger.WithValues("node", nodeName, "namespace", namespace)

	return &handler{
		client:       c,
		pollURL:      pollURL,
		pollInterval: pollInterval,
		nodeName:     nodeName,
		namespace:    namespace,
		log:          logger,
	}, nil
}

// handler implements the logic to check the termination endpoint and delete the
// machine associated with the node
type handler struct {
	client       client.Client
	pollURL      *url.URL
	pollInterval time.Duration
	nodeName     string
	namespace    string
	log          logr.Logger
}

// Run starts the handler and runs the termination logic
func (h *handler) Run(stop <-chan struct{}) error {
	ctx, cancel := context.WithCancel(context.Background())

	errs := make(chan error, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		errs <- h.run(ctx)
	}()

	select {
	case <-stop:
		cancel()
		// Wait for run to stop
		wg.Wait()
		return nil
	case err := <-errs:
		cancel()
		return err
	}
}

func (h *handler) run(ctx context.Context) error {
	logger := h.log.WithValues("node", h.nodeName)
	logger.V(1).Info("Monitoring node termination")

	if err := wait.PollImmediateUntil(h.pollInterval, func() (bool, error) {
		req, err := http.NewRequest("GET", h.pollURL.String(), nil)
		if err != nil {
			return false, fmt.Errorf("could not create request %q: %w", h.pollURL.String(), err)
		}

		req.Header.Add("Metadata", "true")

		resp, err := http.DefaultClient.Do(req)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			return false, fmt.Errorf("could not get URL %q: %w", h.pollURL.String(), err)
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, fmt.Errorf("failed to read responce body: %w", err)
		}

		s := scheduledEvents{}
		err = json.Unmarshal(bodyBytes, &s)
		if err != nil {
			return false, fmt.Errorf("failed to unmarshal responce body: %w", err)
		}

		for _, event := range s.Events {
			if event.EventType == preemptEventType {
				// Instance marked for termination
				return true, nil
			}
		}

		// Instance not terminated yet
		logger.V(2).Info("Instance not marked for termination")
		return false, nil
	}, ctx.Done()); err != nil {
		return fmt.Errorf("error polling termination endpoint: %w", err)
	}

	// Will only get here if the termination endpoint returned FALSE
	logger.V(1).Info("Instance marked for termination, marking Node for deletion")
	if err := h.markNodeForDeletion(ctx); err != nil {
		return fmt.Errorf("error marking node: %v", err)
	}

	return nil
}

func (h *handler) markNodeForDeletion(ctx context.Context) error {
	node := &corev1.Node{}
	if err := h.client.Get(ctx, client.ObjectKey{Name: h.nodeName}, node); err != nil {
		return fmt.Errorf("error fetching node: %v", err)
	}

	addNodeTerminationCondition(node)
	if err := h.client.Status().Update(ctx, node); err != nil {
		return fmt.Errorf("error updating node status")
	}
	return nil
}

// nodeHasTerminationCondition checks whether the node already
// has a condition with the terminatingConditionType type
func nodeHasTerminationCondition(node *corev1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == terminatingConditionType {
			return true
		}
	}
	return false
}

// addNodeTerminationCondition will add a condition with a
// terminatingConditionType type to the node
func addNodeTerminationCondition(node *corev1.Node) {
	now := metav1.Now()
	terminatingCondition := corev1.NodeCondition{
		Type:               terminatingConditionType,
		Status:             corev1.ConditionTrue,
		LastHeartbeatTime:  now,
		LastTransitionTime: now,
		Reason:             terminationRequestedReason,
		Message:            "The cloud provider has marked this instance for termination",
	}

	if !nodeHasTerminationCondition(node) {
		// No need to merge, just add the new condition to the end
		node.Status.Conditions = append(node.Status.Conditions, terminatingCondition)
		return
	}

	// The node already has a terminating condition,
	// so make sure it has the correct status
	conditions := []corev1.NodeCondition{}
	for _, condition := range node.Status.Conditions {
		if condition.Type != terminatingConditionType {
			conditions = append(conditions, condition)
			continue
		}

		// Condition type is terminating
		if condition.Status == corev1.ConditionTrue {
			// Condition already marked true, do not update
			conditions = append(conditions, condition)
			continue
		}

		// The existing terminating condition had the wrong status
		conditions = append(conditions, terminatingCondition)
	}

	node.Status.Conditions = conditions
}

const preemptEventType = "Preempt"

// scheduledEvents represents metadata response, more detailed info can be found here:
// https://docs.microsoft.com/en-us/azure/virtual-machines/linux/scheduled-events#use-the-api
type scheduledEvents struct {
	Events []events `json:"Events"`
}

type events struct {
	EventType string `json:"EventType"`
}

// notFoundMachineForNode this error is returned when no machine for node is found in a list of machines
type notFoundMachineForNode struct{}

func (err notFoundMachineForNode) Error() string {
	return "machine not found for node"
}
