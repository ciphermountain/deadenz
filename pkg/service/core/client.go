package core

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/opts"
	proto "github.com/ciphermountain/deadenz/pkg/proto/core"
)

type Client struct {
	conn       *grpc.ClientConn
	grpcClient proto.DeadenzClient

	closer sync.Once
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
		conn:       conn,
		grpcClient: proto.NewDeadenzClient(conn),
	}, nil
}

func (c *Client) Spawnin(ctx context.Context, profile *components.Profile, options ...opts.Option) ([]string, *components.Profile, error) {
	req := &proto.RunRequest{
		Command: &proto.RunRequest_Spawnin{
			Spawnin: &proto.SpawninCommand{},
		},
		Profile: profileToProto(profile),
	}

	setter := &runRequestLanguageSetter{req: req}
	for _, opt := range options {
		opt(setter)
	}

	return c.run(ctx, profile, req)
}

func (c *Client) Walk(ctx context.Context, profile *components.Profile, options ...opts.Option) ([]string, *components.Profile, error) {
	req := &proto.RunRequest{
		Command: &proto.RunRequest_Walk{
			Walk: &proto.WalkCommand{},
		},
		Profile: profileToProto(profile),
	}

	setter := &runRequestLanguageSetter{req: req}
	for _, opt := range options {
		opt(setter)
	}

	return c.run(ctx, profile, req)
}

func (c *Client) Items(ctx context.Context, options ...opts.Option) ([]components.Item, error) {
	req := &proto.AssetRequest{
		Type: proto.AssetType_ItemAsset,
	}

	setter := &assetRequestLanguageSetter{req: req}
	for _, opt := range options {
		opt(setter)
	}

	resp, err := c.grpcClient.Assets(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Response.Status != proto.Status_OK {
		return nil, fmt.Errorf("service returned an unsuccessful response: %s", resp.Response.Message)
	}

	switch asset := resp.Asset.(type) {
	case *proto.AssetResponse_Item:
		return mutateListValues(asset.Item.GetItems(), protoToItem), nil
	default:
		return nil, fmt.Errorf("unexpected response")
	}
}

func (c *Client) Characters(ctx context.Context, options ...opts.Option) ([]components.Character, error) {
	req := &proto.AssetRequest{
		Type: proto.AssetType_CharacterAsset,
	}

	setter := &assetRequestLanguageSetter{req: req}
	for _, opt := range options {
		opt(setter)
	}

	resp, err := c.grpcClient.Assets(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Response.Status != proto.Status_OK {
		return nil, fmt.Errorf("service returned an unsuccessful response: %s", resp.Response.Message)
	}

	switch asset := resp.Asset.(type) {
	case *proto.AssetResponse_Character:
		return mutateListValues(asset.Character.Characters, protoToCharacter), nil
	default:
		return nil, fmt.Errorf("unexpected response")
	}
}

func (c *Client) Close() error {
	var err error

	c.closer.Do(func() {
		err = c.conn.Close()
	})

	return err
}

func (c *Client) run(
	ctx context.Context,
	profile *components.Profile,
	req *proto.RunRequest,
) ([]string, *components.Profile, error) {
	resp, err := c.grpcClient.Run(ctx, req)
	if err != nil {
		return nil, profile, err
	}

	if resp.Response.Status != proto.Status_OK {
		return nil, profile, fmt.Errorf("%s", resp.Response.Message)
	}

	protoProfile := protoToProfile(resp.Profile)

	return resp.Events, &protoProfile, nil
}

type runRequestLanguageSetter struct {
	req *proto.RunRequest
}

func (s *runRequestLanguageSetter) SetLanguage(lang string) {
	s.req.Language = lang
}

type assetRequestLanguageSetter struct {
	req *proto.AssetRequest
}

func (s *assetRequestLanguageSetter) SetLanguage(lang string) {
	s.req.Language = lang
}
