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

// loginresponse struct defines the response structure for the CheckLoginUserinsertApi.
type loginresponse struct {
	Role           string               `json:"role"`
	LoginHistoryId int                  `json:"loginHistoryID"`
	MResp          structure.MasterResp `json:"resp"`
}

// loginReq struct defines the structure for the login request data.
type loginReq struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

// this API will get the request data userId and Password . it has four methods to validate login and give response data role,status,errMsg
func CheckLoginUserinsertApi(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	log.Println("CheckLogin  +")

	// Handle only POST requests.
	if r.Method == "POST" {
		var llogindetails loginReq
		var lresp loginresponse

		// Set default response status to success.
		lresp.MResp.Status = "S"
		// Read the request body.
		lbody, lerr := ioutil.ReadAll(r.Body)
		log.Print(lbody)
		if lerr != nil {
			// Log the error and update response status and error message if reading body fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "ALogin001" + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Unmarshal the request body into loginReq struct.
			lerr = json.Unmarshal(lbody, &llogindetails)
			if lerr != nil {
				// Log the error and update response status and error message if unmarshalling fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "ALogin002" + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Connect to the local database.
				db, lerr := db.LocalDBConnect()
				if lerr != nil {
					// Log the error and update response status and error message if connection fails.
					log.Println(lerr)
					lresp.MResp.ErrMsg = "ALogin003" + lerr.Error()
					lresp.MResp.Status = "E"
				} else {
					// Ensure the database connection is closed after use.
					defer db.Close()
					// Method 1: Check if the user ID exists in the medapp_login table.
					lloginId, lerr := methods.CheckUserId(db, llogindetails.UserId)
					if lerr != nil {
						// Log the error and update response status and error message if user ID check fails.
						log.Println(lerr)
						lresp.MResp.ErrMsg = "ALogin004" + lerr.Error()
						lresp.MResp.Status = "E"
					} else {
						// Method 2: Check if the entered password matches the stored password.
						passwordMatch := methods.CheckPassword(db, llogindetails.UserId, llogindetails.Password)
						if passwordMatch {
							// Method 3: Get the user role using the login ID.
							lrole, lerr := methods.GetUserRole(db, lloginId)
							if lerr != nil {
								// Log the error and update response status and error message if fetching role fails.
								log.Println(lerr)
								lresp.MResp.ErrMsg = "ALogin005" + lerr.Error()
								lresp.MResp.Status = "E"
							} else {
								lresp.Role = lrole
								// Method 4: Insert login history into medapp_login_history table.
								lloginHistoryId, lerr := methods.PostLoginHistory(db, lloginId, true)
								if lerr != nil {
									// Log the error and update response status and error message if inserting login history fails.
									log.Println(lerr)
									lresp.MResp.ErrMsg = "ALogin006" + lerr.Error()
									lresp.MResp.Status = "E"
								} else {
									lresp.LoginHistoryId = int(lloginHistoryId)
									log.Println("inserted Successfully")
								}
							}
						} else {
							// If password does not match, set error status and message.
							lresp.MResp.Status = "E"
							lresp.MResp.ErrMsg = "ALogin007" + "Invalid password"
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
			fmt.Fprintf(w, "Error taking data"+"ALogin008"+lerr.Error())
		} else {
			// Send the JSON response.
			fmt.Fprint(w, string(ldata))
		}
		log.Println(" CheckLogin (-)")
	}
}
