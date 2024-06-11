package methods

import (
	"database/sql"
	"log"
)

// UpdateStockDetails updates the present quantity or unit price in the stock table
func UpdateStockDetails(ldb *sql.DB, medicineName string, quantity int, unitPrice int) error {
	// Check if the medicine exists in the medicine_master table and get the medicine_master_id
	var medicineMasterID int
	err := ldb.QueryRow("SELECT medicine_master_id FROM sneha.medapp_medicine_master WHERE medicine_name = ?", medicineName).Scan(&medicineMasterID)
	if err != nil {
		log.Println(err)
		return err
	}

	// Update the quantity and unit price in the stock table using the medicine_master_id
	_, err = ldb.Exec("UPDATE sneha.medapp_stock SET quantity = quantity + ?, unit_price = ? WHERE medicine_master_id = ?", quantity, unitPrice, medicineMasterID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
