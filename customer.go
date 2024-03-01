package main

// Have added a customer ID, as this is what I would expect in a production system
type customer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
