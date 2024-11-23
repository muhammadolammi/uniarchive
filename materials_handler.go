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

func (s *state) materialsPOSTHandler(w http.ResponseWriter, r *http.Request) {

	// get params

	params := struct {
		Name     string    `json:"name"`
		CourseId uuid.UUID `json:"course_id"`
		CloudUrl string    `json:"cloud_url"`
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
	if params.CloudUrl == "" {
		respondWithError(w, http.StatusBadRequest, "provide a cloud url for the material to be created")
		return
	}
	if params.CourseId == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide a course id for the material to be created")
		return
	}
	_, err = s.db.CreateMaterial(context.Background(), database.CreateMaterialParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		CourseID:  params.CourseId,
		CloudUrl:  params.CloudUrl,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating material. err:%v", err))
		return
	}
	respondWithJson(w, http.StatusOK, "material created")
}

func (s *state) materialsGETHandler(w http.ResponseWriter, r *http.Request) {
	courseIdString := chi.URLParam(r, "courseID")
	if courseIdString == "default" {
		params := struct {
			DepartmentId uuid.UUID `json:"department_id"`
		}{}
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&params)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error decoding parameters. err: %v", err))
			return
		}
		if params.DepartmentId.String() == "" {
			respondWithError(w, http.StatusBadRequest, "provide a department id to get materials")
			return
		}
		dbMaterials, err := s.db.GetDefaultMaterials(context.Background(), params.DepartmentId)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting materials from db. err: %v", err))
			return
		}
		respondWithJson(w, http.StatusOK, convertDBMaterialsToMainMaterials(dbMaterials))
		return
	}

	courseID, err := uuid.Parse(courseIdString)
	if err != nil {
		respondWithError(w, 501, fmt.Sprintf("error  parsing course id to uuid. err :%v", err))
		return
	}

	dbMaterials, err := s.db.GetCourseMaterials(context.Background(), courseID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting materials from db. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, convertDBMaterialsToMainMaterials(dbMaterials))

}
