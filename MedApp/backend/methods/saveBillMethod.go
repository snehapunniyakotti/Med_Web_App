package methods

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// PostBillMaster inserts data into the medapp_bill_master table
func PostBillMaster(db *sql.DB, billNo int, billAmount int, billGst float32, netPrice int, loginId int) error {
	Date := time.Now().Format("2006-01-02")
	name := "sneha"
	query := `INSERT INTO sneha.medapp_bill_master (bill_no, bill_amount, bill_gst, net_price, login_id, bill_date, created_by, created_date ,updated_by, updated_date) 
	          VALUES (?, ?, ?, ?, ? ,? ,? ,? ,? ,?)`
	_, err := db.Exec(query, billNo, billAmount, billGst, netPrice, loginId, Date, name, Date, name, Date)
	if err != nil {
		return fmt.Errorf("error inserting into medapp_bill_master: %v", err)
	}
	return nil
}

// PostBillDetails inserts data into the medapp_bill_details table
func PostBillDetails(db *sql.DB, billNo int, medicineMasterId int, quantity int, unitPrice int, amount int) error {
	Date := time.Now().Format("2006-01-02")
	name := "sneha"
	query := `INSERT INTO sneha.medapp_bill_details (bill_no, medicine_master_id, quantity, unit_price, amount, created_by, created_date ,updated_by, updated_date) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, billNo, medicineMasterId, quantity, unitPrice, amount, name, Date, name, Date)
	if err != nil {
		return fmt.Errorf("error inserting into medapp_bill_details: %v", err)
	}
	return nil
}

// GetMedicineMasterId retrieves the medicine_master_id based on the medicine name
func GetMedicineMasterId(db *sql.DB, medicineName string) (int, error) {
	var medicineMasterId int

	query := `SELECT medicine_master_id FROM sneha.medapp_medicine_master WHERE medicine_name = ?`
	err := db.QueryRow(query, medicineName).Scan(&medicineMasterId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("medicine not found")
		}
		return 0, fmt.Errorf("error retrieving medicine_master_id: %v", err)
	}
	return medicineMasterId, nil
}

// GetLoginId fetches the login_id from the medapp_login table based on user_id
func GetLoginId(db *sql.DB, userId string) (int, error) {
	var loginId int
	query := "SELECT login_id FROM sneha.medapp_login WHERE user_id = ?"
	err := db.QueryRow(query, userId).Scan(&loginId)
	if err != nil {
		return 0, err
	}
	return loginId, nil
}

// UpdateStock subtracts quantity from existing stock for a given medicine_master_id
func UpdateStock(db *sql.DB, medicineMasterId int, quantity int) error {
	// Fetch existing quantity
	var existingQuantity int
	query := "SELECT quantity FROM sneha.medapp_stock WHERE medicine_master_id = ?"
	err := db.QueryRow(query, medicineMasterId).Scan(&existingQuantity)
	if err != nil {
		return err
	}

	// Calculate new quantity after subtraction
	newQuantity := existingQuantity - quantity
	if newQuantity < 0 {
		return fmt.Errorf("insufficient stock")
	}

	// Update stock in database
	updateQuery := "UPDATE sneha.medapp_stock SET quantity = ? WHERE medicine_master_id = ?"
	_, err = db.Exec(updateQuery, newQuantity, medicineMasterId)
	if err != nil {
		return err
	}

	return nil
}
