package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Global defines a structure for global responses containing a message, status, and error message.
type Global struct {
	Message string `json:"g_message"`
	Status  string `json:"g_status"`
	Errmsg  string `json:"g_errmsg"`
}

// LocalDBConnect establishes a connection to the local MySQL database.
func LocalDBConnect() (*sql.DB, error) {
	log.Println("Local DBConnect  (+)")

	// Construct the connection string using MySQL driver format.
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "192.168.2.5", 3306, "sneha")

	// Open a connection to the MySQL database using the constructed connection string.
	db, err := sql.Open("mysql", connString)
	if err != nil {
		// Log an error if opening the database connection fails and return the error.
		log.Println("Open Database Connection failed: ", err.Error())
		return db, err
	}

	// Log successful database connection establishment.
	log.Println("Local DBConnect  (-)")
	return db, nil
}
