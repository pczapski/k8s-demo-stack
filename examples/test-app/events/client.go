package events

import (
	"context"
	"errors"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type client struct {
	c         cloudevents.Client
	sourceURI string
	targetURI string
}

func NewClient(sourceURI, targetURI string) (*client, error) {
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &client{c: c, sourceURI: sourceURI, targetURI: targetURI}, nil
}
func (c *client) SendEvent(ctx context.Context) error {
	event := cloudevents.NewEvent()
	event.SetSource(c.sourceURI)
	event.SetType("helloworld")
	event.SetData(cloudevents.ApplicationJSON, map[string]string{"name": "world"})
	ctx = cloudevents.ContextWithTarget(ctx, c.targetURI)

	if result := c.c.Send(ctx, event); !cloudevents.IsACK(result) {
		return errors.New(result.Error())
	}
	return nil
}
