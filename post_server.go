package main

import (
	"context"
	"errors"
	pb "github.com/Anarr/gomicrodev/proto/post"
	"github.com/micro/go-micro"
	"log"
)

var lastID int64
var posts []*pb.Post

const (
	ErrPostCreate = "can not create post"
	ErrEmptyPosts = "there is no available posts"
)

type PostService struct{}

func (ps *PostService) All(ctx context.Context, req *pb.PostsRequest, res *pb.PostsResponse) error {
	res.Posts = posts
	return nil
}

func (ps *PostService) Create(ctx context.Context, req *pb.PostCreateRequest, res *pb.Post) error {

	lastID++
	res.Id = lastID
	res.UserId = req.UserId
	res.Description = req.Description
	res.CreatedAt = req.CreatedAt
	posts = append(posts, res)

	return nil
}

func (ps *PostService) Delete(ctx context.Context, req *pb.PostDeleteRequest, res *pb.PostDeleteResponse) error {

	if len(posts) == 0 {
		res.Status = false
		return errors.New(ErrEmptyPosts)
	}

	posts = append(posts[:req.PostId], posts[req.PostId+1:]...)

	res.Status = true
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("post"),
	)

	pb.RegisterPostServiceHandler(service.Server(), new(PostService))

	if err := service.Run(); err != nil {
		log.Fatal("error occurs during running post service", err)
	}
}
