package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/amorindev/go-tmpl/internal/config"
	mongoClient "github.com/amorindev/go-tmpl/internal/mongo"
	"github.com/amorindev/go-tmpl/pkg/app/admin/api/handler"
	authMethodHandler "github.com/amorindev/go-tmpl/pkg/app/auth-methods/handler"
	authMethodService "github.com/amorindev/go-tmpl/pkg/app/auth-methods/service"
	userRepository "github.com/amorindev/go-tmpl/pkg/app/users/repository/mongo"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(port string) *HttpServer {
	mux := http.NewServeMux()

	appEnvs := config.Load()

	// * MongoDB
	dbURI := os.Getenv("MONGO_DB_URI")
	if dbURI == "" {
		log.Fatal("missing required environment variable: MONGO_DB_URI")
	}

	dbName := os.Getenv("MONGO_INITDB_DATABASE")
	if dbName == "" {
		log.Fatal("missing required environment variable: DB_NAME")
	}

	mongoConn := mongoClient.New(dbURI)
	mongoDB := mongoConn.DB.Database(dbName)
	mongoConn.Ping()

	// * Collections
	userColl := mongoDB.Collection("users")

	// * Repositories and indexes
	userRepo := userRepository.NewUserRepo(mongoConn.DB, userColl)

	// * Indexes
	err := userRepo.CreateIndexes()
	if err != nil {
		log.Fatal(err)
	}

	// * Services
	authMethodSrv := authMethodService.NewAuthMethodSrv(userRepo)

	// * Handler
	authMethodHandler.NewAuthMethodHandler(mux, authMethodSrv)

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

	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin/home", http.StatusFound)
	})

	handler.NewAdminHandler(mux, appEnvs.ApiBaseUrl)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serv := &HttpServer{
		server: server,
	}

	return serv
}

func (serv *HttpServer) Start() {
	log.Printf("Http server running http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
