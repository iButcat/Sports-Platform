package api

import "context"

type Service interface {
	CreateContact(ctx context.Context, contact Contact) (string, error)
	GetContact(ctx context.Context, id string) (Contact, error)
	GetAllContact(ctx context.Context) ([]Contact, error)
	DeleteContact(ctx context.Context, id string) (string, error)
}
