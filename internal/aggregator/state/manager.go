package state

import "context"

var CorrelationConfigKey = "_config"

type StateManager interface {
	Append(ctx context.Context, correlationId string, strategyConfig []byte, queueName string, message []byte) error
	Delete(ctx context.Context, correlationId string) error
	All(ctx context.Context, correlationId string) (map[string][]byte, error)
}
