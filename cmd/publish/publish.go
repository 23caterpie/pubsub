package publish

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"

	"cloud.google.com/go/pubsub"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	runner := newRunner()
	return cli.Command{
		Name:      "publish",
		Usage:     "publish a message to pubsub",
		UsageText: `cat message.dat | pubsub publish --pubsub-project-id test-project --pubsub-topic test-topic -`,
		Flags:     runner.flags(),
		Action:    runner.run,
	}
}

func newRunner() *runner {
	return &runner{
		config: newConfig(),
	}
}

type runner struct {
	config
}

func (r *runner) flags() []cli.Flag {
	return r.config.flags()
}

func (r *runner) run(c *cli.Context) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	client, err := pubsub.NewClient(ctx, r.PubSubConfig.ProjectID)
	if err != nil {
		return fmt.Errorf("error making pubsub client: %w", err)
	}

	var messageSource io.Reader
	switch c.Args().First() {
	case "-", "--":
		messageSource = os.Stdin
	default:
		messageSource = bytes.NewBufferString(r.PubSubMessage)
	}

	message, err := io.ReadAll(messageSource)
	if err != nil {
		return fmt.Errorf("error reading message bytes: %w", err)
	}

	_, err = client.Topic(r.PubSubTopic).Publish(ctx, &pubsub.Message{
		Data: message,
	}).Get(ctx)
	if err != nil {
		return fmt.Errorf("error publishing message to pubsub: %w", err)
	}

	return nil
}
