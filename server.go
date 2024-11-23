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
		AllowedOrigins: []string{"http://172.23.206.70:8080", "localhost"}, // You can customize this based on your needs

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

	//universities endpoint
	apiRoute.Post("/universities", s.universitiesPOSTHandler)
	apiRoute.Get("/universities", s.universitiesGETHandler)
	apiRoute.Patch("/universities", s.universitiesPATCHHandler)

	// faculties endpoint
	apiRoute.Post("/faculties", s.facultiesPOSTHandler)
	apiRoute.Get("/faculties/{universityID}", s.facultiesGETHandler)

	// departments endpoint
	apiRoute.Post("/departments", s.departmentsPOSTHandler)
	apiRoute.Get("/departments/{facultyID}", s.departmentsGETHandler)

	// levels endpoint
	apiRoute.Post("/levels", s.levelsPOSTHandler)
	apiRoute.Get("/levels", s.levelsGETHandler)

	// courses endpoint
	apiRoute.Post("/courses", s.coursesPOSTHandler)
	apiRoute.Get("/courses/{userID}", s.coursesGETHandler)

	// materials endpoint
	apiRoute.Post("/materials", s.materialsPOSTHandler)
	apiRoute.Get("/materials/{courseID}", s.materialsGETHandler)
	// auth endpoint
	apiRoute.Post("/signup", s.signUpHandler)
	apiRoute.Post("/signin", s.signInHandler)
	apiRoute.Post("/validate", s.middlewareLoggedIn(s.validateHandler))
	apiRoute.Post("/refresh", s.middlewareLoggedIn(s.refreshHandler))

	// users endpoint
	apiRoute.Get("/users", s.usersGETHandler)
	apiRoute.Post("/user/profilepicture", s.userProfilePictureHandler)

	router.Mount("/api", apiRoute)

	srv := &http.Server{
		Addr:              ":" + s.PORT,
		Handler:           router,
		ReadHeaderTimeout: time.Minute,
	}

	log.Printf("Serving on port: %s\n", s.PORT)
	log.Fatal(srv.ListenAndServe())

}
