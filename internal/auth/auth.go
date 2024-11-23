package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/muhammadolammi/uniarchive/internal/database"
)

func MakeJwtTokenString(tokenSigner []byte, userId, tokenName string, tokenExpiration int) (string, error) {
	if len(tokenSigner) == 0 {
		return "", errors.New("a token signer must be provided")
	}
	claims := jwt.RegisteredClaims{

		Issuer:    fmt.Sprintf("%v", userId),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(tokenExpiration) * time.Minute)),
		Subject:   tokenName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tokenSigner)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func UpdateAccessToken(tokenSigner []byte, userId uuid.UUID, expirationTime int, w http.ResponseWriter) error {

	jwtAccessTokenString, err := MakeJwtTokenString(tokenSigner, userId.String(), "accesstoken", expirationTime)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "accesstoken",
		Value:    jwtAccessTokenString,
		Expires:  time.Now().UTC().Add(time.Duration(expirationTime) * time.Minute),
		HttpOnly: true,
		Path:     "/",
		// TODO REMEBER TO COMMENT ALL THESE OUT
		// SameSite: http.SameSiteLaxMode,
		Secure: false, // Set to true in production when using HTTPS
	})
	return nil

}
func UpdateRefreshToken(signgingKey []byte, userId uuid.UUID, expirationTime int, DB *database.Queries) error {
	jwtRefreshTokenString, err := MakeJwtTokenString(signgingKey, userId.String(), "refreshtoken", expirationTime)
	if err != nil {
		return err
	}

	err = DB.UpdateRefreshToken(context.Background(), database.UpdateRefreshTokenParams{
		RefreshToken: sql.NullString{
			Valid:  true,
			String: jwtRefreshTokenString,
		},
		UserID: userId,
	})
	if err != nil {
		return err
	}
	return nil
}
