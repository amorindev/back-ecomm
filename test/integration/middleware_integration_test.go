package integration

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/amorindev/go-tmpl/internal/auth"
	"github.com/amorindev/go-tmpl/internal/config"
	"github.com/amorindev/go-tmpl/pkg/shared/api/middlewares"
	"github.com/joho/godotenv"
)

// Manual test server (use Postman or curl to test the endpoints)
// This is NOT meant for production, only for verifying middleware and token flow.
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	appEnvs := config.Load()

	tokenSrv := auth.NewTokenSrv(appEnvs.JWTAccessSecret, appEnvs.JWTRefreshSecret, appEnvs.JWTAccessExpIn, 3*time.Second, appEnvs.JWTRefreshRememberMeExpIn, appEnvs.JWTIssuer)

	mdwSrv := middlewares.NewAuthMdw(tokenSrv)

	// Protected endpoint: requires a valid Access Token
	// Example: GET /protected with header Authorization: Bearer <access_token>
	mux.HandleFunc("GET /protected", mdwSrv.AccessTokenMdw(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := struct {
			Msg string `json:"msg"`
		}{
			Msg: "Valid access token. Welcome!",
		}
		json.NewEncoder(w).Encode(resp)
	}))

	// Refresh endpoint: requires a valid Refresh Token
	// Example: POST /refresh with JSON body {"refresh_token": "<your_token>"}
	mux.HandleFunc("POST /refresh", mdwSrv.RefreshTokenMdw(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := struct {
			Msg string `json:"msg"`
		}{
			Msg: "Valid refresh token. Welcome!",
		}
		json.NewEncoder(w).Encode(resp)
	}))

	// Tokens endpoint: generates new Access and Refresh tokens for manual testing
	// Example: GET /tokens → returns {access_token, refresh_token, expires_in}
	mux.HandleFunc("GET /tokens", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		accessToken, expiresIn, err := tokenSrv.CreateAccessToken("1", "test@gmail.com", nil)
		if err != nil {
			resp := struct {
				Msg string `json:"msg"`
			}{
				Msg: err.Error(),
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}

		_, refreshToken, _, err := tokenSrv.CreateRefreshToken("1", false)
		if err != nil {
			resp := struct {
				Msg string `json:"msg"`
			}{
				Msg: err.Error(),
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp := struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			ExpiresIn    int64  `json:"expires_in"`
		}{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    expiresIn,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}
