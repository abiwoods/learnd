package main

// Have added a building ID, as this is what I would expect in a production system
type building struct {
	ID   int64  `json:"buildingID"`
	Name string `json:"building"`
}
