package routes

import (
	"github.com/gorilla/mux"
	controller "github.com/mnsdojo/recruitment-system/controllers"
	"gorm.io/gorm"
)

func SetupRoutes(r *mux.Router, db *gorm.DB) {
	// Initialize the user controller with the database connection

	// Public routes
	userController := controller.NewUserController(db)
	jobsController := controller.NewJobController(db)
	r.HandleFunc("/register", userController.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/login", userController.LoginUserHandler).Methods("POST")

	r.HandleFunc("/api/jobs", jobsController.GetAllJobsHandler).Methods("GET")
	r.HandleFunc("/api/jobs/apply", jobsController).Methods("GET")

	// You can add more public or private routes as needed
}
