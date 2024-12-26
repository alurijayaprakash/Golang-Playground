package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type APIServer struct {
	listenAddr string
	Port       string
}

// NewAPIServer Returns API Server Object
func NewAPIServer(listenAddr, port string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		Port:       port,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
