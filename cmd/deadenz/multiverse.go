package main

/*
import (
	"context"

	"github.com/ciphermountain/deadenz/pkg/multiverse/service"
	"google.golang.org/grpc"
)

func setupClient(addr string) error {
	var opts []grpc.DialOption

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := service.NewMultiverseClient(conn)

	client.PublishEvent(context.Background(), &service.Event{})

	events, err := client.Events(context.Background(), &service.Filter{})
	if err != nil {
		return err
	}

	for {
		event, err := events.Recv()
		if err != nil {
			return err
		}
	}
}
*/
