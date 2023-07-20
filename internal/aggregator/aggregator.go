package aggregator

import (
	"github.com/ashkan-maleki/go-r-msg-rmq-aggregator/internal/aggregator/state"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Aggregator struct {
	Config

	conn         *amqp.Connection
	state        state.StateManager
	scenario     Scenario
	ErrorManager ErrorManager

	producerChannel *amqp.Channel
}
