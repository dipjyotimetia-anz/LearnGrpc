package main

import (
	"log"
	"testing"

	"github.com/dipjyotimetia-anz/gogrpc/greet/greetpb"
	"google.golang.org/grpc"
)

func BenchmarkDoUnary(b *testing.B) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error is, %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	for n := 0; n < b.N; n++ {
		DoUnary(c)
	}
}

func BenchmarkParallelDoUnary(b *testing.B) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error is, %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			DoUnary(c)
		}
	})
}

func BenchmarkDoClientStreaming(b *testing.B) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error is, %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	for n := 0; n < b.N; n++ {
		DoClientStreaming(c)
	}
}

func BenchmarkDoServerStreaming(b *testing.B) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error is, %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	for n := 0; n < b.N; n++ {
		DoServerStreaming(c)
	}
}

func BenchmarkDoClientBidirectional(b *testing.B) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error is, %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	for n := 0; n < b.N; n++ {
		DoClientBidirectional(c)
	}
}
