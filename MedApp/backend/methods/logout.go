package methods

import (
	"database/sql"
	"fmt"
	"time"
)

// UpdateLoginHistory update logout date and time
func UpdateLoginHistory(db *sql.DB, loginHistoryId int) error {
	lquery := `UPDATE medapp_login_history 
			  SET logout_time = ?, logout_date = ?,updated_by=?,updated_date=?
			  WHERE login_history_id = ?`
	llogoutTime := time.Now().Format("15:04:05")   // assuming time format
	llogoutDate := time.Now().Format("2006-01-02") // assuming date format

	_, lerr := db.Exec(lquery, llogoutTime, llogoutDate, "sneha", llogoutDate, loginHistoryId)
	if lerr != nil {
		return fmt.Errorf("failed to execute update: %v", lerr)
	}

	return nil
}
