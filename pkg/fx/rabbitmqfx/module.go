package rabbitmqfx

import (
	"github.com/knadh/koanf/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Config *koanf.Koanf
}

// ProvideConnection provides a connection to RabbitMQ.
func ProvideMQ(p Params) (*amqp.Connection, error) {

	connection, err := amqp.Dial(p.Config.String(p.Config.String("MQ")))
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// ConfigModule provided to fx
var MQModule = fx.Options(
	fx.Provide(ProvideMQ),
)
