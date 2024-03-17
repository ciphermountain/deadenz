package multiverse

import (
	"context"

	"github.com/ciphermountain/deadenz/pkg/multiverse/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MultiverseClient struct {
	grpcClient service.MultiverseClient
}

func NewMultiverseClient(addr string) (*MultiverseClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	return &MultiverseClient{
		grpcClient: service.NewMultiverseClient(conn),
	}, nil
}

func (c *MultiverseClient) PublishEvent(ctx context.Context, in *service.Event) error {
	_, err := c.grpcClient.PublishEvent(ctx, in)

	return err
}

func (c *MultiverseClient) NewEventsStreamReader(ctx context.Context) (*EventsReader, error) {
	events, err := c.grpcClient.Events(ctx, &service.Filter{})
	if err != nil {
		return nil, err
	}

	return &EventsReader{eventsClient: events}, nil
}

type EventsReader struct {
	eventsClient service.Multiverse_EventsClient
}

func (r *EventsReader) Next() (*service.Event, error) {
	return r.eventsClient.Recv()
}

func (r *EventsReader) Close() error {
	return r.eventsClient.CloseSend()
}
