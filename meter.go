package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Meter struct {
	BuildingID int64  `json:"buildingID"`
	CustomerID int64  `json:"customerID"`
	ID         string `json:"ID"`
}

// handler for /customers/:customerID/meters
func getMetersHandler(c *gin.Context) {
	ctx := c.Request.Context()

	custIDParam := c.Param("customerID")
	customerID, err := strconv.ParseInt(custIDParam, 10, 64)
	if err != nil {
		log.Println(fmt.Errorf("invalid customer ID - %s: %w", custIDParam, err))
		// Could add error into response depending on who was ingesting data
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// In production system would do checks for permissions, authentication etc here to make sure user allowed to access info
	metersIDs, err := db.getMetersForCustomer(ctx, customerID)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Println(fmt.Errorf("retrieving meters - customerID %d: customer not found", customerID))
		c.AbortWithStatus(http.StatusNotFound)
		return
	case err != nil:
		log.Println(fmt.Errorf("retrieving meters - customerID %d: %w", customerID, err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, metersIDs)
}
