package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amorindev/go-tmpl/internal/config"
	mongoClient "github.com/amorindev/go-tmpl/internal/mongo"
	"github.com/amorindev/go-tmpl/pkg/app/admin/api/handler"
	authMethodHandler "github.com/amorindev/go-tmpl/pkg/app/auth-methods/handler"
	authMethodService "github.com/amorindev/go-tmpl/pkg/app/auth-methods/service"
	userRepository "github.com/amorindev/go-tmpl/pkg/app/users/repository/mongo"
)

func New() http.Handler {
	mux := http.NewServeMux()

	// Api version
	v1 := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	appEnvs := config.Load()

	// MongoDB
	mongoConn := mongoClient.New(appEnvs.MongoDBUri)
	mongoDB := mongoConn.DB.Database(appEnvs.MongoInitDB)
	mongoConn.Ping()

	// Collections
	userColl := mongoDB.Collection("users")

	// Repositories
	userRepo := userRepository.NewUserRepo(mongoConn.DB, userColl)

	// Indexes
	err := userRepo.CreateIndexes()
	if err != nil {
		log.Fatal(err)
	}

	// Services
	authMethodSrv := authMethodService.NewAuthMethodSrv(userRepo)

	// Handler
	// Note: all subsequent handlers should also be registered using v1
	authMethodHandler.NewAuthMethodHandler(v1, authMethodSrv)

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := struct {
			Msg string `json:"msg"`
		}{
			Msg: "pong",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})

	// Templates
	// Redirects requests from "/admin" to the admin home page under API v1
	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/v1/admin/home", http.StatusFound)
	})

	handler.NewAdminHandler(v1, appEnvs.ApiBaseUrl)

	return mux
}
