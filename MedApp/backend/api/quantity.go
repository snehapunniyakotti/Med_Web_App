package api

import (
	"command/db"
	"command/methods"
	"command/structure"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// billDetailsReq defines the structure of the request body for checking quantity.
type billDetailsReq struct {
	MedicineName string `json:"medicineName"`
	Quantity     int    `json:"quantity"`
}

// billDetailsResp defines the structure of the response body for the quantity check API.
type billDetailsResp struct {
	OneBillItem methods.BillItem     `json:"oneBillItem"`
	MResp       structure.MasterResp `json:"resp"`
}

// CheckQuantityApi handles the request for checking the quantity of a specified medicine.
func CheckQuantityApi(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	// Handle only POST requests.
	if r.Method == "POST" {

		log.Println("Quantity (+)")

		var lreq billDetailsReq
		var lresp billDetailsResp

		// Set default response status to success.
		lresp.MResp.Status = "S"

		// Read the request body.
		lbody, lerr := ioutil.ReadAll(r.Body)
		if lerr != nil {
			// Log the error and update response status and error message if reading the body fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "AQ001: " + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Unmarshal the request body into lreq.
			lerr = json.Unmarshal(lbody, &lreq)
			if lerr != nil {
				// Log the error and update response status and error message if unmarshaling fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "AQ002: " + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Connect to the local database.
				ldb, lerr := db.LocalDBConnect()
				if lerr != nil {
					// Log the error and update response status and error message if connection fails.
					log.Println(lerr)
					lresp.MResp.ErrMsg = "AQ003: " + lerr.Error()
					lresp.MResp.Status = "E"
				} else {
					// Ensure the database connection is closed after use.
					defer ldb.Close()

					// Fetch the bill item data for the specified medicine.
					billItem, lerr := methods.GetBillItemData(ldb, lreq.MedicineName)
					if lerr != nil {
						// Log the error and update response status and error message if fetching data fails.
						log.Println(lerr)
						lresp.MResp.ErrMsg = "AQ004: " +  lerr.Error()
						lresp.MResp.Status = "E"
					} else {
						// Check if the available quantity is sufficient.
						if billItem.Quantity >= lreq.Quantity {
							// Assign the fetched bill item data to the response.
							lresp.OneBillItem = billItem
						} else {
							// Set response status and error message for insufficient quantity.
							lresp.MResp.ErrMsg =  "AQ005: " +  "Insufficient quantity"
							lresp.MResp.Status = "E"
						}
					}
				}
			}
		}

		// Marshal the response structure to JSON.
		ldata, lerr := json.Marshal(lresp)
		if lerr != nil {
			// Handle error in marshaling the response data.
			fmt.Fprintf(w, "Error taking data"+"AQ005: "+lerr.Error())
		} else {
			// Write the response data.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(ldata))
		}
		log.Println("Quantity (-)")
	}
}
