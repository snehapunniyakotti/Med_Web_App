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

// newUserResp defines the structure of the response body for adding a new user.
type newUserResp struct {
	Msg   string               `json:"msg"`
	MResp structure.MasterResp `json:"resp"`
}

// newUserReq defines the structure of the request body for adding a new user.
type newUserReq struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// this API will get the request data userId and Password . it has four methods to validate login and give response data role,status,errMsg
func PostUserApi(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin and specify allowed headers and methods.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	log.Println("PostUserApi  +")

	// Handle only POST requests.
	if r.Method == "POST" {
		var luser newUserReq
		var lresp newUserResp

		// Initialize the response structure and set default response status to success.
		lresp.MResp.Status = "S"
		// Read the request body.
		lbody, lerr := ioutil.ReadAll(r.Body)
		if lerr != nil {
			// Log the error and update response status and error message if reading request body fails.
			log.Println(lerr)
			lresp.MResp.ErrMsg = "AAU001" + lerr.Error()
			lresp.MResp.Status = "E"
		} else {
			// Unmarshal the request body into newUserReq struct.
			lerr = json.Unmarshal(lbody, &luser)
			if lerr != nil {
				// Log the error and update response status and error message if unmarshaling fails.
				log.Println(lerr)
				lresp.MResp.ErrMsg = "AAU002" + lerr.Error()
				lresp.MResp.Status = "E"
			} else {
				// Connect to the local database.
				db, lerr := db.LocalDBConnect()
				if lerr != nil {
					// Log the error and update response status and error message if database connection fails.
					log.Println(lerr)
					lresp.MResp.ErrMsg = "AAU003" + lerr.Error()
					lresp.MResp.Status = "E"
				} else {
					// Ensure the database connection is closed after use.
					defer db.Close()
					// Method 1 - Check if the user ID already exists.
					if methods.IsUserIdExists(db, luser.UserId) {
						lresp.MResp.ErrMsg = "User ID already exists"
						lresp.MResp.Status = "E"
					} else {
						// Method 2 - Add the new user if the user ID is not present.
						lerr = methods.AddUser(db, luser.UserId, luser.Password, luser.Role)
						if lerr != nil {
							// Log the error and update response status and error message if adding user fails.
							log.Println(lerr)
							lresp.MResp.ErrMsg = "AAU005" + lerr.Error()
							lresp.MResp.Status = "E"
						} else {
							lresp.Msg = "User added successfully"
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
			// Log the error if JSON marshaling fails and return an error response.
			fmt.Fprintf(w, "Error taking data"+"AAU008"+lerr.Error())
		} else {
			// Write the JSON response with appropriate status code.
			fmt.Fprint(w, string(ldata))
		}
		log.Println(" PostUserApi (-)")
	}
}
