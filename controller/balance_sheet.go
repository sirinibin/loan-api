package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sirinibin/loan-api/models"
	"github.com/sirinibin/loan-api/utils"
)

// Get balance sheet : handler for POST /balance-sheet
func GetBalanceSheet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	response.Errors = make(map[string]string)

	var business *models.Business
	// Decode business data
	if !utils.Decode(w, r, &business) {
		return
	}

	// Validate business data
	if errs := business.Validate(w, r); len(errs) > 0 {
		response.Status = false
		response.Errors = errs
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get Balancesheet
	balanceSheet, err := business.GetBalanceSheet()
	if err != nil {
		response.Status = false
		response.Errors["balance_sheet"] = "Unable to find balance sheet:" + err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Status = true
	response.Result = balanceSheet
	response.TotalCount = int64(len(balanceSheet.Sheets))

	json.NewEncoder(w).Encode(response)
}

// Get outcome : handler for POST /outcome
func GetOutcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	response.Errors = make(map[string]string)

	var balanceSheet *models.BalanceSheet
	// Decode balance sheet data
	if !utils.Decode(w, r, &balanceSheet) {
		return
	}

	// Validate balance sheet data
	if errs := balanceSheet.Validate(w, r); len(errs) > 0 {
		response.Status = false
		response.Errors = errs
		json.NewEncoder(w).Encode(response)
		return
	}
	// Get Outcome
	outcome, err := balanceSheet.GetOutcome()
	if err != nil {
		response.Status = false
		response.Errors["outcome"] = "Unable to find outcome:" + err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Status = true
	response.Result = outcome

	json.NewEncoder(w).Encode(response)
}
