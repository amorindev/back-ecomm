package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(port string) *HttpServer {
	mux := http.NewServeMux()

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
