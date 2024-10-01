package multiverse

import (
	"context"
	"errors"

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

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		grpcClient: proto.NewMultiverseClient(conn),
	}, nil
}

func (c *Client) PublishGameEvent(ctx context.Context, id string, data []byte) error {
	resp, err := c.grpcClient.PublishGameEvent(ctx, &proto.GameEvent{
		Uid:  id,
		Data: data,
	})

	if resp.Status == proto.Status_Failure {
		return errors.New(resp.Message)
	}

	return err
}

func (c *Client) NewEventsStreamReader(ctx context.Context, id string) (*EventsReader, error) {
	events, err := c.grpcClient.Events(ctx, &proto.Filter{Uid: id, Recipients: []string{id}})
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
