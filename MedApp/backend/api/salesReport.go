// package api

// import (
// 	"command/db"
// 	"command/methods"
// 	"command/structure"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type salesReportReq struct {
// 	FromDate string `json:"fromDate"`
// 	ToDate   string `json:"toDate"`
// }

// type salesReportResp struct {
// 	SalesArr []methods.SalesReport `json:"salesArr"`
// 	MResp    structure.MasterResp  `json:"resp"`
// }

// func GetSalesReport(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")

// 	log.Println("GetSalesReport (+)")

// 	if r.Method == "GET" {
// 		var lresp salesReportResp
// 		var lreq salesReportReq

// 		err := json.NewDecoder(r.Body).Decode(&lreq)
// 		if err != nil {
// 			log.Println("Error decoding request:", err)
// 			lresp.MResp.ErrMsg = "Invalid request format"
// 			lresp.MResp.Status = "E"
// 		} else {
// 			lresp.MResp.Status = "S"

// 			ldb, lerr := db.LocalDBConnect()
// 			if lerr != nil {
// 				log.Println(lerr)
// 				lresp.MResp.ErrMsg = "AGetMed001: " + lerr.Error()
// 				lresp.MResp.Status = "E"
// 			} else {
// 				defer ldb.Close()

// 				// method
// 				// Call FetchSales method from methods package
// 				sales, err := methods.FetchSales(ldb, lreq.FromDate, lreq.ToDate)
// 				if err != nil {
// 					log.Println("Error fetching sales data:", err)
// 					lresp.MResp.ErrMsg = "Error fetching sales data"
// 					lresp.MResp.Status = "E"
// 				} else {
// 					lresp.SalesArr = sales
// 				}
// 			}
// 		}

// 		ldata, lerr := json.Marshal(lresp)
// 		if lerr != nil {
// 			fmt.Fprintf(w, "Error taking data"+"AGetMed003: "+lerr.Error())
// 		} else {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprint(w, string(ldata))
// 		}
// 		log.Println("GetSalesReport (-)")
// 	}
// }

// func sendResponse(w http.ResponseWriter, resp salesReportResp, statusCode int) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		log.Println("Error encoding response:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 	}
// }

package api

import (
	"command/db"
	"command/methods"
	"command/structure"
	"encoding/json"

	// "encoding/json"
	"log"
	"net/http"
)

// salesReportReq defines the structure of the request parameters for the sales report API.
type salesReportReq struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

// salesReportResp defines the structure of the response for the sales report API.
type salesReportResp struct {
	SalesArr []methods.SalesReport `json:"salesArr"`
	MResp    structure.MasterResp  `json:"resp"`
}

// GetSalesReport handles the request for fetching the sales report within a date range.
func GetSalesReport(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")

	log.Println("SalesReport (+)")

	// Handle only GET requests
	if r.Method == "GET" {
		var lresp salesReportResp

		// Get the 'fromDate' and 'toDate' query parameters from the request URL.
		lfromDate := r.URL.Query().Get("fromDate")
		ltoDate := r.URL.Query().Get("toDate")

		// Check if both 'fromDate' and 'toDate' are provided.
		if lfromDate == "" || ltoDate == "" {
			log.Println("Missing required parameters")
			lresp.MResp.ErrMsg = "Missing required parameters"
			lresp.MResp.Status = "E"
			// sendResponse(w, lresp, http.StatusBadRequest)
			return
		}

		// Set default response status to success.
		lresp.MResp.Status = "S"

		// Connect to the local database.
		ldb, lerr := db.LocalDBConnect()
		if lerr != nil {
			// Log the error and update response status and error message if connection fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "ASR001: " + lerr.Error()
			lresp.MResp.Status = "E"
			// sendResponse(w, lresp, http.StatusInternalServerError)
			return
		}
		// Ensure the database connection is closed after use.
		defer ldb.Close()

		// Method 1 -Fetch sales data within the specified date range.
		sales, lerr := methods.FetchSales(ldb, lfromDate, ltoDate)
		if lerr != nil {
			// Log the error and update response status and error message if fetching sales data fails.
			log.Println("Error fetching sales data:", lerr)
			lresp.MResp.ErrMsg = "Error fetching sales data : ASR002: " + lerr.Error()
			lresp.MResp.Status = "E"
			// sendResponse(w, lresp, http.StatusInternalServerError)
			return
		} else {
			// Assign the fetched sales data to the response.
			lresp.SalesArr = sales
		}

		log.Println("SalesReport (-)")

		sendResponse(w, lresp, http.StatusOK)
	}
}

func sendResponse(w http.ResponseWriter, resp salesReportResp, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
