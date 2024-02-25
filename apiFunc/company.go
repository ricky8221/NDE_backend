package apiFunc

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/ricky8221/NDE_DB/sqlc_func"
	"net/http"
)

func CreateCompany(ctx *gin.Context, db *sql.DB) {
	var req sqlc_func.CreateCompanyReq

	// Bind the JSON to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, err := sqlc_func.CreateCompany(context.Background(), db, req)
	if err != nil {
		// Handle the error appropriately
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "company": company})
}
