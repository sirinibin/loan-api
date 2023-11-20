package models

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
)

type Business struct {
	BusinessName    string   `bson:"business_name,omitempty" json:"business_name,omitempty"`
	YearEstablished *int64   `bson:"year_established,omitempty" json:"year_established,omitempty"`
	LoanAmount      *float64 `bson:"loan_amount,omitempty" json:"loan_amount,omitempty"`
	AccountProvider string   `bson:"account_provider,omitempty" json:"account_provider,omitempty"`
}

func (business *Business) Validate(w http.ResponseWriter, r *http.Request) (errs map[string]string) {

	errs = make(map[string]string)

	if govalidator.IsNull(business.BusinessName) {
		errs["business_name"] = "Name is required"
	}

	if business.YearEstablished == nil {
		errs["year_established"] = "Year Est. is required"
	}

	if business.LoanAmount == nil {
		errs["loan_amount"] = "Loan amount is required"
	}

	if govalidator.IsNull(business.AccountProvider) {
		errs["account_provider"] = "Account provider is required"
	}

	return errs
}

func (business *Business) GetBalanceSheet() (*BalanceSheet, error) {
	balanceSheet := BalanceSheet{Business: *business}

	file, err := os.ReadFile("sheets.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &balanceSheet.Sheets)
	if err != nil {
		return nil, err
	}

	return &balanceSheet, nil
}
