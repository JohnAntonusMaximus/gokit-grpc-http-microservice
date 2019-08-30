package service

import (
	"context"

	"github.com/johnantonusmaximus/gokit-grpc-poc/proto"
)

// EncodeGRPCLoremRequest encodes a LoremRequest
func EncodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(LoremRequest)

	return &pb.LoremRequest{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

// DecodeGRPCLoremRequest decodes a LoremRequest
func DecodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoremRequest)

	return LoremRequest{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

// EncodeGRPCLoremResponse encodes a LoremResponse
func EncodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(LoremResponse)

	return &pb.LoremResponse{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}

// DecodeGRPCLoremResponse decodes a LoremResponse
func DecodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.LoremResponse)

	return LoremResponse{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}
