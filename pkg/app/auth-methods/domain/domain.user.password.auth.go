package domain

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserPasswordAuth stores password authentication data.
type UserPasswordAuth struct {
	Password     string     `bson:"-"`
	PasswordHash string     `bson:"password_hash"`
	CreatedAt    *time.Time `bson:"created_at"`
	UpdatedAt    *time.Time `bson:"updated_at"`
}

// HashPassword hashes the plain password and stores it in PasswordHash.
func (upa *UserPasswordAuth) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(upa.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}
	upa.PasswordHash = string(passwordHash)
	return nil
}
