package repository

import (
	"database/sql"

	"github.com/diorshelton/job-tracker/internal/models"
)

type ApplicationRepository struct {
	db *sql.DB
}

func NewApplicationRepository(db *sql.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

func (r *ApplicationRepository) CreateApplication(app *models.Application) error {

	application := &models.Application {


	}
	return nil
}