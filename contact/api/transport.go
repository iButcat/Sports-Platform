package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(service Service, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	endpoints := MakeServerEndpoints(service)

	router.Methods("POST").Path("/contact/create").Handler(httptransport.NewServer(
		endpoints.CreateContactEndpoint,
		decodeCreateRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/contact/get/{id}").Handler(httptransport.NewServer(
		endpoints.GetContactEndpoint,
		decodeGetRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/contact/getall").Handler(httptransport.NewServer(
		endpoints.GetAllContactEndpoint,
		decodeGetAllRequest,
		encodeResponse,
	))

	router.Methods("DELETE").Path("/contact/delete/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteContactEndpoint,
		decodeDeleteRequest,
		encodeResponse,
	))

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Contact); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := GetContactRequest{
		ID: id,
	}
	return req, nil
}

func decodeGetAllRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req GetAllContactRequest
	return req, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id := vars["id"]
	return DeleteContactRequest{
		ID: id,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
