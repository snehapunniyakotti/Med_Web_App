package main

import (
	"command/api"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server is running ")
	http.HandleFunc("/login", api.CheckLoginUserinsertApi)
	http.HandleFunc("/updateloginHistory", api.UpdateLoginHistoryApi)
	http.HandleFunc("/addStock", api.PostStockApi)
	http.HandleFunc("/getMedicineNames", api.GetMedicineName)
	http.HandleFunc("/mndetails", api.FetchMedicineDetails)
	http.HandleFunc("/updatestock", api.UpdateStock)
	http.HandleFunc("/getStockView", api.GetStockView)
	http.HandleFunc("/managerDashboard", api.ManagerDashboard)
	http.HandleFunc("/history", api.GetHistory)
	http.HandleFunc("/addUser", api.PostUserApi)
	http.HandleFunc("/checkQuantity", api.CheckQuantityApi)
	http.HandleFunc("/saveBill", api.SaveBillinsertApi)
	http.HandleFunc("/salesReport", api.GetSalesReport)
	http.HandleFunc("/billerDashboard", api.BillerDashboard)

	http.ListenAndServe(":8081", nil)

}
