package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/kaizenlabs/gokit-grpc-http-microservice/proto"
	"github.com/kaizenlabs/gokit-grpc-http-microservice/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Errors takes all possible error types in a struct
type Errors struct {
	grpc   error
	http   error
	system os.Signal
}

func main() {
	var (
		gRPCAddr = flag.String("grpc", ":50051", "gRPC listen address")
		//endpoint = flag.String("endpoint", "localhost:9090", "go.micro.srv.greeter address")
	)

	flag.Parse()

	// init lorem service
	svc := service.LoremService{}
	errs := make(chan Errors)

	endpoints := service.Endpoints{
		LoremEndpoint: service.MakeLoremEndpoint(svc),
	}

	listener, err := net.Listen("tcp", *gRPCAddr)
	if err != nil {
		errs <- Errors{
			grpc:   err,
			http:   nil,
			system: nil,
		}
		return
	}

	handler := service.NewGRPCServer(endpoints)
	gRPCServer := grpc.NewServer()
	pb.RegisterLoremServer(gRPCServer, handler)
	reflection.Register(gRPCServer)

	// start the grpc server
	go func() {
		err := gRPCServer.Serve(listener)
		if err != nil {
			errs <- Errors{
				grpc:   err,
				http:   nil,
				system: nil,
			}
			return
		}
	}()

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}

		err := pb.RegisterLoremHandlerFromEndpoint(ctx, mux, *gRPCAddr, opts)
		if err != nil {
			log.Fatalln("Error Registering Lorem Handler Form Endpoint: ", err)
		}
		fmt.Println("Server listening on PORT 8080...open another terminal window to call the service with gRPC, or use PostMan for HTTP.")
		err = http.ListenAndServe(":8080", mux)
		if err != nil {
			errs <- Errors{
				grpc:   nil,
				http:   err,
				system: nil,
			}
			return
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		err := <-c
		errs <- Errors{
			grpc:   nil,
			http:   nil,
			system: err,
		}
	}()

	fmt.Println(<-errs)
}
