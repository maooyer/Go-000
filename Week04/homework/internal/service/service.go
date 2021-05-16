package service

import (
	"github.com/google/wire"

	pb "github.com/maooyer/Go-000/Week04/homework/api/blog/V1"
	"github.com/maooyer/Go-000/Week04/homework/internal/biz"
	"github.com/maooyer/Go-000/Week04/homework/pkg/log"
)

var ProviderSet = wire.NewSet(NewBlogService)

type BlogService struct {
	pb.UnimplementedBlogServiceServer
	log     *log.Helper
	article *biz.ArticleUsecase
}
