package dto

type CreateCustomerRequest struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Gender       string `json:"gender"`
	Tel          string `json:"tel"`
	EmailAddress string `json:"email_address"`
}

type UpdateCustomerRequest struct {
	ID        uint    `uri:"id"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Gender    *string `json:"gender"`
	Tel       *string `json:"tel"`
}

type GetCustomerResponse struct {
	ID           uint   `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Gender       string `json:"gender"`
	Tel          string `json:"tel"`
	EmailAddress string `json:"email_address"`
}

type CustomerRequest struct {
	ID uint `uri:"id"`
}
