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

// updateStockResp defines the structure of the response body for the UpdateStock API.
type updateStockResp struct {
	Msg   string               `json:"msg"`
	MResp structure.MasterResp `json:"resp"`
}

// updateStockReq defines the structure of the request body for the UpdateStock API.
type updateStockReq struct {
	MedicineName string `json:"medicineName"`
	Quantity     int    `json:"quantity"`
	UnitPrice    int    `json:"unitPrice"`
	Brand        string `json:"brand"`
}

// This api is used to update the login history while clicking logout in ui. request data is loginHistoryid
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")

	log.Println("PutStock +")

	// Handle only PUT requests.
	if r.Method == "PUT" {
		var ldetails updateStockReq
		var lresp updateStockResp

		// Set default response status to success.
		lresp.MResp.Status = "S"
		// Read the request body.
		lbody, lerr := ioutil.ReadAll(r.Body)
		if lerr != nil {
			// Log the error and update response status and error message if reading the body fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "APS001" + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Unmarshal the request body into ldetails.
			lerr = json.Unmarshal(lbody, &ldetails)
			if lerr != nil {
				// Log the error and update response status and error message if unmarshaling fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "APS002" + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Connect to the local database.
				ldb, lerr := db.LocalDBConnect()
				if lerr != nil {
					// Log the error and update response status and error message if connection fails.
					log.Println(lerr)
					lresp.MResp.ErrMsg = "APS003" + lerr.Error()
					lresp.MResp.Status = "E"
				} else {
					// Ensure the database connection is closed after use.
					defer ldb.Close()

					// log.Print(ldetails)
					// Methods 1 - Call UpdateStockDetails method from methods package to update the stock details.
					lerr := methods.UpdateStockDetails(ldb, ldetails.MedicineName, ldetails.Quantity, ldetails.UnitPrice)
					if lerr != nil {
						// Log the error and update response status and error message if updating stock details fails.
						log.Println(lerr)
						lresp.MResp.ErrMsg = "APS004" + lerr.Error()
						lresp.MResp.Status = "E"
					} else {
						// Update the response message on successful update.
						lresp.Msg = "updated Successfully"
						log.Println("updated Successfully")
					}
				}
			}
		}
		// here we marshal the lresp struct to JSON
		data, lerr := json.Marshal(lresp)
		if lerr != nil {
			// Handle error in marshaling the response data.
			fmt.Fprintf(w, "Error taking data"+"ALogout005"+lerr.Error())
		} else {
			// Write the response data.
			fmt.Fprint(w, string(data))
		}
		log.Println(" PutStock (-)")
	}

}
