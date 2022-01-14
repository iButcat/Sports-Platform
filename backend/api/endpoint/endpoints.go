package endpoint

import (
	"backend/api/models"
	"backend/api/service"
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
		UpdateEndpoint:  MakeUpdateEndpoint(service),
		DeleteEndpoint:  MakeDeleteEndpoint(service),
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

func MakeUpdateEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateRequest)
		updated, err := service.Update(ctx, req.Sports)
		return UpdateResponse{
			Updated: updated,
			Err:     err,
		}, nil
	}
}

func MakeDeleteEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteRequest)
		deleted, err := service.Delete(ctx, req.Id)
		if err != nil {
			return DeleteResponse{
				Deleted: deleted,
				Err:     err,
			}, nil
		}
		return DeleteResponse{
			Deleted: deleted,
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

	UpdateRequest struct {
		Sports models.Sports
	}

	UpdateResponse struct {
		Updated bool  `json:"updated"`
		Err     error `json:"error,omitempty"`
	}

	DeleteRequest struct {
		Id string `json:"id"`
	}

	DeleteResponse struct {
		Deleted bool  `json:"deleted"`
		Err     error `json:"error,omitempty"`
	}
)
