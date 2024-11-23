package main

// acess  token  5 minutes, that refresh every 5 minutes and we use refresh token to refresh
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/muhammadolammi/uniarchive/internal/auth"
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
		return
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
		MatricNumber: params.MatricNumber,
	})
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), `pq: duplicate key value violates unique constraint "users_email_key"`) {
			respondWithError(w, http.StatusUnauthorized, "user with that email already exist")
			return
		}
		if strings.Contains(err.Error(), `pq: duplicate key value violates unique constraint "users_matric_number_key"`) {
			respondWithError(w, http.StatusUnauthorized, "user with that matric number already exist")
			return
		}

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
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		if strings.Contains(err.Error(), `hashedPassword is not the hash of the given password`) {
			respondWithError(w, http.StatusUnauthorized, "Wrong password.")
			return
		}
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("error hasing password err: %v", err))
		return
	}
	// 6 minute here to give the frontend time to call refresh

	err = auth.UpdateAccessToken([]byte(s.JWTSIGNER), user.ID, accessTokenExpirationTime, w)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("error updating access token err: %v", err))
		return

	}

	err = auth.UpdateRefreshToken([]byte(s.JWTSIGNER), user.ID, refreshTokenExpirationTime, s.db)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("error updating refresh token err: %v", err))
		return

	}

	respondWithJson(w, http.StatusOK, "log in successful")

}

func (s *state) validateHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, http.StatusOK, convertDBUserToMainUser(user))
}

func (s *state) refreshHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	refreshToken := user.RefreshToken
	// Verify the token
	token, err := jwt.Parse(refreshToken.String, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.JWTSIGNER), nil
	})

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("error parsing token string to jwt. err: %v", err))
		return
	}

	if !token.Valid {
		respondWithError(w, http.StatusUnauthorized, "invalid refresh token")
		return
	}

	// If the refresh token is valid
	err = auth.UpdateAccessToken([]byte(s.JWTSIGNER), user.ID, accessTokenExpirationTime, w)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("error refreshing access token. err: %v", err))
		return
	}

	respondWithJson(w, http.StatusOK, "tokens refreshed")

}
