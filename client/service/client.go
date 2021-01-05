package service

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	pb "github.com/kaizenlabs/gokit-grpc-http-microservice/proto"
	"github.com/kaizenlabs/gokit-grpc-http-microservice/server/service"
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
