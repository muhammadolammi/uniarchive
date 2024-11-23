package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/muhammadolammi/uniarchive/internal/database"
)

func (s *state) usersGETHandler(w http.ResponseWriter, r *http.Request) {
	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting users from db. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, convertDBUsersToMainUsers(dbUsers))
}

func (s *state) userProfilePictureHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		UserId     uuid.UUID `json:"user_id"`
		ProfileUrl string    `json:"profile_url"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error decoding params. err: %v", err))
		return
	}
	if params.UserId == uuid.Nil {
		respondWithError(w, http.StatusUnauthorized, "provide a user id")
		return

	}
	if params.ProfileUrl == "" {
		respondWithError(w, http.StatusUnauthorized, "provide a profile url")
		return
	}
	err = s.db.UpdateUserProfilePicture(context.Background(), database.UpdateUserProfilePictureParams{
		UserID: params.UserId,
		ProfilePicture: sql.NullString{
			Valid:  true,
			String: params.ProfileUrl,
		},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error uploading profile picture to server. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, "profile picture uploaded")
}
