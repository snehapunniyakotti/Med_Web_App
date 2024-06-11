package methods

import (
	"database/sql"
	"log"
	"time"
)

// AddUser adds a new user to the database
func AddUser(db *sql.DB, userId, password, role string) error {
	// Prepare the SQL statement to insert a new user into the database
	lquery := "INSERT INTO sneha.medapp_login (user_id, password, role,created_by,created_date) VALUES (? ,? ,?, ?, ?)"

	// Execute the SQL statement with the user data
	_, lerr := db.Exec(lquery, userId, password, role, "sneha", time.Now().Format("2006-01-02"))
	if lerr != nil {
		log.Println("Error executing SQL statement:", lerr)
		return lerr
	}

	log.Println("User added successfully to the database")
	return nil
}

// IsUserIdExists checks if a user ID already exists in the database.
func IsUserIdExists(db *sql.DB, userId string) bool {
	lquery := "SELECT COUNT(*) FROM sneha.medapp_login WHERE user_id = ?"

	var count int
	lerr := db.QueryRow(lquery, userId).Scan(&count)
	if lerr != nil {
		log.Println("Error checking user ID:", lerr)
		return false // Assume false on error
	}

	return count > 0
}
