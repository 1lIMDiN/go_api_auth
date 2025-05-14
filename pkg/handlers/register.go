package handlers

import (
	"encoding/json"
	"net/http"

	"auth/pkg/database"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user database.User
	// Получаем логин и пароль пользователя
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// Хэшируем пароль 'bcrypt'
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "failed to hashed password", http.StatusInternalServerError)
		return
	}

	// Добавляем полученные данные в БД
	_, err = database.DB.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", user.Username, string(hashedPassword))
	if err != nil {
		http.Error(w, "failed to add user id DB", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}