package server

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time-tracker/cmd/api/handlers/user"
	userRepo "time-tracker/internal/user/repositories/mysql/user"
	"time-tracker/internal/user/services"
)

type ServerChi struct {
	Addr        string
	MySQLDSN    string
	DB          *sql.DB
	UserRepo    userRepo.Repository
	UserService services.Service
	UserHandler user.Handler
}

func New(cfg ServerChi) *ServerChi {
	return &ServerChi{
		Addr:     cfg.Addr,
		MySQLDSN: cfg.MySQLDSN,
	}
}

func (s *ServerChi) Run() {
	db, err := s.initDB()
	if err != nil {
		log.Fatalf("Error initializing DB: %v", err)
		return
	}

	defer db.Close()
	s.DB = db

	s.registerRouter(db)
}

func (s *ServerChi) initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", s.MySQLDSN)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *ServerChi) registerRouter(db *sql.DB) {
	// router
	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	s.UserRepo = userRepo.Repository{
		Db: db,
	}
	s.UserService = services.Service{
		Repo: s.UserRepo,
	}
	s.UserHandler = user.Handler{
		UserService: s.UserService,
	}
	router.Route("/users", func(rt chi.Router) {
		rt.Post("/", s.UserHandler.CreateUser)
	})

	log.Printf("ServerChi running on port %s", s.Addr)

	err := http.ListenAndServe(s.Addr, router)
	if err != nil {
		log.Fatalf("Error running server: %v", err)
		return
	}

}
