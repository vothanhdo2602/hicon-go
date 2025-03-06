package grpctil

import (
	"github.com/vothanhdo2602/hicon/external/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn *grpc.ClientConn
)

func NewClient() (*grpc.ClientConn, error) {
	return grpc.NewClient(config.GetAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
}
