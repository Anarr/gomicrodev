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
	ErrEmptyPosts = "there is no available posts"
	ErrInvalidPostDeleteRequest = "request data is invalid"
)

type PostService struct{}

//All return posts list
func (ps *PostService) All(ctx context.Context, req *pb.PostsRequest, res *pb.PostsResponse) error {
	res.Posts = posts
	return nil
}

//Create create new post
func (ps *PostService) Create(ctx context.Context, req *pb.PostCreateRequest, res *pb.Post) error {
	lastID++
	res.Id = lastID
	res.UserId = req.UserId
	res.Description = req.Description
	res.CreatedAt = req.CreatedAt
	posts = append(posts, res)

	return nil
}
//Delete delete post by PostId and UserId
func (ps *PostService) Delete(ctx context.Context, req *pb.PostDeleteRequest, res *pb.PostDeleteResponse) error {

	deletedIndex := -1

	if len(posts) == 0 {
		res.Status = false
		return errors.New(ErrEmptyPosts)
	}

	//iterate posts and find deleted index
	for index, value := range posts {
		if value.Id == req.PostId && value.UserId == req.UserId {
			deletedIndex = index
		}
	}

	if deletedIndex > -1 {
		posts = append(posts[:deletedIndex], posts[deletedIndex+1:]...)
		res.Status = true
		return nil
	}

	return errors.New(ErrInvalidPostDeleteRequest)
}

func main() {
	service := micro.NewService(
		micro.Name("post"),
	)

	pb.RegisterPostServiceHandler(service.Server(), new(PostService))

	if err := service.Run(); err != nil {
		log.Println("Post server error", err)
	}
}
