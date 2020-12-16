package api

import (
	"context"
)

type Contact struct {
	ID        string `gorm:"PrimaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Message   string `json:"message"`
	User      string `json:"user"` // should use other service for user
}

type Repository interface {
	Create(ctx context.Context, contact Contact) (string, error)
	Get(ctx context.Context, id string) (Contact, error)
	GetAll(ctx context.Context) ([]Contact, error)
	Delete(ctx context.Context, id string) (string, error)
}
