package methods

import (
	"database/sql"
	"fmt"
)

// struct for GetBillItemData method
type BillItem struct {
	MedicineName string  `json:"medicineName"`
	Brand        string  `json:"brand"`
	Quantity     int     `json:"quantity"`
	UnitPrice    float64 `json:"unitPrice"`
}

// GetBillItemData gives the data selected
func GetBillItemData(db *sql.DB, medicineName string) (BillItem, error) {
	query := `
		SELECT mm.medicine_name, mm.brand, ms.quantity, ms.unit_price
		FROM medapp_medicine_master mm
		JOIN medapp_stock ms ON mm.medicine_master_id = ms.medicine_master_id
		WHERE mm.medicine_name = ?`

	var item BillItem
	err := db.QueryRow(query, medicineName).Scan(&item.MedicineName, &item.Brand, &item.Quantity, &item.UnitPrice)
	if err != nil {
		return item, fmt.Errorf("AGetMed006: %v", err)
	}

	return item, nil
}
