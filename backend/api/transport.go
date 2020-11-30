package api

import (
	"context"
	"encoding/json"
	"net/http"

	"backend/config" // internal pkg

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(service Service, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	endpoints := MakeServerEndpoints(service)

	router.Methods("GET").Path("/sports/{id}").Handler(httptransport.NewServer(
		endpoints.GetEndpoint,
		decodeGetRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/sports").Handler(httptransport.NewServer(
		endpoints.GetAllEndpoint,
		decodeGetAllRequest,
		encodeResponse,
	))

	router.Methods("PUT").Path("/sports/update").Handler(httptransport.NewServer(
		endpoints.UpdateEndpoint,
		decodeUpdateRequest,
		encodeResponse,
	))

	router.Methods("DELETE").Path("/sports/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteEndpoint,
		decodeDeleteRequest,
		encodeResponse,
	))

	router.Methods("POST").Path("/fetch").Handler(httptransport.NewServer(
		endpoints.FetchSportsEndpoint,
		decodeFetchApiRequest,
		encodeResponse,
	))

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}

func decodeFetchApiRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	config, err := config.LoadConfig("./config")
	if err != nil {
		return nil, err
	}
	req := FetchRequest{
		url: config.URL,
	}
	return req, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req GetRequest
	vars := mux.Vars(r)
	req = GetRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeGetAllRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req GetAllRequest
	return req, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req = UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req.sites); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req = DeleteRequest{}
	vars := mux.Vars(r)
	req = DeleteRequest{
		Id: vars["id"],
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return json.NewEncoder(w).Encode(response)
}
