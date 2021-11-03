package machine

import (
	machinev1 "github.com/openshift/api/machine/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

const (
	machineCreationSucceedReason  = "MachineCreationSucceeded"
	machineCreationSucceedMessage = "machine successfully created"
	machineCreationFailedReason   = "MachineCreationFailed"
)

func shouldUpdateCondition(
	oldCondition machinev1.AzureMachineProviderCondition,
	newCondition machinev1.AzureMachineProviderCondition,
) bool {
	if oldCondition.Status != newCondition.Status ||
		oldCondition.Reason != newCondition.Reason ||
		oldCondition.Message != newCondition.Message {
		return true
	}
	return false
}

// setMachineProviderCondition sets the condition for the machine and
// returns the new slice of conditions.
// If the machine does not already have a condition with the specified type,
// a condition will be added to the slice.
// If the machine does already have a condition with the specified type,
// the condition will be updated if either of the following are true.
// 1) Requested Status is different than existing status.
// 2) requested Reason is different that existing one.
// 3) requested Message is different that existing one.
func setMachineProviderCondition(conditions []machinev1.AzureMachineProviderCondition, newCondition machinev1.AzureMachineProviderCondition) []machinev1.AzureMachineProviderCondition {
	now := metav1.Now()
	currentCondition := findCondition(conditions, newCondition.Type)
	if currentCondition == nil {
		klog.V(4).Infof("Adding new provider condition %v", newCondition)
		conditions = append(
			conditions,
			machinev1.AzureMachineProviderCondition{
				Type:               newCondition.Type,
				Status:             newCondition.Status,
				Reason:             newCondition.Reason,
				Message:            newCondition.Message,
				LastTransitionTime: now,
				LastProbeTime:      now,
			},
		)
	} else {
		if shouldUpdateCondition(
			*currentCondition,
			newCondition,
		) {
			klog.V(4).Infof("Updating provider condition %v", newCondition)
			if currentCondition.Status != newCondition.Status {
				currentCondition.LastTransitionTime = now
			}
			currentCondition.Status = newCondition.Status
			currentCondition.Reason = newCondition.Reason
			currentCondition.Message = newCondition.Message
			currentCondition.LastProbeTime = now
		}
	}
	return conditions
}

// findCondition finds in the machine the condition that has the
// specified condition type. If none exists, then returns nil.
func findCondition(conditions []machinev1.AzureMachineProviderCondition, conditionType machinev1.ConditionType) *machinev1.AzureMachineProviderCondition {
	for i, condition := range conditions {
		if condition.Type == conditionType {
			return &conditions[i]
		}
	}
	return nil
}
