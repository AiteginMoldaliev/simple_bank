package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	db "simple-bank/db/sqlc"
	"simple-bank/token"
	"simple-bank/util"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	fromAccount, ok := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !ok {	
		return
	}
	authPayload, ok := ctx.MustGet(util.AuthorizationPayloadKey).(*token.Payload)
	if !ok {
		err := errors.New("mismatched payload type")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if authPayload.Username != fromAccount.Owner {
		err := errors.New("from account doesn't belong ot the authenticiated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	} 

	_, ok = server.validAccount(ctx, req.ToAccountID, req.Currency)
	if !ok {
		return
	}

	arg := db.TransferTxParams{
		FromAccountId: req.FromAccountID,
		ToAccountId:   req.ToAccountID,
		Amount:        req.Amount,
	}

	res, err := server.store.TransferEx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %v vs %v", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true
}