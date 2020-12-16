package test

import (
	"testing"

	"blog/api" // internal package

	"github.com/stretchr/testify/assert"
)

func TestCreateService(t *testing.T) {
	service, ctx := setup()
	title, article, image := "Test Title", "Test Article", "Test image"
	createService, err := service.CreatePost(ctx, title, article, image)
	if err != nil {
		return
	}
	assert.Equal(t, "Saved", createService)
}

func TestGetService(t *testing.T) {
	service, ctx := setup()
	id := "5f819098-38d8-4032-b6be-7fa85faa5d21"
	getService, err := service.GetPost(ctx, id)
	if err != nil {
		return
	}
	post := api.Post{
		ID:      id,
		Title:   "Hello World2",
		Article: "this is my first article",
		Image:   "special image path",
	}
	assert.Equal(t, post, getService)
}

func TestGetNotFound(t *testing.T) {
	id := ""
	service, ctx := setup()
	getNotFound, err := service.GetPost(ctx, id)
	if err != nil {
		return
	}
	// empty object
	post := api.Post{}
	assert.Equal(t, post, getNotFound)
}

func TestUpdateService(t *testing.T) {
	service, ctx := setup()
	post := api.Post{
		ID:      "c9f17bfe-7f81-46f3-8bbd-53b58233627f",
		Title:   "Update",
		Article: "Update article",
		Image:   "update image",
	}
	updateService, err := service.UpdatePost(ctx, post)
	if err != nil {
		return
	}
	assert.Equal(t, "Updated", updateService)
}

func TestDeleteService(t *testing.T) {
	service, ctx := setup()
	id := "bd550d0c-98b3-41c4-a58d-28d316795709"
	deleteService, err := service.DeletePost(ctx, id)
	if err != nil {
		return
	}
	assert.Equal(t, "Deleted", deleteService)
}
