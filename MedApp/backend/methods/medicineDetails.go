package methods

import (
	"database/sql"
	"log"
)

type MedicineDetails struct {
	MedicineName string `json:"medicineName"`
	Brand        string `json:"brand"`
	Quantity     string `json:"quantity"`
	UnitPrice    int    `json:"unitPrice"`
}

// GetMedicineDetails retrieves the brand and unit price based on the provided medicine name
func GetMedicineDetails(db *sql.DB, medicineName string) (*MedicineDetails, error) {
	var details MedicineDetails

	query := `
		SELECT mmm.medicine_name, mmm.brand, ms.quantity, ms.unit_price
		FROM medapp_medicine_master mmm
		JOIN medapp_stock ms ON mmm.medicine_master_id = ms.medicine_master_id
		WHERE mmm.medicine_name = ?
	`

	err := db.QueryRow(query, medicineName).Scan(&details.MedicineName, &details.Brand, &details.Quantity, &details.UnitPrice)
	if err != nil {
		log.Println("Error fetching medicine details:", err)
		return nil, err
	}

	return &details, nil
}
