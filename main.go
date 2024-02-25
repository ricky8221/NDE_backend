package main

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ricky8221/NDE_DB/sqlc_func"
	"github.com/ricky8221/NDE_DB/util"
	"log"
	"net/http"
)

func main() {
	// Initialize the Gin engine.
	r := gin.Default()

	// DB connect string
	connStr := util.GetConnStr()

	// Open a database connection.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify our connection is good.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r.POST("/createCompany", func(c *gin.Context) {
		var req sqlc_func.CreateCompanyReq

		// Bind the JSON to the struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		company, err := sqlc_func.CreateCompany(context.Background(), db, req)
		if err != nil {
			// Handle the error appropriately
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "company": company})
	})

	r.Run()
}
