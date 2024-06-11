package methods

import (
	"database/sql"
	"log"
)

type LoginHistory struct {
	LoginId    int            `json:"loginId"`
	UserId     string         `json:"userId"`
	LoginDate  string         `json:"loginDate"`
	LoginTime  string         `json:"loginTime"`
	LogoutDate sql.NullString `json:"logoutDate"`
	LogoutTime sql.NullString `json:"logoutTime"`
}

// FetchLoginHistory fetches login history from the medapp_login_history table
func FetchLoginHistory(ldb *sql.DB) ([]LoginHistory, error) {
	query := `SELECT lh.login_id,ml.user_id , lh.login_date, lh.login_time, lh.logout_date, lh.logout_time 
	FROM sneha.medapp_login_history lh
	join sneha.medapp_login ml 
	on lh.login_id = ml.login_id 
	order by lh.login_date`
	rows, err := ldb.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var loginHistories []LoginHistory

	for rows.Next() {
		var history LoginHistory
		err = rows.Scan(&history.LoginId, &history.UserId, &history.LoginDate, &history.LoginTime, &history.LogoutDate, &history.LogoutTime)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		loginHistories = append(loginHistories, history)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return loginHistories, nil
}
