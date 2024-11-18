package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *state) materialsPOSTHandler(w http.ResponseWriter, r *http.Request) {

	// get params

	params := struct {
		Name         string    `json:"name"`
		CourseCode   string    `json:"course_code"`
		LevelId      uuid.UUID `json:"level_id"`
		DepartmentId uuid.UUID `json:"department_id"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error decoding parameters. err: %v", err))
		return
	}
	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "provide a name for the material to be created")
		return
	}
}

func (s *state) materialsGETHandler(w http.ResponseWriter, r *http.Request) {
	dbMaterials, err := s.db.GetMaterials(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting materials from db. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, convertDBMaterialsToMainMaterials(dbMaterials))
}
