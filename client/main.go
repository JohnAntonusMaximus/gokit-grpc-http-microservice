package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	grpcClient "github.com/johnantonusmaximus/gokit-grpc-http-microservice/client/service"
	"github.com/johnantonusmaximus/gokit-grpc-http-microservice/server/service"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr = flag.String("addr", ":50051", "gRPC address")
	)

	flag.Parse()
	ctx := context.Background()

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial: ", err)
	}

	defer conn.Close()

	loremService := grpcClient.New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)

	switch cmd {
	case "lorem":
		var requestType, minStr, maxStr string

		requestType, args = pop(args)
		minStr, args = pop(args)
		maxStr, args = pop(args)

		min, _ := strconv.Atoi(minStr)
		max, _ := strconv.Atoi(maxStr)
		lorem(ctx, loremService, requestType, min, max)
	default:
		log.Fatalln("Unknown command: ", cmd)
	}
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

// call the lorem service
func lorem(ctx context.Context, svc service.Service, requestType string, min, max int) {
	msg, err := svc.Lorem(ctx, requestType, min, max)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Message From Server: ", msg)
}
