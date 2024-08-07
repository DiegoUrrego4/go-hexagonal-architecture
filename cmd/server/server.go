package server

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type ServerChi struct {
	Addr     string
	MySQLDSN string
	DB       *sql.DB
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
		log.Printf("Error initializing DB: %v", err)
		return
	}

	defer db.Close()
	s.DB = db

	s.registerRouter()
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

func (s *ServerChi) registerRouter() {
	// router
	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Route("/users", func(rt chi.Router) {
		//rt.Get("/")
	})

	log.Printf("ServerChi running on port %s", s.Addr)

	err := http.ListenAndServe(s.Addr, router)
	if err != nil {
		log.Printf("Error running server: %v", err)
		return
	}

}
