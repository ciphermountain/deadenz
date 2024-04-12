package multiverse

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "github.com/ciphermountain/deadenz/pkg/proto/multiverse"
)

type Client struct {
	grpcClient proto.MultiverseClient
}

func NewClient(addr string) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		grpcClient: proto.NewMultiverseClient(conn),
	}, nil
}

func (c *Client) PublishEvent(ctx context.Context, in *proto.Event) error {
	_, err := c.grpcClient.PublishEvent(ctx, in)

	return err
}

func (c *Client) NewEventsStreamReader(ctx context.Context) (*EventsReader, error) {
	events, err := c.grpcClient.Events(ctx, &proto.Filter{})
	if err != nil {
		return nil, err
	}

	return &EventsReader{eventsClient: events}, nil
}

type EventsReader struct {
	eventsClient proto.Multiverse_EventsClient
}

func (r *EventsReader) Next() (*proto.Event, error) {
	return r.eventsClient.Recv()
}

func (r *EventsReader) Close() error {
	return r.eventsClient.CloseSend()
}
