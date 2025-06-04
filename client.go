package hicon

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client *Client
)

type Client struct {
	conn *grpc.ClientConn
}

func NewClient(ctx context.Context, addr string) (*Client, error) {
	if client == nil {
		newConn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}

		client = &Client{conn: newConn}
	}

	return client, nil
}

// NewUpsertConfig New creates a new NewUpsertConfig with default settings
func (s *Client) NewUpsertConfig(opts ...UpsertConfigOption) *UpsertConfig {
	cfg := &UpsertConfig{}
	cfg.build(opts...)

	return cfg
}

type ExecOptions struct {
	RequestID string
}
