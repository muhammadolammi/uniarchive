package main

import (
	"context"
	"fmt"
	"net/http"
)

func (s *state) usersGETHandler(w http.ResponseWriter, r *http.Request) {
	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting users from db. err: %v", err))
		return
	}
	respondWithJson(w, http.StatusOK, convertDBUsersToMainUsers(dbUsers))
}
