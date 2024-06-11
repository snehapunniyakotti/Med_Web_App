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

// addStockResp defines the structure of the response body for adding stock.
type addStockResp struct {
	Msg   string               `json:"msg"`
	MResp structure.MasterResp `json:"resp"`
}

// addStockReq defines the structure of the request body for adding stock.
type addStockReq struct {
	MedicineName string `json:"medicineName"`
	Brand        string `json:"Brand"`
}

// this API will get the request data userId and Password . it has four methods to validate login and give response data role,status,errMsg
func PostStockApi(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	log.Println("CheckLoginUserinsert  +")

	// Handle only POST requests.
	if r.Method == "POST" {
		var lstockdetails addStockReq
		var lresp addStockResp

		// Initialize the response structure and set default response status to success.
		lresp.MResp.Status = "S"
		// Read the request body.
		lbody, lerr := ioutil.ReadAll(r.Body)
		if lerr != nil {
			// Log the error and update response status and error message if reading request body fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "AAddStock001" + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Unmarshal the request body into addStockReq struct.
			lerr = json.Unmarshal(lbody, &lstockdetails)
			if lerr != nil {
				// Log the error and update response status and error message if unmarshaling fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "AAddStock002" + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Connect to the local database.
				ldb, lerr := db.LocalDBConnect()
				if lerr != nil {
					// Log the error and update response status and error message if database connection fails.
					log.Println(lerr)
					lresp.MResp.ErrMsg = "AAddStock003" + lerr.Error()
					lresp.MResp.Status = "E"
				} else {
					// Ensure the database connection is closed after use
					defer ldb.Close()
					// Method 1 - Check if the entered stock medicineName and brand are present in the medicine master table in the database.
					lmedicineMasterID, lerr := methods.CheckStock(ldb, lstockdetails.MedicineName, lstockdetails.Brand)
					if lerr != nil {
						// Log the error and update response status and error message if checking stock fails.
						log.Println(lerr)
						lresp.MResp.ErrMsg = "AAddStock004: " + lerr.Error()
						lresp.MResp.Status = "E"
					} else if lmedicineMasterID != 0 {
						// If stock is already present, update response status and error message accordingly.
						lresp.MResp.ErrMsg = "Stock already present"
						lresp.MResp.Status = "E"
					} else {
						// Method 2 - Add the new stock if not already present in the system.
						lmedicineMasterID, lerr = methods.AddMedicineMaster(ldb, lstockdetails.MedicineName, lstockdetails.Brand)
						if lerr != nil {
							// Log the error and update response status and error message if adding medicine master fails.
							log.Println(lerr)
							lresp.MResp.ErrMsg = "AAddStock005: " + lerr.Error()
							lresp.MResp.Status = "E"
						} else {
							// Method 3 -Insert data into the stock table.
							lerr = methods.AddStock(ldb, lmedicineMasterID)
							if lerr != nil {
								// Log the error and update response status and error message if adding stock fails.
								log.Println(lerr)
								lresp.MResp.ErrMsg = "AAddStock006: " + lerr.Error()
								lresp.MResp.Status = "E"
							} else {
								// Log success message and update response message.
								log.Println("Inserted Successfully")
								lresp.Msg = "Added Successfully"
							}
						}

					}
				}
			}
		}
		// here we marshal the lresp struct to JSON
		ldata, lerr := json.Marshal(lresp)
		if lerr != nil {
			// Log the error if JSON marshaling fails and return an error response.
			fmt.Fprintf(w, "Error taking data"+"AAddStock008"+lerr.Error())
		} else {
			// Write the JSON response with appropriate status code.
			fmt.Fprint(w, string(ldata))
		}
		log.Println(" CheckLoginUserinsert (-)")
	}
}
