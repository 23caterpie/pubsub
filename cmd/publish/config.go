package publish

import (
	"github.com/urfave/cli"
)

func newConfig() config {
	return config{
		PubSubConfig: NewPubSubConfig(),
	}
}

type config struct {
	PubSubConfig  PubSubConfig
	PubSubTopic   string
	PubSubMessage string
}

func (c *config) flags() []cli.Flag {
	return append(
		[]cli.Flag{
			cli.StringFlag{
				Name:        "pubsub-topic",
				EnvVar:      "PUBSUB_TOPIC",
				Usage:       "the pubsub topic to publish the message",
				Required:    true,
				Value:       "",
				Destination: &c.PubSubTopic,
			},
			cli.StringFlag{
				Name:        "pubsub-message",
				EnvVar:      "PUBSUB_MESSAGE",
				Usage:       "the message payload to publish to pubsub; can be overridden by stdin if dash supplied as first argument in command",
				Required:    false,
				Value:       "",
				Destination: &c.PubSubMessage,
			},
		},
		c.PubSubConfig.Flags()...,
	)
}

func NewPubSubConfig() PubSubConfig {
	return PubSubConfig{}
}

type PubSubConfig struct {
	ProjectID string
}

func (c *PubSubConfig) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "pubsub-project-id",
			EnvVar:      "PUBSUB_PROJECT_ID",
			Usage:       "the pubsub project id to which to connect",
			Required:    true,
			Value:       "",
			Destination: &c.ProjectID,
		},
	}
}
