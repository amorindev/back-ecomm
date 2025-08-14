package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/app/auth-methods/core"
	userCore "github.com/amorindev/go-tmpl/pkg/app/users/core"
	"github.com/amorindev/go-tmpl/pkg/app/users/domain"
	coreShared "github.com/amorindev/go-tmpl/pkg/shared/core"
)

// Signup handles user registration, validates input, creates a new user, and returns a JSON response.
func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.SignUpReq

	// Decode JSON request body into SignUpReq struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(coreShared.ErrorMsg{Msg: "Invalid request format"})
		return
	}

	defer r.Body.Close()

	// Validate the sign-up request
	err = req.IsSignUpValid()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(coreShared.ErrorMsg{Msg: err.Error()})
		return
	}

	// Create a new user domain object
	user := domain.NewUser(req.Email, req.Password)

	err = h.AuthMethodSrv.SignUp(context.Background(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(coreShared.ErrorMsg{Msg: err.Error()})
		return
	}

	// Create response from the created user domain
	resp := userCore.NewFromUserDomain(user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
