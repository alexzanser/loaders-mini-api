package handler

import (
	"fmt"
	"loaders/internal/models"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func formToCustomerRegister(ct *models.Customer, req *http.Request) error {
	ct.Username = req.PostFormValue("username")
	ct.Passwd = req.PostFormValue("password")

	b, err := strconv.Atoi(req.PostFormValue("balance"))
	if err != nil {
		return fmt.Errorf("error when create customer: %v", err)
	}
	ct.Balance = b

	if ct.Username == "" || ct.Passwd == "" || ct.Balance <= 0 {
		return fmt.Errorf("error when create customer: ivalid arguments")
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(ct.Passwd), 10)
	if err != nil {
		return fmt.Errorf("error when create customer: %v", err)
	}
	ct.PasswdHash = string(hashBytes)

	return nil
}

func formToLoaderRegister(ld *models.Loader, req *http.Request) error {
	ld.Username = req.PostFormValue("username")
	ld.Passwd = req.PostFormValue("password")

	if ld.Username == "" || ld.Passwd == "" {
		return fmt.Errorf("error when create customer: ivalid arguments")
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(ld.Passwd), 10)
	if err != nil {
		return fmt.Errorf("error when create customer: %v", err)
	}
	ld.PasswdHash = string(hashBytes)

	return nil
}

func formToLogin(req *http.Request) (string, string, string, error) {
	username := req.PostFormValue("username")
	passwd := req.PostFormValue("password")
	role := req.PostFormValue("role")

	if username == "" || passwd == "" || role != "loader" && role != "customer" {
		return "", "", "", fmt.Errorf("error when login user: ivalid arguments")
	}

	return username, passwd, role, nil
}
