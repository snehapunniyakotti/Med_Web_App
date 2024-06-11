package methods

import (
	"database/sql"
	"log"
)

// struct for  FetchAllStock api
type StockView struct {
	MedicineName string `json:"medicineName"`
	Brand        string `json:"brand"`
	Quantity     int    `json:"quantity"`
	UnitPrice    int    `json:"unitPrice"`
}

// FetchAllStock gives the data selected
func FetchAllStock(ldb *sql.DB) ([]StockView, error) {
	query := `
		SELECT mmm.medicine_name, mmm.brand, ms.quantity, ms.unit_price
		FROM sneha.medapp_medicine_master mmm
		JOIN sneha.medapp_stock ms ON mmm.medicine_master_id = ms.medicine_master_id;
	`

	rows, err := ldb.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var stockViews []StockView

	for rows.Next() {
		var stockView StockView
		err := rows.Scan(&stockView.MedicineName, &stockView.Brand, &stockView.Quantity, &stockView.UnitPrice)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		stockViews = append(stockViews, stockView)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return stockViews, nil
}
