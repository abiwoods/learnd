package main

import (
	"context"
	"database/sql"
	"time"
)

type database struct{}

// Here, I am only setting up the data we need for the examples
var usage map[string]int
var customerMeters map[int64][]string

func initDB() (dataSource, error) {
	db := database{}

	// This data would be set up completely differently in a real db,
	// this is just to mock for the endpoints we need...

	// In a production system, you would expect to use customer IDs rather than
	// name - here I have used
	// 1 = Aquaflow
	// 2 = Albers Facilities Management
	customerMeters = map[int64][]string{
		1: []string{"1111-1111-1111", "1111-1111-2222"},
		2: []string{"1111-1111-3333"},
	}

	usage = map[string]int{
		"1111-1111-1111": 20,
		"1111-1111-2222": 30,
		"1111-1111-3333": 40,
	}
	return db, nil
}

func (db database) getMetersForCustomer(ctx context.Context, customerID int64) ([]string, error) {
	meterIDs, ok := customerMeters[customerID]
	if !ok {
		return nil, sql.ErrNoRows
	}

	return meterIDs, nil
}

func (db database) getReadingAtDate(customerID int64, meterID string, date time.Time) (int64, error) {
	// This just gives the reading for the current month usage, could be adjusted to give different functionality
	reading, ok := usage[meterID]
	if !ok {
		return 0, sql.ErrNoRows
	}

	return int64(date.Day() * reading), nil
}
