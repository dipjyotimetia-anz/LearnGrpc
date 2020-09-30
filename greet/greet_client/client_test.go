package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/dipjyotimetia/gogrpc/greet/greet_client/cases"
	"github.com/dipjyotimetia/gogrpc/greet/greetpb"
	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type getGreetSuite struct {
	suite.Suite
	conn      *grpc.ClientConn
	rpcClient greetpb.GreetServiceClient
}

func (s *getGreetSuite) SetupSuite() {
	var err error
	s.conn, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	s.rpcClient = greetpb.NewGreetServiceClient(s.conn)
}

func (s *getGreetSuite) TearDownSuite() {
	if err := s.conn.Close(); err != nil {
		fmt.Println(err)
	}
}

func TestGreet(t *testing.T) {
	suite.Run(t, new(getGreetSuite))
}

func (s *getGreetSuite) TestSingleDoUnary() {
	s.T().Run("TestGreetRequestSingle", func(t *testing.T) {
		assert := tassert.New(t)
		req := &greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Dipjyoti",
				LastName:  "Metia",
			}}

		res, err := s.rpcClient.Greet(context.Background(), req)
		if !assert.NoError(err) {
			s.T().Fail()
		}
		assert.EqualValues("Dipjyoti", res.Result)
	})
}

func (s *getGreetSuite) TestSuiteDoUnary() {
	tCases := cases.TestCases
	for _, c := range tCases {
		tCase := c
		s.T().Run(tCase.Name, func(t *testing.T) {
			assert := tassert.New(t)
			res, err := s.rpcClient.Greet(context.Background(), tCase.Request)
			if !assert.NoError(err) {
				s.T().Fail()
			}
			assert.EqualValues(tCase.Response.Result, res.Result)
		})
	}
}
