package domain

import (
	"time"

	"github.com/amorindev/go-tmpl/pkg/app/auth-methods/domain"
)

// User represents a user in the system
type User struct {
	ID            interface{}              `bson:"_id"`
	Email         string                   `bson:"email"`
	EmailVerified bool                     `bson:"email_verified"`
	IsActive      bool                     `bson:"is_active"`
	UserPassAuth  *domain.UserPasswordAuth `bson:"pass_method"`
	CreatedAt     *time.Time               `bson:"created_at"`
	UpdatedAt     *time.Time               `bson:"updated_at"`
}

func NewUser(email string, password string) *User{
	return &User{
		Email: email,
		UserPassAuth: &domain.UserPasswordAuth{
			Password: password,
		},
	}
}