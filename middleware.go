package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/muhammadolammi/uniarchive/internal/database"
)

//TODO anyone can create users, new user can create account. But only devs can upgrade user to admin.
// i may also have to add devs with sql on server since they control most authentication

func (s *state) middlewareLoggedIn(handler func(w http.ResponseWriter, r *http.Request, user database.User)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get the access token from cookies
		cookie, err := r.Cookie("accesstoken")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println(err)
				respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("missing access token. err: %v", err))
				return
			}
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error retrieving access token from cookies. err: %v", err))
			return
		}
		accessToken := cookie.Value

		// Verify the token
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.JWTSIGNER), nil
		})

		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("error parsing token string to jwt. err: %v", err))
			return
		}

		if !token.Valid {
			respondWithError(w, http.StatusUnauthorized, "invalid access token")
			return
		}

		// Extract the user ID from the claims
		userIDSTRING, err := token.Claims.GetIssuer()
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("error getting user id from token claims err: %v", err))
			return
		}

		userID, err := uuid.Parse(userIDSTRING)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error parsing id to uuid. err: %v", err))
			return
		}
		currentUser, err := s.db.GetUser(context.Background(), userID)
		if err != nil {

			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting current user token from db id. err: %v", err))
			return
		}

		// CALL THE HANDLER
		handler(w, r, currentUser)
		// TODO

	}

}
