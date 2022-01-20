package transport

import (
	"backend/api/endpoint"
	"backend/service"
	"context"
	"encoding/json"
	"net/http"

	// internal pkg

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(service service.Service, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	endpoints := endpoint.MakeServerEndpoints(service)

	router.Methods("GET").Path("/sports/{id}").Handler(httptransport.NewServer(
		endpoints.GetEndpoint,
		decodeGetRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/sports/name/{name}").Handler(httptransport.NewServer(
		endpoints.GetNameEndpoint,
		decodeGetNameRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/sports").Handler(httptransport.NewServer(
		endpoints.GetAllEndpoint,
		decodeGetAllRequest,
		encodeResponse,
	))

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}

func decodeGetRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.GetRequest
	vars := mux.Vars(r)
	req = endpoint.GetRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeGetNameRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.GetNameRequest
	vars := mux.Vars(r)["name"]
	req = endpoint.GetNameRequest{
		Name: vars,
	}
	return req, nil
}

func decodeGetAllRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.GetAllRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return json.NewEncoder(w).Encode(response)
}
