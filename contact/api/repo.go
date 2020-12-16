package api

import (
	"context"

	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
)

type repo struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepo(db *gorm.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) Create(ctx context.Context, contact Contact) (string, error) {
	repo.db.AutoMigrate(&contact)
	if err := repo.db.Create(&contact).Error; err != nil {
		return "", err
	}
	return "Created", nil
}

func (repo *repo) Get(ctx context.Context, id string) (Contact, error) {
	var contact Contact
	sqlGet := "SELECT * FROM contacts WHERE id=id"
	if err := repo.db.Raw(sqlGet, id).Scan(&contact).Error; err != nil {
		return contact, err
	}
	return contact, nil
}

func (repo *repo) GetAll(ctx context.Context) ([]Contact, error) {
	var contacts []Contact
	sqlGetAll := "SELECT * FROM contact"
	if err := repo.db.Raw(sqlGetAll).Scan(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (repo *repo) Delete(ctx context.Context, id string) (string, error) {
	sqlDelete := "DELETE FROM contacts WHERE id=id"
	if err := repo.db.Exec(sqlDelete, id).Error; err != nil {
		return "", err
	}
	return "Deleted", nil
}
