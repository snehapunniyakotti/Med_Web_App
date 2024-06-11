package api

import (
	"command/db"
	"command/methods"
	"command/structure"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// medicineDetailsResp defines the structure of the response body for fetching medicine details.
type medicineDetailsResp struct {
	MedicineName string                   `json:"medicineName"`
	Brand        string                   `json:"brand"`
	Quantity     string                   `json:"quantity"`
	UnitPrice    int                      `json:"unitPrice"`
	MedDetails   *methods.MedicineDetails `json:"medDetails"`
	MResp        structure.MasterResp     `json:"resp"`
}

// medicineDetailsReq defines the structure of the request body for fetching medicine details.
type medicineDetailsReq struct {
	MedicineName string `json:"medicineName"`
}

// FetchMedicineDetails handles the request to fetch details of a specific medicine.
func FetchMedicineDetails(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")

	log.Println("FetchMedicineDetails (+)")

	// Handle only GET requests.
	if r.Method == "GET" {
		var lreq medicineDetailsReq

		// Extract the medicine name from the query parameters.
		lreq.MedicineName = r.URL.Query().Get("medicineName")

		// Initialize the response structure and set default response status to success.
		var lresp medicineDetailsResp
		lresp.MResp.Status = "S"

		// Connect to the local database.
		ldb, lerr := db.LocalDBConnect()
		if lerr != nil {
			// Log the error and update response status and error message if connection fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "MedDetails001: " + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Ensure the database connection is closed after use.
			defer ldb.Close()

			// Call GetMedicineDetails method to fetch medicine details
			ldetails, lerr := methods.GetMedicineDetails(ldb, lreq.MedicineName)
			if lerr != nil {
				// Log the error and update response status and error message if fetching details fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "MedDetails002: " + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Update the response with fetched medicine details.
				lresp.MedDetails = ldetails
				lresp.MedicineName = ldetails.MedicineName
				lresp.Brand = ldetails.Brand
				lresp.Quantity = ldetails.Quantity
				lresp.UnitPrice = ldetails.UnitPrice
			}
		}

		// Marshal the response structure to JSON.
		ldata, lerr := json.Marshal(lresp)
		if lerr != nil {
			// Log the error if JSON marshaling fails and return internal server error.
			log.Println("JSON marshal error: MedDetails003:", lerr)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header and write the response data with appropriate status code.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(ldata))
		log.Println("FetchMedicineDetails (-)")
	}
}
