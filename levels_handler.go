package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/muhammadolammi/uniarchive/internal/database"
)

func (s *state) levelsPOSTHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Name string `json:"name"`
		Code int    `json:"code"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error decoding parameters. err: %v", err))
		return
	}
	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "provide a name for the department to be created")
		return
	}
	if params.Code == 0 {
		respondWithError(w, http.StatusBadRequest, "provide a level code to create a level, an int version of the name")
		return
	}
	_, err = s.db.CreateLevel(context.Background(), database.CreateLevelParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Code:      int32(params.Code),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating level on db err: %v", err))
		return
	}
	// respond with a success message
	respondWithJson(w, http.StatusOK, "level created")
}

func (s *state) levelsGETHandler(w http.ResponseWriter, r *http.Request) {
	dbLevels, err := s.db.GetLevels(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting levels from db. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, convertDBLevelsToMainLevels(dbLevels))
}
