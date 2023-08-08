package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type EncryptRequest struct {
	Password     string `json:"password"`
	EncryptionMethod string `json:"encryption_method"`
}

type EncryptResponse struct {
	EncryptedPassword string `json:"encrypted_password"`
}

func encryptPassword(password string, method string) (string, error) {
	switch method {
	case "sha256":
		hash := sha256.Sum256([]byte(password))
		return fmt.Sprintf("%x", hash), nil
	case "bcrypt":
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}
		return string(hashedPassword), nil
	default:
		return "", fmt.Errorf("unsupported encryption method")
	}
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	var req EncryptRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	encryptedPassword, err := encryptPassword(req.Password, req.EncryptionMethod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := EncryptResponse{EncryptedPassword: encryptedPassword}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/encrypt", encryptHandler)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
