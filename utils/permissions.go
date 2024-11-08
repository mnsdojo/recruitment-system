package utils

import (
	"context"
	"net/http"
	"strings"
)

// AuthMiddleware validates JWT token and sets user context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header from the request
		authHeader := r.Header.Get("Authorization")

		// Check if the header is missing or doesn't start with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			JSONResponse(w, false, "Authorization header missing or invalid", nil)
			return
		}

		// Extract the token by trimming the "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT token
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			JSONResponse(w, false, "Invalid or expired token", nil)
			return
		}

		// Store claims in request context for downstream handlers
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", claims.UserId)
		ctx = context.WithValue(ctx, "userType", claims.UserType)
		r = r.WithContext(ctx)

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// AdminOnly middleware restricts access to admin users
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the user type from the request context
		userType := r.Context().Value("userType").(string)

		// Check if the user type is "Admin"
		if userType != "Admin" {
			JSONResponse(w, false, "Admin access required", nil)
			return
		}

		// If the user is an Admin, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// ApplicantOnly middleware restricts access to applicant users
func ApplicantOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the user type from the request context
		userType := r.Context().Value("userType").(string)

		// Check if the user type is "Applicant"
		if userType != "Applicant" {
			JSONResponse(w, false, "Applicant access required", nil)
			return
		}

		// If the user is an Applicant, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
