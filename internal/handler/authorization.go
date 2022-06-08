package handler

import (
	"fmt"
	"loaders/internal/service"
	"net/http"
	"strings"
	"context"
	"github.com/dgrijalva/jwt-go/v4"
)

type authorizationHandler struct {
	service *service.Service
}

func newAuthorizationHandler (service *service.Service) *authorizationHandler {
	return &authorizationHandler{service: service}
}

func (a *authorizationHandler) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token, err := retrieveToken(req)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusUnauthorized)
			return
		}


		username, role, err  := parseToken(token)
		if err != nil || role != "loader" && role != "customer" {
			http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(context.TODO(), "username", username)
		ctx	= context.WithValue(ctx, "role", role)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

// func (a *authorizationHandler) AuthorizeForCustomer(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		token, err := retrieveToken(req)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusUnauthorized)
// 			return
// 		}

// 		username, role, err  := parseToken(token)
// 		if err != nil || role != "customer" { 
// 			http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusUnauthorized)
// 			return
// 		}

// 		ctx := context.WithValue(req.Context(), "username", username)
// 		ctx	= context.WithValue(ctx, "role", role)
// 		req = req.WithContext(ctx)
// 		next.ServeHTTP(w, req)
// 	})
// }

func retrieveToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", fmt.Errorf("invalid auth header")
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", fmt.Errorf("invalid auth header")
	}

	return headerParts[1], nil
}

func parseToken(accessToken string) (string, string, error) {
	type Claims struct {
		jwt.StandardClaims
		Username	string
		PasswdHash	string
		Role		string
	}
	
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	
	if err != nil {
		return "", "", fmt.Errorf("invalid acess token: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, claims.Role, nil
	}

	return "", "", fmt.Errorf("invalid acess token")
}
