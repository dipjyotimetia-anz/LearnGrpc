package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dipjyotimetia/gogrpc/calculator/calcpb"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked %v\n", req)
	firstArgument := req.GetFirstNumber()
	secondArgument := req.GetSecondNumber()

	result := firstArgument + secondArgument
	res := &calcpb.SumResponse{
		Result: result,
	}
	return res, nil
}

var (
	grpcServerEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)

func run() error {
	fmt.Println("Hello calc")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := calcpb.RegisterSumServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		fmt.Println(err)
	}
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
