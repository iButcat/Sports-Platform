package endpoint

import (
	"backend/models"
	"backend/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetEndpoint     endpoint.Endpoint
	GetNameEndpoint endpoint.Endpoint
	GetAllEndpoint  endpoint.Endpoint
	UpdateEndpoint  endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
}

func MakeServerEndpoints(service service.Service) Endpoints {
	return Endpoints{
		GetEndpoint:     MakeGetEndpoint(service),
		GetNameEndpoint: MakeGetNameEndpoint(service),
		GetAllEndpoint:  MakeGetAllEndpoint(service),
	}
}

func MakeGetEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		data, err := service.Get(ctx, req.Id)
		return GetResponse{
			Data: data,
			Err:  err,
		}, nil
	}
}

func MakeGetNameEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetNameRequest)
		data, err := service.GetName(ctx, req.Name)
		return GetNameResponse{
			Data: data,
			Err:  err,
		}, nil
	}
}

func MakeGetAllEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(GetAllRequest) // request is not required
		sports, err := service.GetAll(ctx)
		return GetAllResponse{
			Sports: sports,
			Err:    err,
		}, nil
	}
}

type (
	GetRequest struct {
		Id string `json:"id"`
	}

	GetResponse struct {
		Data models.Data
		Err  error `json:"error,omitempty"`
	}

	GetNameRequest struct {
		Name string
	}

	GetNameResponse struct {
		Data models.Data `json:"data"`
		Err  error       `json:"err,omitempty"`
	}

	GetAllRequest struct {
	}

	GetAllResponse struct {
		Sports *models.Sports `json:"sports"`
		Err    error          `json:"error,omitempty"`
	}
)
