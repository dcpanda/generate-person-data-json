package api

import (
	"github.com/dcpanda/generate-person-data-json/pkg/person"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SetupRouter configures the Gin router with all necessary routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Add a health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Add an endpoint to generate person data
	router.GET("/api/persons", getPersons)

	return router
}

// getPersons handles requests to generate person data
// Query parameter: n - number of records to generate (default: 10)
func getPersons(c *gin.Context) {
	// Parse the number of records from query parameter
	numRecordsStr := c.DefaultQuery("n", "10")
	numRecords, err := strconv.Atoi(numRecordsStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid number of records",
		})
		return
	}

	// Validate the number of records
	if numRecords <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Number of records must be greater than 0",
		})
		return
	}

	// Generate the person records
	personRecords, err := person.GeneratePersonRecords(numRecords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the person records as JSON
	c.JSON(http.StatusOK, personRecords)
}

// StartServer starts the HTTP server on the specified port
func StartServer(port string) error {
	router := SetupRouter()
	return router.Run(port)
}
