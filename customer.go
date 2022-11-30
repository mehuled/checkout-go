package main

type Customer struct {
	Id string `json:"id"`
}
type CreateCustomerRequest struct {
	Name         string
	Contact      string
	Email        string
	FailExisting bool
	GSTIN        string
	Notes        string
}
