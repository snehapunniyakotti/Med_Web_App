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

// BillerDashboardresponse struct defines the request structure for the BillerDashboard API.
type BillerDashboardrequest struct {
	UserId string `json:"userId"`
}

// BillerDashboardresponse struct defines the response structure for the BillerDashboard API.
type BillerDashboardresponse struct {
	Arr   methods.BillerDashboard `json:"arr"`
	MResp structure.MasterResp    `json:"resp"`
}

// BillerDashboard handles HTTP requests for the biller dashboard data.

func BillerDashboard(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	log.Println("BillerDashboard (+)")
	// Handle only GET requests.
	if r.Method == "POST" {
		var lresp BillerDashboardresponse
		// var lreq BillerDashboardrequest
		var l_data map[string]string

		// Set default response status to success.
		lresp.MResp.Status = "S"
		// Read the request body.
		lbody, lerr := ioutil.ReadAll(r.Body)
		if lerr != nil {
			// Log the error and update response status and error message if reading request body fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "ABD001" + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Unmarshal the request body into newUserReq struct.
			// log.Println("Received JSON data:", string(lbody))
			lerr = json.Unmarshal(lbody, &l_data)
			if lerr != nil {
				// Log the error and update response status and error message if unmarshaling fails.
				log.Println("ABD002", lerr)
				lresp.MResp.ErrMsg = "ABD002" + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Connect to the local database.
				ldb, lerr := db.LocalDBConnect()
				if lerr != nil {
					// Log the error and update response status and error message if connection fails.
					log.Println(lerr)
					lresp.MResp.ErrMsg = "ABD001: " + lerr.Error()
					lresp.MResp.Status = "E"
				} else {
					// Ensure the database connection is closed after use.
					defer ldb.Close()

					//method - Fetch sales data for the biller from the database.
					dashboard, lerr := methods.GetSalesOfBiller(ldb, l_data["userId"])
					if lerr != nil {
						// Log the error and update response status and error message if fetching data fails.
						log.Println(lerr)
						lresp.MResp.ErrMsg = "ABD002: " + lerr.Error()
						lresp.MResp.Status = "E"
					} else {
						// Set the fetched data to the response structure.
						lresp.Arr = dashboard
					}
				}
			}
		}

		// Marshal the response structure to JSON format.
		ldata, lerr := json.Marshal(lresp)
		if lerr != nil {
			// Send an error message if marshalling fails.
			fmt.Fprintf(w, "Error taking data"+"ABD003: "+lerr.Error())
		} else {
			// Set response content type to JSON and write the response with status OK.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(ldata))
		}
		log.Println("BillerDashboard (-)")
	}
}
