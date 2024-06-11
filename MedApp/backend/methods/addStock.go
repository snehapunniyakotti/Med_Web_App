package methods

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// CheckStock checks if a particular stock item exists in the database.
func CheckStock(ldb *sql.DB, lmedicineName, lbrand string) (int, error) {
	log.Println(" CheckStock (+)")
	var lmedicineMasterID int
	lquery := `SELECT medicine_master_id FROM sneha.medapp_medicine_master WHERE medicine_name = ? AND brand = ?`
	lerr := ldb.QueryRow(lquery, lmedicineName, lbrand).Scan(&lmedicineMasterID)
	if lerr != nil && lerr != sql.ErrNoRows {
		log.Println(" CheckStock err(-)")
		return 0, lerr
	}
	log.Println(" CheckStock success(-)")
	return lmedicineMasterID, nil
}

// AddMedicineMaster adds a new medicine master record to the database.
func AddMedicineMaster(ldb *sql.DB, lmedicineName, lbrand string) (int, error) {
	log.Println(" AddMedicineMaster (+)")
	lname := "sneha"
	ldate := time.Now().Format("2006-01-02")
	lquery := `INSERT INTO sneha.medapp_medicine_master (medicine_name, brand, created_by, created_date, updated_by, updated_date) VALUES (?, ?, ?, ?, ?, ?)`
	lresult, lerr := ldb.Exec(lquery, lmedicineName, lbrand, lname, ldate, lname, ldate)
	if lerr != nil {
		log.Println(" AddMedicineMaster err(-)")
		return 0, lerr
	}
	lmedicineMasterID, lerr := lresult.LastInsertId()
	if lerr != nil {
		log.Println(" AddMedicineMaster err lastindex(-)")
		return 0, lerr
	}
	log.Println(" AddMedicineMaster (-)")
	return int(lmedicineMasterID), nil
}

// AddStock adds a new stock entry for a given medicine master ID.
func AddStock(ldb *sql.DB, lmedicineMasterID int) error {
	log.Println(" AddStock (+)")
	lname := "sneha"
	ldate := time.Now().Format("2006-01-02")
	lquery := `INSERT INTO sneha.medapp_stock (medicine_master_id, unit_price, quantity, created_by, created_date) VALUES (?, ?, ?, ?, ?)`
	_, lerr := ldb.Exec(lquery, lmedicineMasterID, 1, 0, lname, ldate)
	if lerr != nil {
		log.Println(" AddStock err(-)")
		return fmt.Errorf("failed to insert stock: %v", lerr)
	}
	log.Println(" AddStock success(-)")
	return nil
}
