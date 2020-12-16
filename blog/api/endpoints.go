package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreatePostEndpoints endpoint.Endpoint
	GetPostEndpoint     endpoint.Endpoint
	GetAllPostEndpoint  endpoint.Endpoint
	UpdatePostEndpoint  endpoint.Endpoint
	DeletePostEndpoint  endpoint.Endpoint
}

func MakeServerEndpoints(service Service) Endpoints {
	return Endpoints{
		CreatePostEndpoints: MakeCreatePostEndpoint(service),
		GetPostEndpoint:     MakeGetPostEndpoint(service),
		GetAllPostEndpoint:  MakeGetAllPostEndpoint(service),
		UpdatePostEndpoint:  MakeUpdatePostEndpoint(service),
		DeletePostEndpoint:  MakeDeletePostEndpoint(service),
	}
}

func MakeCreatePostEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreatePostRequest)
		ok, err := service.CreatePost(ctx, req.Title, req.Article, req.Image)
		return CreatePostResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func MakeGetPostEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetPostRequest)
		post, err := service.GetPost(ctx, req.ID)
		return GetPostResponse{
			Post: post,
			Err:  err,
		}, err
	}
}

func MakeGetAllPostEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(GetAllPostRequest)
		posts, err := service.GetAllPost(ctx)
		return GetAllPostResponse{
			Post: posts,
			Err:  err,
		}, err
	}
}

func MakeUpdatePostEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdatePostRequest)
		ok, err := service.UpdatePost(ctx, req.post)
		return UpdatePostResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func MakeDeletePostEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeletePostRequest)
		ok, err := service.DeletePost(ctx, req.ID)
		return DeletePostResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

type (
	CreatePostRequest struct {
		Title   string
		Article string
		Image   string
	}

	CreatePostResponse struct {
		Ok  string `json:"ok"`
		Err error  `json:"-"`
	}

	GetPostRequest struct {
		ID string
	}

	GetPostResponse struct {
		Post Post
		Err  error `json:"-"`
	}

	GetAllPostRequest struct {
	}

	GetAllPostResponse struct {
		Post []Post `json:"posts"`
		Err  error  `json:"-"`
	}

	UpdatePostRequest struct {
		post Post
	}

	UpdatePostResponse struct {
		Ok  string `json:"ok"`
		Err error  `json:"-"`
	}

	DeletePostRequest struct {
		ID string
	}

	DeletePostResponse struct {
		Ok  string `json:"ok"`
		Err error  `json:"-"`
	}
)
