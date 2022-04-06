package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "kuann.me/simplebank/db/sqlc"
)

type createAccountRequestParam struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required,oneof=1 2 3"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequestParam
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Username: req.Username,
		Password: req.Password,
	}

	account, err := server.store.CreateAccountTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
