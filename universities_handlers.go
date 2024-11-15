package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadolammi/uniarchive/internal/database"
)

func (s *state) universitiesPOSTHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Name  string `json:"name"`
		Alias string `json:"alias"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error decoding parameters. err: %v", err))
		return
	}
	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "provide a name for the univerity to be created")
		return
	}
	// default the alias to the university name if non  was provided
	if params.Alias == "" {
		params.Alias = params.Name
	}
	_, err = s.db.CreateUniversity(context.Background(), database.CreateUniversityParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Alias:     params.Alias,
	})
	// if there is an error creating university let the client know
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating university on db err: %v", err))
		return
	}
	// respond with a success message
	respondWithJson(w, http.StatusOK, "university created")

}

func (s *state) universitiesGETHandler(w http.ResponseWriter, r *http.Request) {
	dbUnis, err := s.db.GetUniversities(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting universities from db. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, convertDBUnisToMainUnis(dbUnis))
}

func (s *state) universitiesPATCHHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Name  string    `json:"name"`
		Alias string    `json:"alias"`
		ID    uuid.UUID `json:"id"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error decoding parameters. err: %v", err))
		return
	}
	if params.ID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide an id to edit university")
		return
	}
	if params.Name == "" && params.Alias == "" {
		respondWithError(w, http.StatusBadRequest, "provide a name or alias to update")
		return
	}

	updateParams := database.EditUniversityParams{
		UpdatedAt: time.Now(),
		ID:        params.ID,
	}
	if params.Alias != "" {
		updateParams.Alias = params.Alias
	}
	if params.Name != "" {
		updateParams.Name = params.Name
	}

	err = s.db.EditUniversity(context.Background(), updateParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error editing university. err: %v", err))
	}

	respondWithJson(w, http.StatusOK, "university edited")
}
