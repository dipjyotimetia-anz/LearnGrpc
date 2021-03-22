package cases

import (
	blogpb "github.com/dipjyotimetia-anz/gogrpc/blog/blogPb"
	"github.com/google/uuid"
)

var _id = uuid.New().String()

type Suite []Case

type Case struct {
	Name               string
	CreateBlogRequest  *blogpb.CreateBlogRequest
	CreateBlogResponse *blogpb.CreateBlogResponse
	ReadBlogRequest    *blogpb.ReadBlogRequest
	ReadBlogResponse   *blogpb.ReadBlogResponse
}

var TestCases = Suite{
	{
		Name: "CreateBlog",
		CreateBlogRequest: &blogpb.CreateBlogRequest{
			Blog: &blogpb.Blog{
				Id:       _id,
				AuthorId: "Dip",
				Title:    "Dip gRPC Mongo",
				Content:  "mongo service",
			}},
		CreateBlogResponse: &blogpb.CreateBlogResponse{
			Blog: &blogpb.Blog{
				Id:       _id,
				AuthorId: "Dip",
				Title:    "Dip gRPC Mongo",
				Content:  "mongo service",
			}},
	},
	{
		Name:            "ReadBlog",
		ReadBlogRequest: &blogpb.ReadBlogRequest{BlogId: _id},
		ReadBlogResponse: &blogpb.ReadBlogResponse{
			Blog: &blogpb.Blog{
				Id:       _id,
				AuthorId: "Dip",
				Title:    "Dip gRPC Mongo",
				Content:  "mongo service",
			}},
	},
}
