package methods

import "database/sql"

type MedNameArr struct {
	MedName string `json:"medName"`
}

//  FetchMedicines  This method will get the list of medicine available in medapp_medicine_master table in the database having fields medicine_name
func FetchMedicines(ldb *sql.DB) ([]MedNameArr, error) {
	query := "SELECT medicine_name FROM sneha.medapp_medicine_master"
	rows, err := ldb.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meds []MedNameArr
	for rows.Next() {
		var med MedNameArr
		if err := rows.Scan(&med.MedName); err != nil {
			return nil, err
		}
		meds = append(meds, med)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return meds, nil
}
