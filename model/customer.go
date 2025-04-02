package model

type CreateCustomerRequest struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Gender       string `json:"gender"`
	Tel          string `json:"tel"`
	EmailAddress string `json:"email_address"`
}

type UpdateCustomerRequest struct {
	ID           string `uri:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Gender       string `json:"gender"`
	Tel          string `json:"tel"`
	EmailAddress string `json:"email_address"`
}
