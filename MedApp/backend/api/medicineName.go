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

// medNameResp defines the structure of the response body for the GetMedicineName API.
type medNameResp struct {
	MedArr []methods.MedNameArr `json:"medArr"`
	MResp  structure.MasterResp `json:"resp"`
}

// GetMedicineName handles the request to fetch medicine names.
func GetMedicineName(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")

	log.Println("MedicineName (+)")

	// Handle only GET requests.
	if r.Method == "GET" {
		var lresp medNameResp

		// Set default response status to success.
		lresp.MResp.Status = "S"

		// Connect to the local database.
		ldb, lerr := db.LocalDBConnect()
		if lerr != nil {
			// Log the error and update response status and error message if connection fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "AGetMed001: " + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Ensure the database connection is closed after use.
			defer ldb.Close()

			// Method 1 -Call FetchMedicines method from methods package to fetch medicine names.
			lmeds, lerr := methods.FetchMedicines(ldb)
			if lerr != nil {
				// Log the error and update response status and error message if fetching medicines fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "AGetMed002: " + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Update the response with fetched medicine names.
				lresp.MedArr = lmeds
			}
		}

		// Marshal the response structure to JSON.
		ldata, lerr := json.Marshal(lresp)
		if lerr != nil {
			// Handle error in marshaling the response data.
			fmt.Fprintf(w, "Error taking data"+"AGetMed003: "+lerr.Error())
		} else {
			// Write the response data with appropriate headers and status code.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(ldata))
		}
		log.Println("MedicineName (-)")
	}
}
