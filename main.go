package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/muhammadolammi/uniarchive/internal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("there is no port provided kindly provide a port.")
		return
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Println("there is no port provided kindly provide a port.")
		return
	}
	jwtSigner := os.Getenv("JWT_SIGNER")
	if jwtSigner == "" {
		log.Println("there is jwtSigner provided kindly provide a jwtSigner.")
		return
	}
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Println("error opening an sql connection. err:", err)
		os.Exit(1)
	}
	db := database.New(dbConn)

	s := state{
		db:        db,
		PORT:      port,
		JWTSIGNER: jwtSigner,
	}

	server(&s)
}
