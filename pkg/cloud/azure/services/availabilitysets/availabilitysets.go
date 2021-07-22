package availabilitysets

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-30/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
)

// Spec input specification for Get/CreateOrUpdate/Delete calls
type Spec struct {
	Name string
}

func (s *Service) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	availabilitysetsSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("invalid availability set specification")
	}

	asParams := compute.AvailabilitySet{
		Name: to.StringPtr(availabilitysetsSpec.Name),
		Sku: &compute.Sku{
			Name: to.StringPtr(string(compute.Aligned)),
		},
		Location: to.StringPtr(s.Scope.Location()),
		// Todo: figure out if we need to set the tags
		// Tags: ,
		AvailabilitySetProperties: &compute.AvailabilitySetProperties{
			PlatformFaultDomainCount:  to.Int32Ptr(int32(2)),
			PlatformUpdateDomainCount: to.Int32Ptr(int32(5)),
		},
	}

	_, err := s.Client.CreateOrUpdate(ctx, s.Scope.MachineConfig.ResourceGroup, availabilitysetsSpec.Name, asParams)
	if err != nil {
		return fmt.Errorf("failed to create availability set %s: %w", availabilitysetsSpec.Name, err)
	}

	return nil
}

// Get no-op.
func (s *Service) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	// Not implemented since there is nothing to get
	return nil, nil
}

// Delete no-op.
func (s *Service) Delete(ctx context.Context, spec azure.Spec) error {
	// Not implemented since there is nothing to delete
	return nil
}
