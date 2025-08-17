package core

import (
	"errors"
	"strings"
	"unicode"
    "github.com/go-playground/validator/v10"

)

// SignUpReq represents the request structure for user registration
type SignUpReq struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// IsSignUpValid performs comprehensive validation on the sign-up request
// Returns an error if any validation fails
func (req *SignUpReq) IsSignUpValid() error {

	// Validate email field is not empty
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
    
	// Validate email format using validator
	validate := validator.New()
	err := validate.Var(req.Email, "email")
	if err != nil {
		return errors.New("invalid email format")
	}

	// Validate password field is not empty
	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	// Validate confirm password field is not empty
	if strings.TrimSpace(req.ConfirmPassword) == "" {
		return errors.New("confirm password is required")
	}

	// Validate password minimum length
	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Validate password strength (at least one uppercase, one lowercase, one number)
	if !isPasswordStrong(req.Password) {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, and one number")
	}

	// Validate passwords match
	if req.Password != req.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	// Additional validation using validator library for password confirmation
	err = validate.VarWithValue(req.Password, req.ConfirmPassword, "eqfield")
	if err != nil {
		return errors.New("passwords do not match")
	}

	return nil
}

// isPasswordStrong checks if the password meets strength requirements
// Password must contain at least one uppercase letter, one lowercase letter, and one number
func isPasswordStrong(password string) bool {
	var (
		hasUpper  bool
		hasLower  bool
		hasNumber bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	// Require at least uppercase, lowercase, and number
	return hasUpper && hasLower && hasNumber
}
