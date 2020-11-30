package api

import (
	"context"
)

type Service interface {
	FetchSportsAPI(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, id string) (Data, error)
	GetAll(ctx context.Context) ([]Data, []Sites, error)
	Update(ctx context.Context, sites Sites) (string, error)
	Delete(ctx context.Context, id string) (string, error)
}
