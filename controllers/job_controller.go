package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mnsdojo/recruitment-system/models"
	"github.com/mnsdojo/recruitment-system/repository"
	"github.com/mnsdojo/recruitment-system/services"
	"gorm.io/gorm"
)

type JobController struct {
	jobService services.JobService
}

func NewJobController(db *gorm.DB) *JobController {
	jobsRepo := repository.NewJobRepository(db)
	jobsService := services.NewJobService(jobsRepo)
	return &JobController{jobService: jobsService}
}

// CreateJobHandler handles the creation of a new job posting
func (c *JobController) CreateJobHandler(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.jobService.CreateJob(&job); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}

// GetJobHandler handles fetching a job by its ID
func (c *JobController) GetJobHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	job, err := c.jobService.GetJobByID(id)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(job)
}

// GetAllJobsHandler handles fetching all available jobs
func (c *JobController) GetAllJobsHandler(w http.ResponseWriter, r *http.Request) {
	jobs, err := c.jobService.GetAllJobs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(jobs)
}

// UpdateJobHandler handles updating an existing job
func (c *JobController) UpdateJobHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	job, err := c.jobService.GetJobByID(id)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.jobService.UpdateJob(job); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(job)
}

// DeleteJobHandler handles deleting a job by its ID
func (c *JobController) DeleteJobHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.jobService.DeleteJob(id); err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
