package models

type Outcome struct {
	LoanApproved  bool    `bson:"loan_approved,omitempty" json:"loan_approved,omitempty"`
	PreAssessment float64 `bson:"pre_assessment,omitempty" json:"pre_assessment,omitempty"`
}
