package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/johnantonusmaximus/gokit-grpc-poc/proto"
	"github.com/johnantonusmaximus/gokit-grpc-poc/server/service"
	"google.golang.org/grpc"
)

func main() {
	var (
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)

	flag.Parse()

	// init lorem service
	svc := service.LoremService{}
	errs := make(chan error)

	endpoints := service.Endpoints{
		LoremEndpoint: service.MakeLoremEndpoint(svc),
	}

	// start the grpc server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errs <- err
			return
		}

		handler := service.NewGRPCServer(endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterLoremServer(gRPCServer, handler)
		errs <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errs)
}
