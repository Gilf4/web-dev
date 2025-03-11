package api

import (
	"GoForBeginner/internal/repository"
	"GoForBeginner/internal/service/user"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
)

type APIServer struct {
	db   *pgxpool.Pool
	addr string
}

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()

	userRepo := repository.NewUserStore(s.db)

	router.Route("/api/v1", func(r chi.Router) {
		userHandler := user.NewHandler(userRepo)
		userHandler.RegisterRoutes(r)
	})

	log.Println("Listening on", s.addr)

	err := http.ListenAndServe(s.addr, router)
	if err != nil {
		return err
	}
	return nil
}
