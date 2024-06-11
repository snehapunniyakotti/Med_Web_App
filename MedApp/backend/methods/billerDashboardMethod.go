package methods

import (
	"database/sql"
	"fmt"
	"log"
)

// BillerDashboard struct defines the structure to hold sales data for the current and previous day.
type BillerDashboard struct {
	CurrentDaySales  float64 `json:"current_day_sales"`
	PreviousDaySales float64 `json:"previous_day_sales"`
}

// GetSalesOfBiller fetches the sales data for today and yesterday for the given user from the database.
func GetSalesOfBiller(db *sql.DB, userId string) (BillerDashboard, error) {
	var dashboard BillerDashboard

	// Query to get today's sales.
	queryToday := `SELECT IFNULL(SUM(net_price), 0) AS total_sales
		FROM sneha.medapp_bill_master mbm
		JOIN sneha.medapp_login ml 
		ON mbm.login_id = ml.login_id 
		WHERE mbm.bill_date = CURDATE() AND ml.user_id = ?`
	err := db.QueryRow(queryToday, userId).Scan(&dashboard.CurrentDaySales)
	log.Print(&dashboard.CurrentDaySales)
	if err != nil {
		log.Println("Error fetching today's sales:", err)
		return dashboard, fmt.Errorf("error fetching today's sales: %v", err)
	}

	// Query to get yesterday's sales.
	queryYesterday := `SELECT IFNULL(SUM(net_price), 0) AS total_sales
		FROM sneha.medapp_bill_master mbm
		JOIN sneha.medapp_login ml 
		ON mbm.login_id = ml.login_id 
		WHERE mbm.bill_date = CURDATE()-1 AND ml.user_id = ?`
	err = db.QueryRow(queryYesterday, userId).Scan(&dashboard.PreviousDaySales)
	log.Print(&dashboard.PreviousDaySales)
	if err != nil {
		log.Println("Error fetching yesterday's sales:", err)
		return dashboard, fmt.Errorf("error fetching yesterday's sales: %v", err)
	}

	return dashboard, nil
}
