package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateContactEndpoint endpoint.Endpoint
	GetContactEndpoint    endpoint.Endpoint
	GetAllContactEndpoint endpoint.Endpoint
	DeleteContactEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(service Service) Endpoints {
	return Endpoints{
		CreateContactEndpoint: MakeCreateContactEndoint(service),
		GetContactEndpoint:    MakeGetContactEndpoint(service),
		GetAllContactEndpoint: MakeGetAllContactEndpoint(service),
		DeleteContactEndpoint: MakeDeleteContactEndpoint(service),
	}
}

func MakeCreateContactEndoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateContactRequest)
		ok, err := service.CreateContact(ctx, req.Contact)
		return CreateContactResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func MakeGetContactEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetContactRequest)
		contact, err := service.GetContact(ctx, req.ID)
		return GetContactResponse{
			Contact: contact,
			Err:     err,
		}, err
	}
}

func MakeGetAllContactEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		contacts, err := service.GetAllContact(ctx)
		return GetAllContactResponse{
			Contact: contacts,
			Err:     err,
		}, err
	}
}

func MakeDeleteContactEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteContactRequest)
		ok, err := service.DeleteContact(ctx, req.ID)
		return DeleteContactResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

type (
	CreateContactRequest struct {
		Contact Contact
	}

	CreateContactResponse struct {
		Ok  string `json:"ok"`
		Err error  `json:"-"`
	}

	GetContactRequest struct {
		ID string
	}

	GetContactResponse struct {
		Contact Contact `json:"contact"`
		Err     error   `json:"-"`
	}

	GetAllContactRequest struct {
	}

	GetAllContactResponse struct {
		Contact []Contact `json:"contacts"`
		Err     error     `json:"-"`
	}

	DeleteContactRequest struct {
		ID string
	}

	DeleteContactResponse struct {
		Ok  string `json:"ok"`
		Err error  `json:"-"`
	}
)
