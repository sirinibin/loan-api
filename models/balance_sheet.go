package models

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
)

type Sheet struct {
	Year         *int64   `bson:"year,omitempty" json:"year,omitempty"`
	Month        *int64   `bson:"month,omitempty" json:"month,omitempty"`
	ProfitOrLoss *float64 `bson:"profitOrLoss,omitempty" json:"profitOrLoss,omitempty"`
	AssetsValue  *float64 `bson:"assetsValue,omitempty" json:"assetsValue,omitempty"`
}

type BalanceSheet struct {
	Business Business `bson:"business,omitempty" json:"business,omitempty"`
	Sheets   []Sheet  `bson:"sheets,omitempty" json:"sheets,omitempty"`
}

func (balanceSheet *BalanceSheet) Validate(w http.ResponseWriter, r *http.Request) (errs map[string]string) {

	errs = make(map[string]string)
	if govalidator.IsNull(balanceSheet.Business.BusinessName) {
		errs["business_name"] = "Name is required"
	}

	if balanceSheet.Business.YearEstablished == nil {
		errs["year_established"] = "Year Est. is required"
	}

	if balanceSheet.Business.LoanAmount == nil {
		errs["loan_amount"] = "Loan amount is required"
	}

	if govalidator.IsNull(balanceSheet.Business.AccountProvider) {
		errs["account_provider"] = "Account provider is required"
	}

	for index, sheet := range balanceSheet.Sheets {
		if sheet.Year == nil {
			errs["year_"+strconv.Itoa(index)] = "Year is required for balance sheet"
		}

		if sheet.Month == nil {
			errs["month_"+strconv.Itoa(index)] = "Month is required for balance sheet"
		}

		if sheet.ProfitOrLoss == nil {
			errs["profit_or_loss_"+strconv.Itoa(index)] = "Profit or Loss is required for balance sheet"
		}

		if sheet.AssetsValue == nil {
			errs["assets_value_"+strconv.Itoa(index)] = "Assets value is required for balance sheet"
		}
	}

	return errs
}

func (balanceSheet *BalanceSheet) GetOutcome() (*Outcome, error) {
	outcome := Outcome{PreAssessment: 20, LoanApproved: false}

	profit := float64(0)
	totalAssetsValue := float64(0)
	averageAssetsValue := float64(0)

	month := 1
	for _, sheet := range balanceSheet.Sheets {
		if month > 12 {
			break
		}

		profit += *sheet.ProfitOrLoss
		totalAssetsValue += *sheet.AssetsValue
		month++
	}

	averageAssetsValue = totalAssetsValue / float64(len(balanceSheet.Sheets))

	if profit > 0 {
		outcome.PreAssessment = 60
	}

	if averageAssetsValue > *balanceSheet.Business.LoanAmount {
		outcome.PreAssessment = 100
		outcome.LoanApproved = true
	}

	return &outcome, nil
}
