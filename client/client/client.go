package client

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/johnantonusmaximus/gokit-grpc-poc/proto"
	"github.com/johnantonusmaximus/gokit-grpc-poc/server/service"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) service.Service {
	var loremEndpoint = grpctransport.NewClient(
		conn, "pb.Lorem", "Lorem",
		service.EncodeGRPCLoremRequest,
		service.DecodeGRPCLoremResponse,
		pb.LoremResponse{},
	).Endpoint()

	return service.Endpoints{
		LoremEndpoint: loremEndpoint,
	}
}
