package controller

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/mnsdojo/recruitment-system/models"
	"github.com/mnsdojo/recruitment-system/repository"
	"github.com/mnsdojo/recruitment-system/services"
	"github.com/mnsdojo/recruitment-system/utils"
	"gorm.io/gorm"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(db *gorm.DB) *UserController {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	return &UserController{userService: userService}
}

func (uc *UserController) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate email format
	if !isValidEmail(user.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Validate user password (example: length check, more checks can be added)
	if len(user.PasswordHash) < 6 {
		http.Error(w, "Password must be at least 6 characters long", http.StatusBadRequest)
		return
	}

	createdUser, err := uc.userService.RegisterUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (uc *UserController) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate email format for login
	if !isValidEmail(creds.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Authenticate the user with the provided email and password
	user, err := uc.userService.AuthenticateUser(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate the JWT token with user ID and user type (use the values from the user object)
	token, err := utils.GenerateJWT(user.ID, string(user.UserType))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// isValidEmail checks the validity of an email using regex
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
