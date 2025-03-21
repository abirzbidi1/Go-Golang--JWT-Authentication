package models

import (
	"database/sql"
	"errors"
	"go-jwt-api/config"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Insérer un utilisateur
func CreateUser(username, password string) error {
	_, err := config.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}

// Récupérer un utilisateur par username
func GetUserByUsername(username string) (*User, error) {
	row := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("utilisateur non trouvé")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
