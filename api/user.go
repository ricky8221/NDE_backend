package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	ndedb "github.com/ricky8221/NDE_DB/db/sqlc"
	"github.com/ricky8221/NDE_DB/dbFunc"
	"github.com/ricky8221/NDE_DB/util"
	"net/http"
	"time"
)

type loginUserRequest struct {
	Username string `json:"username"  binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	Username         string    `json:"username"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreatedAt        time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	Username string `json:"username"  binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"accessToken"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, "Do not found the user")
		}
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
		user.Role,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := &loginUserResponse{
		Username:    user.Username,
		FullName:    user.FullName,
		AccessToken: accessToken,
	}
	ctx.JSON(http.StatusOK, res)
}

func (server *Server) createUser(ctx *gin.Context) {
	var req ndedb.CreateUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	req.Password = hashedPassword

	user, err := server.store.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := dbFunc.UserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	ctx.JSON(http.StatusOK, res)
}
