package methods

import (
	"database/sql"
	"log"
)

type ManagerDashboard struct {
	TotalSales            float64 `json:"totalSales"`
	CurrentInventoryValue float64 `json:"curInventry"`
}

// GetTotalSalesAndCurInventryValue gives total sales and current inventory value
func GetTotalSalesAndCurInventryValue(db *sql.DB) (ManagerDashboard, error) {
	var dashboard ManagerDashboard

	// Query for total sales
	query1 := `SELECT IFNULL(SUM(net_price), 0) AS total_sales
	           FROM sneha.medapp_bill_master
	           WHERE bill_date = CURDATE();`
	err := db.QueryRow(query1).Scan(&dashboard.TotalSales)
	if err != nil {
		log.Println("Error fetching total sales:", err)
		return dashboard, err
	}

	// Query for current inventory value
	query2 := `SELECT IFNULL(SUM(ms.quantity * ms.unit_price), 0) AS current_inventory_value
	           FROM sneha.medapp_stock ms;`
	err = db.QueryRow(query2).Scan(&dashboard.CurrentInventoryValue)
	if err != nil {
		log.Println("Error fetching current inventory value:", err)
		return dashboard, err
	}

	return dashboard, nil
}
