package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func server(s *state) {

	// Define CORS options
	corsOptions := cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://muhammaddev.com", "http://192.168.246.175:3000"}, // You can customize this based on your needs

		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // You can customize this based on your needs
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for cache, in seconds
	}
	router := chi.NewRouter()
	apiRoute := chi.NewRouter()
	// ADD MIDDLREWARE
	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(corsOptions))
	apiRoute.Get("/ready", helloReady)
	apiRoute.Get("/error", errorReady)

	apiRoute.Post("/universities", s.universitiesPOSTHandler)
	apiRoute.Get("/universities", s.universitiesGETHandler)
	apiRoute.Patch("/universities", s.universitiesPATCHHandler)

	router.Mount("/api", apiRoute)

	srv := &http.Server{
		Addr:              ":" + s.PORT,
		Handler:           router,
		ReadHeaderTimeout: time.Minute,
	}

	log.Printf("Serving on port: %s\n", s.PORT)
	log.Fatal(srv.ListenAndServe())

}
