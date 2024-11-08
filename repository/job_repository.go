package repository

import (
	"github.com/mnsdojo/recruitment-system/models"
	"gorm.io/gorm"
)

type JobRepository interface {
	Create(job *models.Job) error
	GetByID(id int) (*models.Job, error)
	GetAll() ([]models.Job, error)
	Update(job *models.Job) error
	Delete(id int) error
}
type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{db}
}

// Create inserts a new job into the database
func (r *jobRepository) Create(job *models.Job) error {
	return r.db.Create(job).Error
}

// GetByID fetches a job by its ID
func (r *jobRepository) GetByID(id int) (*models.Job, error) {
	var job models.Job
	if err := r.db.First(&job, id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

// GetAll fetches all jobs from the database
func (r *jobRepository) GetAll() ([]models.Job, error) {
	var jobs []models.Job
	if err := r.db.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

// Update modifies an existing job
func (r *jobRepository) Update(job *models.Job) error {
	return r.db.Save(job).Error
}

// Delete removes a job by ID
func (r *jobRepository) Delete(id int) error {
	return r.db.Delete(&models.Job{}, id).Error
}
