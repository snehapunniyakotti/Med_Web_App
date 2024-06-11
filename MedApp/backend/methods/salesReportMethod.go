package methods

import (
	"database/sql"
)

// struct for FetchSales method
type SalesReport struct {
	BillNo       int    `json:"billNo"`
	BillDate     string `json:"billDate"`
	MedicineName string `json:"medicineName"`
	Quantity     int    `json:"quantity"`
	Amount       int    `json:"amount"`
}

// FetchSales gives data selected
func FetchSales(ldb *sql.DB, fromDate, toDate string) ([]SalesReport, error) {
	var sales []SalesReport

	// Construct the SQL query using placeholders for fromDate and toDate
	query := `
		SELECT mbm.bill_no, mbm.bill_date, mmm.medicine_name, mbd.quantity, mbd.amount
		FROM sneha.medapp_bill_details mbd
		JOIN sneha.medapp_bill_master mbm ON mbd.bill_no = mbm.bill_no
		JOIN sneha.medapp_medicine_master mmm ON mbd.medicine_master_id = mmm.medicine_master_id
		WHERE mbm.bill_date BETWEEN ? AND ?
	`

	// Execute the query using the provided database connection and fromDate/toDate values
	rows, err := ldb.Query(query, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the query results and populate the sales slice
	for rows.Next() {
		var salesReport SalesReport
		err := rows.Scan(&salesReport.BillNo, &salesReport.BillDate, &salesReport.MedicineName, &salesReport.Quantity, &salesReport.Amount)
		if err != nil {
			return nil, err
		}
		sales = append(sales, salesReport)
	}

	return sales, nil
}
