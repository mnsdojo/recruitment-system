package services

import (
	"github.com/mnsdojo/recruitment-system/models"
	"github.com/mnsdojo/recruitment-system/repository"
)

type JobService interface {
	CreateJob(job *models.Job) error
	GetJobByID(id int) (*models.Job, error)
	GetAllJobs() ([]models.Job, error)
	UpdateJob(job *models.Job) error
	DeleteJob(id int) error
}

type jobService struct {
	jobRepo repository.JobRepository
}


func NewJobService(jobRepo repository.JobRepository) JobService {
	return &jobService{
		jobRepo:jobRepo,
	}
}

// CreateJob creates a new job listing
func (s *jobService) CreateJob(job *models.Job) error {
	return s.jobRepo.Create(job)
}

// GetJobByID returns a job by its ID
func (s *jobService) GetJobByID(id int) (*models.Job, error) {
	return s.jobRepo.GetByID(id)
}

// GetAllJobs returns all available jobs
func (s *jobService) GetAllJobs() ([]models.Job, error) {
	return s.jobRepo.GetAll()
}

// UpdateJob updates an existing job
func (s *jobService) UpdateJob(job *models.Job) error {
	return s.jobRepo.Update(job)
}

// DeleteJob deletes a job by its ID
func (s *jobService) DeleteJob(id int) error {
	return s.jobRepo.Delete(id)
}
