package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	ndedb "github.com/ricky8221/NDE_DB/db/sqlc"
	"net/http"
)

func (server *Server) createCompany(ctx *gin.Context) {
	var req ndedb.CreateCompanyParams

	// Bind the JSON to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company: " + err.Error()})
		return
	}

	company, err := server.store.CreateCompany(ctx, req)
	if err != nil {
		// Handle the error appropriately
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "company": company})
}

type GetCompanyRequest struct {
	CompanyName string `json:"companyName"`
}

func (server *Server) getCompany(ctx *gin.Context) {
	var req GetCompanyRequest

	// Bind the JSON to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, err := server.store.GetCompany(ctx, req.CompanyName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadRequest, err)
		}
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, company)

}
