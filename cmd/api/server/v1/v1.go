package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amorindev/go-tmpl/internal/config"
	minioClient "github.com/amorindev/go-tmpl/internal/minio"
	mongoClient "github.com/amorindev/go-tmpl/internal/mongo"
	adminHandler "github.com/amorindev/go-tmpl/pkg/app/admin/api/handler"
	authMethodHandler "github.com/amorindev/go-tmpl/pkg/app/auth-methods/handler"
	authMethodService "github.com/amorindev/go-tmpl/pkg/app/auth-methods/service"
	categoryHandler "github.com/amorindev/go-tmpl/pkg/app/ecomm/category/api/handler"
	categoryRepository "github.com/amorindev/go-tmpl/pkg/app/ecomm/category/repository/mongo"
	categoryService "github.com/amorindev/go-tmpl/pkg/app/ecomm/category/service"
	variationHandler "github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/api/handler"
	varOptionRepository "github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/repository/var-option/mongo"
	variationRepository "github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/repository/variation/mongo"
	variationService "github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/service"
	userRepository "github.com/amorindev/go-tmpl/pkg/app/users/repository/mongo"
	minioAdapter "github.com/amorindev/go-tmpl/pkg/file-storage/adapter/minio"
	fileStgService "github.com/amorindev/go-tmpl/pkg/file-storage/service"
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

	// Minio
	minioC, err := minioClient.NewClient(appEnvs.MinioEndpoint, appEnvs.MinioAccessKey, appEnvs.MinioSecretKey, appEnvs.MinioUseSSL)
	if err != nil {
		log.Fatal(err)
	}

	err = minioC.CreateStorage(appEnvs.MinioBucketName)
	if err != nil {
		log.Fatal(err)
	}

	minioApt := minioAdapter.NewMinioAdt(minioC.Client, appEnvs.MinioBucketName)
	_ = fileStgService.NewFileStgSrv(minioApt)

	// Collections
	userColl := mongoDB.Collection("users")
	categoryColl := mongoDB.Collection("categories")
	variationColl := mongoDB.Collection("variations")
	varOptionColl := mongoDB.Collection("var_options")

	// Repositories
	userRepo := userRepository.NewUserRepo(mongoConn.DB, userColl)
	categoryRepo := categoryRepository.NewCategoryRepo(mongoConn.DB, categoryColl)
	variationRepo := variationRepository.NewVariationRepo(mongoConn.DB, variationColl)
	varOptionRepo := varOptionRepository.NewVarOptionRepo(mongoConn.DB, varOptionColl)

	// Indexes
	err = userRepo.CreateIndexes()
	if err != nil {
		log.Fatal(err)
	}

	// Services
	authMethodSrv := authMethodService.NewAuthMethodSrv(userRepo)
	categorySrv := categoryService.NewCategorySrv(categoryRepo)
	variationSrv := variationService.NewVariationSrv(variationRepo, varOptionRepo)

	// Handler
	// Note: all subsequent handlers should also be registered using v1
	authMethodHandler.NewAuthMethodHandler(v1, authMethodSrv)
	categoryHandler.NewCategoryHandler(v1, categorySrv)
	variationHandler.NewVariationHandler(v1, variationSrv)

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
		http.Redirect(w, r, "/v1/admin/categories", http.StatusFound)
	})

	adminHandler.NewAdminHandler(v1, appEnvs.ApiBaseUrl)

	return mux
}
