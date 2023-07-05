package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is empty")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handleError)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	fmt.Println("Listening port ", port)
	srv.ListenAndServe()
}