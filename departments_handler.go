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

func (s *state) departmentsPOSTHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Name      string    `json:"name"`
		FacultyId uuid.UUID `json:"faculty_id"`
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
	if params.FacultyId == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide a faculty id to add add a department")
		return
	}
	_, err = s.db.CreateDepartment(context.Background(), database.CreateDepartmentParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		FacultyID: params.FacultyId,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating department on db err: %v", err))
		return
	}
	// respond with a success message
	respondWithJson(w, http.StatusOK, "department created")
}

func (s *state) departmentsGETHandler(w http.ResponseWriter, r *http.Request) {
	facultyIdString := chi.URLParam(r, "facultyID")
	facultyID, err := uuid.Parse(facultyIdString)
	if err != nil {
		respondWithError(w, 501, fmt.Sprintf("error  parsing faculty id to uuid. err :%v", err))
		return
	}
	dbDepartments, err := s.db.GetDepartments(context.Background(), facultyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting faculties from db. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, convertDBDepartmentsToMainDepartments(dbDepartments))
}
