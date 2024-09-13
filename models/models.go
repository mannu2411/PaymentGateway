package models

type PaymentInfo struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Receipt  string `json:"receipt"`
}
type CustomerInfo struct {
	Name         string   `json:"name"`
	Contact      int64    `json:"contact"`
	Email        string   `json:"email"`
	FailExisting int      `json:"fail_existing"`
	Gstin        string   `json:"gstin"`
	Notes        []string `json:"notes"`
}

type UpdateCustomerInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Contact int64  `json:"contact"`
	Email   string `json:"email"`
}
