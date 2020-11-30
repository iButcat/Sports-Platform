package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	FetchSportsEndpoint endpoint.Endpoint
	GetEndpoint         endpoint.Endpoint
	GetAllEndpoint      endpoint.Endpoint
	UpdateEndpoint      endpoint.Endpoint
	DeleteEndpoint      endpoint.Endpoint
}

func MakeServerEndpoints(service Service) Endpoints {
	return Endpoints{
		FetchSportsEndpoint: MakeFetchSportsEndpoint(service),
		GetEndpoint:         MakeGetEndpoint(service),
		GetAllEndpoint:      MakeGetAllEndpoint(service),
		UpdateEndpoint:      MakeUpdateEndpoint(service),
		DeleteEndpoint:      MakeDeleteEndpoint(service),
	}
}

func MakeFetchSportsEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FetchRequest)
		valid, err := service.FetchSportsAPI(ctx, req.url)
		return FetchResponse{
			V:   valid,
			Err: err,
		}, err
	}
}

func MakeGetEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		data, err := service.Get(ctx, req.Id)
		return GetResponse{
			data: data,
			Err:  err,
		}, nil
	}
}

func MakeGetAllEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(GetAllRequest) // request is not required
		allData, allSites, err := service.GetAll(ctx)
		return GetAllResponse{
			Data:  allData,
			Sites: allSites,
			Err:   err,
		}, nil
	}
}

func MakeUpdateEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateRequest)
		msg, err := service.Update(ctx, req.sites)
		return UpdateResponse{
			V:   msg,
			Err: err,
		}, nil
	}
}

func MakeDeleteEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteRequest)
		msg, err := service.Delete(ctx, req.Id)
		if err != nil {
			return DeleteResponse{
				Msg: msg,
				Err: err,
			}, nil
		}
		return DeleteResponse{
			Msg: msg,
		}, nil
	}
}

type (
	FetchRequest struct {
		url string `json:"url"`
	}

	FetchResponse struct {
		V   string `json:"valid"`
		Err error  `json:"-"`
	}

	GetRequest struct {
		Id string `json:"id"`
	}

	GetResponse struct {
		data Data
		Err  error `json:"-"`
	}

	GetAllRequest struct {
	}

	GetAllResponse struct {
		Data  []Data  `json:"data"`
		Sites []Sites `json:"sites"`
		Err   error   `json:"-"`
	}

	UpdateRequest struct {
		sites Sites
	}

	UpdateResponse struct {
		V   string `json:"valid"`
		Err error  `json:"-"`
	}

	DeleteRequest struct {
		Id string `json:"id"`
	}

	DeleteResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error"`
	}
)
