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

// HistoryResp struct defines the response structure for the GetHistory API.
type HistoryResp struct {
	HistoryArr []methods.LoginHistory `json:"historyArr"`
	MResp      structure.MasterResp   `json:"resp"`
}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")

	log.Println("GetHistory (+)")

	// Handle only GET requests.
	if r.Method == "GET" {
		var lresp HistoryResp

		// Set default response status to success.
		lresp.MResp.Status = "S"

		// Connect to the local database.
		ldb, lerr := db.LocalDBConnect()
		if lerr != nil {
			// Log the error and update response status and error message if connection fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "AGetHistory001: " + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Ensure the database connection is closed after use.
			defer ldb.Close()

			// Method 1: Fetch the login history data.
			historyArr, lerr := methods.FetchLoginHistory(ldb)
			if lerr != nil {
				// Log the error and update response status and error message if fetching data fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "AGetHistory002: " + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Set the fetched data to the response structure.
				lresp.HistoryArr = historyArr
			}
		}

		// Marshal the response structure to JSON format.
		ldata, lerr := json.Marshal(lresp)
		if lerr != nil {
			// Send an error message if marshalling fails.
			fmt.Fprintf(w, "Error taking data"+"AGetHistory003: "+lerr.Error())
		} else {
			// Set response content type to JSON and write the response with status OK.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(ldata))
		}
		log.Println("GetHistory (-)")
	}
}
