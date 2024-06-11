package methods

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// CheckUserId checks if a user ID exists in the database and returns the corresponding login ID.
func CheckUserId(db *sql.DB, userId string) (int, error) {
	log.Println(" checkUserId (+)")
	var lloginId int
	query := "SELECT login_id FROM sneha.medapp_login WHERE user_id = ?"
	lerr := db.QueryRow(query, userId).Scan(&lloginId)
	if lerr != nil {
		log.Println(" checkUserId (-)")
		return 0, lerr
	}
	log.Println(" checkuserId (-)")
	return lloginId, nil
}

// CheckPassword checks if the provided password matches the stored password for a given user ID.
func CheckPassword(db *sql.DB, userId string, password string) bool {
	log.Println(" checkPassword (+)")
	var lstoredPassword string
	query := "SELECT password FROM sneha.medapp_login WHERE user_id = ?"
	lerr := db.QueryRow(query, userId).Scan(&lstoredPassword)
	if lerr != nil {
		log.Println(lerr)
		log.Println(" checkpassword (-)")
		return false
	}
	log.Println(" checkpassword (-)")
	return lstoredPassword == password
}

// GetUserRole retrieves the role of a user based on their login ID.
func GetUserRole(db *sql.DB, loginId int) (string, error) {
	log.Println(" getUserRole (+)")
	var lrole string
	query := "SELECT role FROM sneha.medapp_login WHERE login_id = ?"
	lerr := db.QueryRow(query, loginId).Scan(&lrole)
	if lerr != nil {
		log.Println(" getUserRole (-)")
		return "", lerr
	}
	log.Println(" getUserRole (-)")
	return lrole, nil
}

// PostLoginHistory logs a user's login history, including successful or failed attempts.
func PostLoginHistory(db *sql.DB, loginId int, success bool) (int64, error) {
	log.Println(" postLoginHistory (+)")
	if !success {
		log.Println("  postLoginHistory (-)")
		return 0, fmt.Errorf("login failed for user %d", loginId)
	}

	lquery := "INSERT INTO sneha.medapp_login_history(login_id, login_date, login_time,created_by,created_date) VALUES (?, ?, ?, ?, ?)"
	lresult, lerr := db.Exec(lquery, loginId, time.Now().Format("2006-01-02"), time.Now().Format("15:04:05"), "sneha", time.Now().Format("2006-01-02"))
	if lerr != nil {
		log.Println(" error postLoginHistory (-)")
		return 0, lerr
	}
	loginHistoryId, lerr := lresult.LastInsertId()
	if lerr != nil {
		return 0, lerr
	}
	log.Println(" nil postLoginHistory (-)")
	return loginHistoryId, nil
}
