package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	dsn := "root:password@tcp(127.0.0.1:3306)/go_jwt_api"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Erreur de connexion :", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Impossible de se connecter à la base de données :", err)
	}

	fmt.Println("Connexion réussie à MySQL !")

	// Création de la table users si elle n'existe pas
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(100) NOT NULL UNIQUE,
        password TEXT NOT NULL
    )`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Erreur lors de la création de la table :", err)
	}
	fmt.Println("Table 'users' prête !")
}
package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	dsn := "root:password@tcp(127.0.0.1:3306)/go_jwt_api"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Erreur de connexion :", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Impossible de se connecter à la base de données :", err)
	}

	fmt.Println("Connexion réussie à MySQL !")

	// Création de la table users si elle n'existe pas
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(100) NOT NULL UNIQUE,
        password TEXT NOT NULL
    )`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Erreur lors de la création de la table :", err)
	}
	fmt.Println("Table 'users' prête !")
}
