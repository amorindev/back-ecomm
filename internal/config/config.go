package config

import (
	"cmp"
	"log"
	"os"
	"time"
)

type Config struct {
	// MongoDB
	MongoDBUri  string
	MongoInitDB string

	// Minio
	MinioEndpoint   string
	MinioAccessKey  string
	MinioSecretKey  string
	MinioUseSSL     bool
	MinioBucketName string

	// Jwt
	JWTAccessSecret           string
	JWTRefreshSecret          string
	JWTIssuer                 string
	JWTAccessExpIn            time.Duration
	JWTRefreshExpIn           time.Duration
	JWTRefreshRememberMeExpIn time.Duration

	// App
	Port string

	// Templates
	ApiBaseUrl string
}

func Load() *Config {
	// Mongo DB
	mongoInitDB := cmp.Or(os.Getenv("MONGO_INITDB_DATABASE"), "auth-tmpl")

	// Minio
	minioBucketName := cmp.Or(os.Getenv("MINIO_BUCKET_NAME"), "auth-tmpl")
	useSSL := mustGetEnv("MINIO_SECURE")

	var useSSLbool bool
	if useSSL == "true" || useSSL == "yes" {
		useSSLbool = true
	} else {
		useSSLbool = false
	}

	// Auth - tokens
	accessExp := cmp.Or(os.Getenv("JWT_ACCESS_EXP_IN"), "15m")
	refreshExp := cmp.Or(os.Getenv("JWT_REFRESH_EXP_IN"), "168h")
	refreshExpRememberMe := cmp.Or(os.Getenv("JWT_REFRESH_EXP_IN_REMEMBER"), "720h")

	accessDur, err := time.ParseDuration(accessExp)
	if err != nil {
		log.Fatalf("Invalid JWT_ACCESS_EXP_IN format: %v", err)
	}

	refreshDur, err := time.ParseDuration(refreshExp)
	if err != nil {
		log.Fatalf("Invalid JWT_REFRESH_EXP_IN format: %v", err)
	}

	refreshRememberMeDur, err := time.ParseDuration(refreshExpRememberMe)
	if err != nil {
		log.Fatalf("Invalid JWT_REFRESH_EXP_IN_REMEMBER format: %v", err)
	}

	// App
	port := cmp.Or(os.Getenv("HTTP_SERVER_PORT"), "8000")
	apiBaseUrl := cmp.Or(os.Getenv("API_BASE_URL"), "http://localhost:"+port)

	return &Config{
		MongoDBUri:                mustGetEnv("MONGO_DB_URI"),
		MongoInitDB:               mongoInitDB,
		MinioEndpoint:             mustGetEnv("MINIO_ENDPOINT"),
		MinioAccessKey:            mustGetEnv("MINIO_ACCESS_KEY"),
		MinioSecretKey:            mustGetEnv("MINIO_SECRET_KEY"),
		MinioUseSSL:               useSSLbool,
		MinioBucketName:           minioBucketName,
		JWTAccessSecret:           mustGetEnv("JWT_ACCESS_TOKEN"),
		JWTRefreshSecret:          mustGetEnv("JWT_REFRESH_TOKEN"),
		JWTIssuer:                 mustGetEnv("JWT_ISS"),
		JWTAccessExpIn:            accessDur,
		JWTRefreshExpIn:           refreshDur,
		JWTRefreshRememberMeExpIn: refreshRememberMeDur,
		Port:                      port,
		ApiBaseUrl:                apiBaseUrl,
	}
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return val
}
