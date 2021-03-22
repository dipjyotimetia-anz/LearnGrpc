package integration

import (
	"context"
	"fmt"
	blogpb "github.com/dipjyotimetia-anz/gogrpc/blog/blogPb"
	"github.com/google/uuid"
	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"testing"
)

var _id = uuid.New().String()

type createBlogSuite struct {
	suite.Suite
	conn      *grpc.ClientConn
	rpcClient blogpb.BlogServiceClient
}

func (s *createBlogSuite) SetupSuite() {
	var err error
	s.conn, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	s.rpcClient = blogpb.NewBlogServiceClient(s.conn)
}

func (s *createBlogSuite) TearDownSuite() {
	if err := s.conn.Close(); err != nil {
		fmt.Println(err)
	}
}

func TestBlog(t *testing.T) {
	suite.Run(t, new(createBlogSuite))
}

func (s *createBlogSuite) TestCreateBlog() {
	s.T().Run("TestCreateBlog", func(t *testing.T) {
		assert := tassert.New(t)
		req := &blogpb.CreateBlogRequest{Blog: &blogpb.Blog{
			Id:       _id,
			AuthorId: "Dip",
			Title:    "Dip gRPC Mongo",
			Content:  "mongo service",
		}}
		res, err := s.rpcClient.CreateBlog(context.Background(), req)
		if !assert.NoError(err) {
			s.T().Fail()
		}
		fmt.Println(res)
	})
}

func (s *createBlogSuite) TestReadBlog() {
	s.T().Run("TestReadBlog", func(t *testing.T) {
		assert := tassert.New(t)
		req := &blogpb.ReadBlogRequest{BlogId: "6058152ac21979a0d1a770e4"}
		res, err := s.rpcClient.ReadBlog(context.Background(), req)
		if !assert.NoError(err) {
			s.T().Fail()
		}
		fmt.Println(res)
	})
}
