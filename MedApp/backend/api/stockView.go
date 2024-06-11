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

// stockViewResp struct defines the response structure for the GetStockView API.
type stockViewResp struct {
	StockArr []methods.StockView  `json:"stockArr"`
	MResp    structure.MasterResp `json:"resp"`
}

func GetStockView(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")

	log.Println("StockView (+)")

	// Handle only GET requests.
	if r.Method == "GET" {
		var lresp stockViewResp

		// Set default response status to success.
		lresp.MResp.Status = "S"

		// Connect to the local database.
		ldb, lerr := db.LocalDBConnect()
		if lerr != nil {
			// Log the error and update response status and error message if connection fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "ASV001: " + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Ensure the database connection is closed after use.
			defer ldb.Close()

			// Method 1: Fetch all stock data.
			stockArr, lerr := methods.FetchAllStock(ldb)
			if lerr != nil {
				// Log the error and update response status and error message if fetching data fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "ASV002: " + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Set the fetched data to the response structure.
				lresp.StockArr = stockArr
			}
		}

		// Marshal the response structure to JSON format.
		ldata, lerr := json.Marshal(lresp)
		if lerr != nil {
			// Send an error message if marshalling fails.
			fmt.Fprintf(w, "Error taking data"+"ASV003: "+lerr.Error())
		} else {
			// Set response content type to JSON and write the response with status OK.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(ldata))
		}
		log.Println("StockView (-)")
	}
}
