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

// saveBillresponse defines the structure of the response for the SaveBillinsertApi.
type saveBillresponse struct {
	Msg   structure.MasterResp `json:"msg"`
	MResp structure.MasterResp `json:"resp"`
}

// saveBillReq defines the structure of the request body for the SaveBillinsertApi.
type saveBillReq struct {
	BillNo     int             `json:"billNo"`
	BillAmount int             `json:"billAmount"`
	BillGst    float32         `json:"billGst"`
	NetPrice   int             `json:"netPrice"`
	UserId     string             `json:"userId"`
	MedArr     []medicineEntry `json:"medArr"`
}

// medicineEntry defines the structure of each medicine entry in the bill.
type medicineEntry struct {
	MedicineName string `json:"medicineName"`
	Quantity     int    `json:"quantity"`
	UnitPrice    int    `json:"unitPrice"`
	Amount       int    `json:"amount"`
}

// this API will get the request data userId and Password . it has four methods to validate login and give response data role,status,errMsg
func SaveBillinsertApi(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	log.Println("SaveBill  (+)")

	// Handle only POST requests.
	if r.Method == "POST" {
		var lreq saveBillReq
		var lresp saveBillresponse

		// Set default response status to success.
		lresp.MResp.Status = "S"
		// Read the request body.
		lbody, lerr := ioutil.ReadAll(r.Body)
		if lerr != nil {
			// Log the error and update response status and error message if reading the body fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "AL001" + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Unmarshal the request body into the lreq struct.
			lerr = json.Unmarshal(lbody, &lreq)
			if lerr != nil {
				// Log the error and update response status and error message if unmarshalling fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "AL002" + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Connect to the local database.
				db, lerr := db.LocalDBConnect()
				if lerr != nil {
					// Log the error and update response status and error message if connection fails.
					log.Println(lerr)
					lresp.MResp.ErrMsg = "AL003" + lerr.Error()
					lresp.MResp.Status = "E"
				} else {
					// Ensure the database connection is closed after use.
					defer db.Close()
					// Method 1 - Get the login ID for the given user ID.
					log.Print(lreq.UserId)
					loginId, lerr := methods.GetLoginId(db, lreq.UserId)
					log.Print(loginId)
					if lerr != nil {
						// Log the error and update response status and error message if fetching login ID fails.
						log.Println(lerr)
						lresp.MResp.ErrMsg = "AL004" + lerr.Error()
						lresp.MResp.Status = "E"
					} else {
						// Method 2 - Insert bill master details.
						lerr := methods.PostBillMaster(db, lreq.BillNo, lreq.BillAmount, lreq.BillGst, lreq.NetPrice, loginId)
						if lerr != nil {
							// Log the error and update response status and error message if inserting bill master fails.
							log.Println(lerr)
							lresp.MResp.ErrMsg = "AL004" + lerr.Error()
							lresp.MResp.Status = "E"
						} else {

							// Iterate over each medicine entry in the request.
							for _, med := range lreq.MedArr {
								// Method 3 -  Get the medicine master ID for the given medicine name.
								medicineMasterId, lerr := methods.GetMedicineMasterId(db, med.MedicineName)
								if lerr != nil {
									// Log the error and update response status and error message if fetching medicine master ID fails.
									log.Println(lerr)
									lresp.MResp.ErrMsg = "AL005" + lerr.Error()
									lresp.MResp.Status = "E"
									break
								}
								// Method 4 -  Insert bill details for each medicine.
								lerr = methods.PostBillDetails(db, lreq.BillNo, medicineMasterId, med.Quantity, med.UnitPrice, med.Amount)
								if lerr != nil {
									// Log the error and update response status and error message if inserting bill details fails.
									log.Println(lerr)
									lresp.MResp.ErrMsg = "AL006" + lerr.Error()
									lresp.MResp.Status = "E"
									break
								}
								// Method 5 - Update the stock for each medicine.
								lerr = methods.UpdateStock(db, medicineMasterId, med.Quantity)
								if lerr != nil {
									// Log the error and update response status and error message if updating stock fails.
									log.Println(lerr)
									lresp.MResp.ErrMsg = "AL007" + lerr.Error()
									lresp.MResp.Status = "E"
									break
								}
							}
						}

					}
				}
			}
		}
		// log.Print(lresp)
		// here we marshal the lresp struct to JSON
		ldata, lerr := json.Marshal(lresp)
		// log.Print(ldata)
		if lerr != nil {
			// Send an error message if marshalling fails.
			fmt.Fprintf(w, "Error taking data"+"AL008"+lerr.Error())
		} else {
			// Write the JSON response to the client.
			fmt.Fprint(w, string(ldata))
		}
		log.Println(" SaveBill (-)")
	}
}
