package main

type Payment struct {
	Id     string `json:"id"`
	Amount int64  `json:"amount"`
}

type FetchAllPaymentsResponse struct {
	Count  int64     `json:"count"`
	Entity int64     `json:"entity"`
	Items  []Payment `json:"items"`
}
