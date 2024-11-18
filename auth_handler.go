package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadolammi/uniarchive/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (s *state) signUpHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Email        string    `json:"email"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name"`
		OtherName    string    `json:"other_name"`
		MatricNumber string    `json:"matric_number"`
		Password     string    `json:"password"`
		UniversityID uuid.UUID `json:"university_id"`
		FacultyID    uuid.UUID `json:"faculty_id"`
		DepartmentID uuid.UUID `json:"department_id"`
		LevelID      uuid.UUID `json:"level_id"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error decoding parameters. err: %v", err))
		return
	}
	if params.Email == "" {
		respondWithError(w, http.StatusBadRequest, "provide a email for the user to be created")
		return
	}
	if params.MatricNumber == "" {
		respondWithError(w, http.StatusBadRequest, "provide a matric number for the user to be created")
		return
	}
	if params.FirstName == "" {
		respondWithError(w, http.StatusBadRequest, "provide a first name for the user to be created")
		return
	}
	if params.LastName == "" {
		respondWithError(w, http.StatusBadRequest, "provide a last name for the user to be created")
		return
	}
	if params.Password == "" {
		respondWithError(w, http.StatusBadRequest, "user must provide a password")
		return
	}
	if params.LevelID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide the level id for the user to be created")
		return
	}
	if params.DepartmentID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide the department id for the user to be created")
		return
	}
	if params.UniversityID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide the university id for the user to be created")
		return
	}
	if params.FacultyID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide the faculty id for the user to be created")
		return
	}
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error encrypting password. err: %v", err))
	}
	_, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Email:        params.Email,
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		OtherName:    params.OtherName,
		Password:     string(hashedPassword),
		UniversityID: params.UniversityID,
		FacultyID:    params.FacultyID,
		DepartmentID: params.DepartmentID,
		LevelID:      params.LevelID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating user. err:%v", err))
		return
	}
	respondWithJson(w, http.StatusOK, "user created")
}

func (s *state) signInHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error decoding parameters. err: %v", err))
		return
	}
	if params.Email == "" {
		respondWithError(w, http.StatusBadRequest, "provide an email to sign in")
		return
	}
	if params.Password == "" {
		respondWithError(w, http.StatusBadRequest, "user must provide a password")
		return
	}
	// get the user with the email
	user, err := s.db.GetUserWithEmail(context.Background(), params.Email)
	if err != nil {
		if strings.Contains(err.Error(), `sql: no rows in result set`) {
			respondWithError(w, http.StatusUnauthorized, "user doesn't exist")
			return
		}
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting user with that email. err: %v", err))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		if strings.Contains(err.Error(), `hashedPassword is not the hash of the given password`) {
			respondWithError(w, http.StatusUnauthorized, "Wrong password.")
			return
		}
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf(" err: %v", err))
		return
	}
	// TODO write JWT that generate a logged in token on the server http
	respondWithJson(w, http.StatusOK, "log in successful")

}
