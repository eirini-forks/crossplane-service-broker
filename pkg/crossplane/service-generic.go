package crossplane

import (
	"context"
	"encoding/json"

	"code.cloudfoundry.org/lager"
)

type GenericServiceBinder struct {
	serviceBinder
}

func NewGenericServiceBinder(c *Crossplane, instance *Instance, logger lager.Logger) *GenericServiceBinder {
	return &GenericServiceBinder{
		serviceBinder: serviceBinder{
			instance: instance,
			cp:       c,
			logger:   logger,
		},
	}
}

func (g GenericServiceBinder) Bind(ctx context.Context, bindingID string) (Credentials, error) {
	return g.GetBinding(ctx, bindingID)
}

func (g GenericServiceBinder) Unbind(ctx context.Context, bindingID string) error {
	return nil
}

func (g GenericServiceBinder) Deprovisionable(ctx context.Context) error {
	return nil
}

func (g GenericServiceBinder) GetBinding(ctx context.Context, bindingID string) (Credentials, error) {
	connectionSecret, err := g.cp.GetConnectionDetails(ctx, g.instance.Composite)
	if err != nil {
		return nil, err
	}

	credentials := Credentials{}
	for key, value := range connectionSecret.Data {
		credentials[key] = string(value)
	}

	return credentials, nil
}

func (g GenericServiceBinder) ValidateProvisionParams(ctx context.Context, params json.RawMessage) (map[string]any, error) {
	// TODO: unmarshal raw params
	return map[string]any{}, nil
}
