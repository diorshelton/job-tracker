package repository

import (
	"database/sql"

	"github.com/diorshelton/job-tracker/internal/models"
)

type CompanyOfInterestRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyOfInterestRepository {
	return &CompanyOfInterestRepository{db: db}
}

func (r *CompanyOfInterestRepository) NewCompany(coi *models.CompanyOfInterest) error {
	return nil
}