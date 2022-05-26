package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"movie-backend/models"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	jwt struct {
		secret string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var cfg config
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Printf("Error converting $PORT to an int: %q - Using default\n", err)
	}

	flag.IntVar(&cfg.port, "port", port, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
	flag.StringVar(&cfg.db.dsn, "dsn", "root://Goodman8349**@localhost/go_movies?sslmode=disable", "Postgres connection string")
	flag.Parse()
	cfg.jwt.secret = os.Getenv("GO_MOVIE_JWT")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	password := os.Getenv("DB_PASSWORD")
	db, err := sql.Open("postgres", "user=postgres password="+password+" host=127.0.0.1 port=5432 dbname=go_movies sslmode=disable")

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
