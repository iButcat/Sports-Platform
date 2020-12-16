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

	router.Methods("POST").Path("/create").Handler(httptransport.NewServer(
		endpoints.CreatePostEndpoints,
		decodeCreateRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/get/{id}").Handler(httptransport.NewServer(
		endpoints.GetPostEndpoint,
		decodeGetRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/getall").Handler(httptransport.NewServer(
		endpoints.GetAllPostEndpoint,
		decodeGetAllRequest,
		encodeResponse,
	))

	router.Methods("PUT").Path("/update").Handler(httptransport.NewServer(
		endpoints.UpdatePostEndpoint,
		decodeUpdateRequest,
		encodeResponse,
	))

	router.Methods("DELETE").Path("/delete/{id}").Handler(httptransport.NewServer(
		endpoints.DeletePostEndpoint,
		decodeDeleteRequest,
		encodeResponse,
	))

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := GetPostRequest{
		ID: id,
	}
	return req, nil
}

func decodeGetAllRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req GetAllPostRequest
	return req, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req.post); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req DeletePostRequest
	vars := mux.Vars(r)
	id := vars["id"]
	req = DeletePostRequest{
		ID: id,
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
