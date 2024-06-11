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

// logoutresponse struct defines the response structure for the UpdateLoginHistoryApi.
type logoutresponse struct {
	Msg   string               `json:"msg"`
	MResp structure.MasterResp `json:"resp"`
}

// logoutReq struct defines the structure for the logout request data.
type logoutReq struct {
	LoginHistoryId int `json:"loginHistoryId"`
}

// UpdateLoginHistoryApi is used to update the login history when a user logs out.
// It takes the loginHistoryId as request data.
func UpdateLoginHistoryApi(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")

	log.Println("logout +")

	// Handle only PUT requests.
	if r.Method == "PUT" {
		var ldetails logoutReq
		var lresp logoutresponse

		// Set default response status to success.
		lresp.MResp.Status = "S"

		// Read the request body
		lbody, lerr := ioutil.ReadAll(r.Body)
		if lerr != nil {
			// Log the error and update response status and error message if reading body fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "ALogout001" + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Unmarshal the request body into logoutReq struct.
			lerr = json.Unmarshal(lbody, &ldetails)
			if lerr != nil {
				// Log the error and update response status and error message if unmarshalling fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "ALogout002" + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Connect to the local database.
				ldb, lerr := db.LocalDBConnect()
				if lerr != nil {
					// Log the error and update response status and error message if connection fails.
					log.Println(lerr)
					lresp.MResp.ErrMsg = "ALogout003" + lerr.Error()
					lresp.MResp.Status = "E"
				} else {
					// Ensure the database connection is closed after use.
					defer ldb.Close()

					// Call the method to update the login history table.
					lerr := methods.UpdateLoginHistory(ldb, ldetails.LoginHistoryId)
					if lerr != nil {
						// Log the error and update response status and error message if update fails.
						log.Println(lerr)
						lresp.MResp.ErrMsg = "ALogout004" + lerr.Error()
						lresp.MResp.Status = "E"
					} else {
						// If update is successful, set success message.
						lresp.Msg = "updated Successfully"
						log.Println("updated Successfully")
					}
				}
			}
		}
		// Marshal the response structure to JSON format.
		data, lerr := json.Marshal(lresp)
		if lerr != nil {
			// Send an error message if marshalling fails.
			fmt.Fprintf(w, "Error taking data"+"ALogout005"+lerr.Error())
		} else {
			// Send the JSON response.
			fmt.Fprint(w, string(data))
		}
		log.Println(" logout (-)")
	}

}
