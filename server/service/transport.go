package service

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/johnantonusmaximus/gokit-grpc-poc/proto"
)

type grpcServer struct {
	lorem grpctransport.Handler
}

// Lorem implements the LoremServer interface in lorem.pb.go
func (s *grpcServer) Lorem(ctx context.Context, req *pb.LoremRequest) (*pb.LoremResponse, error) {
	_, resp, err := s.lorem.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.LoremResponse), nil
}

// NewGRPCServer creates a new GRPC server
func NewGRPCServer(endpoint Endpoints) pb.LoremServer {
	return &grpcServer{
		lorem: grpctransport.NewServer(
			endpoint.LoremEndpoint,
			DecodeGRPCLoremRequest,
			EncodeGRPCLoremResponse,
		),
	}
}
