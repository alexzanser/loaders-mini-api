package handler

import (
	"context"
	"fmt"
	"loaders/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

const signingKey = "la-la-la-la-la-la-la-la-la-la"

type loginHandler struct {
	service *service.Service
}

func newLoginHandler(service *service.Service) *loginHandler {
	return &loginHandler{
		service: service,
	}
}

func (l *loginHandler) Login(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling login at %s\n", req.URL.Path)

	err := req.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "can't parse request": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	username, passwd, role, err := formToLogin(req)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "can't parse request": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	token, err := l.GenerateToken(context.TODO(), username, passwd, role)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "can't generate token": "%v"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})

	rp := response{
		Token:      token,
		Result:     fmt.Sprintf(`{user with username: %s and role %s authorized}`, username, role),
		HTTPStatus: http.StatusAccepted,
	}
	renderResponse(w, rp)
}

func (l *loginHandler) GenerateToken(ctx context.Context, username, passwd, role string) (string, error) {
	var passwdHash string

	if role == "customer" {
		user, err := l.service.GetCustomer(ctx, username, passwd)
		if err != nil {
			return "", fmt.Errorf("can't get customer: %v", err)
		}
		passwdHash = user.PasswdHash
	}

	if role == "loader" {
		user, err := l.service.GetLoader(ctx, username, passwd)
		if err != nil {
			return "", fmt.Errorf("can't get loader: %v", err)
		}
		passwdHash = user.PasswdHash
	}

	type Claims struct {
		jwt.StandardClaims
		Username   string
		PasswdHash string
		Role       string
	}

	expire := time.Now().Add(time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(expire),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username:   username,
		PasswdHash: passwdHash,
		Role:       role,
	})

	return token.SignedString([]byte(signingKey))
}
