package test

import (
	"testing"

	"contact/api" // internal pkg

	"github.com/stretchr/testify/assert"
)

func TestCreateService(t *testing.T) {
	service, ctx := setup()
	contact := api.Contact{
		ID:        "10000",
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test.com",
		Message:   "test message",
	}
	createService, err := service.CreateContact(ctx, contact)
	if err != nil {
		return
	}
	assert.Equal(t, "Created", createService)
}

func TestGetService(t *testing.T) {
	service, ctx := setup()
	contact := api.Contact{
		ID:        "10000",
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test.com",
		Message:   "test message",
	}
	getService, err := service.GetContact(ctx, contact.ID)
	if err != nil {
		return
	}
	assert.Equal(t, contact, getService)
}

func TestDeleteServie(t *testing.T) {
	service, ctx := setup()
	contact := api.Contact{
		ID:        "10000",
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test.com",
		Message:   "test message",
	}
	deleteService, err := service.DeleteContact(ctx, contact.ID)
	if err != nil {
		return
	}
	assert.Equal(t, "Deleted", deleteService)
}
