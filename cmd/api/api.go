package api

import (
	"database/sql"
	userServices "ecom-tiago/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type APIServer struct {
	addr	string;
	db		*sql.DB;
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr, db};
}

// docs: Run function for API Server
func (s *APIServer) Run() error {
	router := mux.NewRouter();
	subrouter := router.PathPrefix("/api/v1").Subrouter();

	// docs: Register the routes with injected dependencies
	userStore := userServices.NewStore(s.db);
	userHandler := userServices.NewHandler(userStore);
	userHandler.RegisterRoutes(subrouter);
	
	log.Println("Listening on", s.addr);
	
	return http.ListenAndServe(s.addr, router);
}