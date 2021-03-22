package main

import (
	"log"
	"testing"

	"github.com/dipjyotimetia-anz/gogrpc/calculator/calcpb"
	"google.golang.org/grpc"
)

func BenchmarkClientSum(b *testing.B) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error is %v", err)
	}

	defer cc.Close()

	c := calcpb.NewSumServiceClient(cc)
	for n := 0; n < b.N; n++ {
		DoUnary(c)
	}
}
