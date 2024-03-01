package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// handler for /customers/:customerID/meters/:serialID/usage?readingDate=:readingDate
// readingDate has format 2006-01-02
func getUsageHandler(c *gin.Context) {
	custIDParam := c.Param("customerID")
	customerID, err := strconv.ParseInt(custIDParam, 10, 64)
	if err != nil {
		log.Println(fmt.Errorf("invalid customer ID - %s: %w", custIDParam, err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	serialID := c.Param("serialID")
	if serialID == "" {
		log.Println(errors.New("missing serial ID"))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	readingDate := c.Query("readingDate")
	if readingDate == "" {
		log.Println(errors.New("missing reading date"))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	parsedDate, err := time.Parse("2006-01-02", readingDate)
	if err != nil {
		log.Println(fmt.Errorf("invalid date - %s: %w", readingDate, err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	reading, err := db.getReadingAtDate(customerID, serialID, parsedDate)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Println(fmt.Errorf("reading not found for customer %d; meter %s; date %s", customerID, serialID, readingDate))
		c.AbortWithStatus(http.StatusNotFound)
		return
	case err != nil:
		log.Println(fmt.Errorf("error rerieving reading for customer %d; meter %s; date %s: %w", customerID, serialID, readingDate, err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, reading)
}
