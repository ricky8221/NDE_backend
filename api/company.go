package api

import (
	"github.com/gin-gonic/gin"
	ndedb "github.com/ricky8221/NDE_DB/db/sqlc"
	"github.com/ricky8221/NDE_DB/sqlc_func"
	"github.com/sqlc-dev/pqtype"
	"net/http"
)

func (server *Server) createCompany(ctx *gin.Context) {
	var req sqlc_func.CreateCompanyReq

	// Bind the JSON to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var remark pqtype.NullRawMessage
	if len(req.Remark) > 0 {
		remark.Valid = true
		remark.RawMessage = req.Remark
	} else {
		remark.Valid = false
	}

	var createCompanyParams ndedb.CreateCompanyParams = ndedb.CreateCompanyParams{
		CompanyName:          req.CompanyName,
		CompanyContactName:   req.CompanyContactName,
		CompanyContactNumber: req.CompanyContactNumber,
		Remark:               remark,
	}

	company, err := server.store.CreateCompany(ctx, createCompanyParams)
	if err != nil {
		// Handle the error appropriately
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "company": company})
}
