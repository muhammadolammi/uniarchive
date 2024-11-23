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

func (s *state) coursesPOSTHandler(w http.ResponseWriter, r *http.Request) {
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
		respondWithError(w, http.StatusBadRequest, "provide a name for the course to be created")
		return
	}
	if params.LevelId == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide the level id for the course to be created")
		return
	}
	if params.DepartmentId == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "provide the department id for the course to be created")
		return
	}
	_, err = s.db.CreateCourse(context.Background(), database.CreateCourseParams{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Name:         params.Name,
		CourseCode:   params.CourseCode,
		LevelID:      params.LevelId,
		DepartmentID: params.DepartmentId,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating course err:%v", err))
		return
	}

	respondWithJson(w, http.StatusOK, "course created")
}

func (s *state) coursesGETHandler(w http.ResponseWriter, r *http.Request) {
	userIdString := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIdString)
	if err != nil {
		respondWithError(w, 501, fmt.Sprintf("error  parsing user id to uuid. err :%v", err))
		return
	}
	dbCourses, err := s.db.GetUserCourses(context.Background(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting courses from db. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, convertDBCoursesToMainCourses(dbCourses))
}
