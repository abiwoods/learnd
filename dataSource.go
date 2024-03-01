package main

import (
	"context"
	"time"
)

type dataSource interface {
	getMetersForCustomer(context.Context, int64) ([]string, error)
	getReadingAtDate(int64, string, time.Time) (int64, error)
}
