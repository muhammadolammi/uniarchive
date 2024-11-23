package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/muhammadolammi/uniarchive/internal/database"
)

func (s *state) facultiesPOSTHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Name         string    `json:"name"`
		UniversityId uuid.UUID `json:"university_id"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error decoding parameters. err: %v", err))
		return
	}
	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "provide a name for the faculty to be created")
		return
	}
	if params.UniversityId == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide a university id to add add a faculty")
		return
	}
	_, err = s.db.CreateFaculty(context.Background(), database.CreateFacultyParams{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Name:         params.Name,
		UniversityID: params.UniversityId,
	})
	// if there is an error creating faculty let the client know
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating faculty on db err: %v", err))
		return
	}
	// respond with a success message
	respondWithJson(w, http.StatusOK, "faculty created")

}

func (s *state) facultiesGETHandler(w http.ResponseWriter, r *http.Request) {
	universityIdString := chi.URLParam(r, "universityID")
	universityID, err := uuid.Parse(universityIdString)
	if err != nil {
		respondWithError(w, 501, fmt.Sprintf("error  parsing university id to uuid. err :%v", err))
		return
	}
	dbFaculties, err := s.db.GetFaculties(context.Background(), universityID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting faculties from db. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, convertDBFacultiesToMainFaculties(dbFaculties))
}
