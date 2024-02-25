package api

import (
	"NDE_backend/util"
	"github.com/gin-gonic/gin"
	ndedb "github.com/ricky8221/NDE_DB/db/sqlc"
)

// Server servers HTTP requests
type Server struct {
	router *gin.Engine
}

func NewServer(config util.Config, store *ndedb.Store) (*Server, error) {

}
