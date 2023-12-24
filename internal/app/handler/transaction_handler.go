package handler

import (
	"github/ardaberrun/credit-app-go/internal/app/model"
	"github/ardaberrun/credit-app-go/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


type TransactionHandler struct {
	transactionService service.ITransactionService
	router *gin.Engine
}

func InitializeTransactionHandler(ts service.ITransactionService, r *gin.Engine) *TransactionHandler {
	return &TransactionHandler{transactionService: ts, router: r};
}


func (h *TransactionHandler) RegisterRoutes() {
	txs := h.router.Group("/transactions")
	{
		txs.GET("/:id", h.GetTransactionsByUserId);
		txs.POST("/credit", h.GiveCredit);
		txs.POST("/transfer", h.TransferUserCredit);
		txs.POST("/withdraw", h.WithdrawMoney);
		txs.POST("/balance", h.GetUserBalanceAtDate);
	}
}

func (h *TransactionHandler) GetTransactionsByUserId(c *gin.Context) {
	claims, exists := c.Get("claims");
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;
	}

	id, err := strconv.Atoi(c.Param("id"));
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"});

		return;
	}

	userId := int(claims.(jwt.MapClaims)["user_id"].(float64));
	userRole := claims.(jwt.MapClaims)["user_role"].(string);

	if userRole != "admin" && userId != id {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;		
	}

	txs, err := h.transactionService.GetTransactionsByUserId(id);
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user's transactions not found"});

		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": txs});
}


func (h *TransactionHandler) GiveCredit(c *gin.Context) {
	var gcr model.GiveCreditRequest;

	if err := c.ShouldBindJSON(&gcr); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	claims, exists := c.Get("claims");
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;
	}

	userRole := claims.(jwt.MapClaims)["user_role"].(string);

	if userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;		
	}

	if err := h.transactionService.GiveCredit(gcr.Id, gcr.Amount); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()});

		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"});
}

func (h *TransactionHandler) TransferUserCredit(c *gin.Context) {
	var tuc model.TransferUserCreditRequest;

	if err := c.ShouldBindJSON(&tuc); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	claims, exists := c.Get("claims");
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;
	}

	userId := int(claims.(jwt.MapClaims)["user_id"].(float64));

	if userId != tuc.From {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;		
	}

	if err := h.transactionService.TransferUserCredit(tuc.From, tuc.To, tuc.Amount); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()});

		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"});
}

func (h *TransactionHandler) WithdrawMoney(c *gin.Context) {
	var wmr model.WithdrawRequest;

	if err := c.ShouldBindJSON(&wmr); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	claims, exists := c.Get("claims");
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;
	}

	userId := int(claims.(jwt.MapClaims)["user_id"].(float64));

	if userId != wmr.Id {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;		
	}

	if err := h.transactionService.WithdrawMoney(wmr.Id, wmr.Amount); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()});

		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"});
}

func (h *TransactionHandler) GetUserBalanceAtDate(c *gin.Context) {
	var gubar model.GetUserBalanceAtDateRequest;

	if err := c.ShouldBindJSON(&gubar); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	claims, exists := c.Get("claims");
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;
	}

	userId := int(claims.(jwt.MapClaims)["user_id"].(float64));
	userRole := claims.(jwt.MapClaims)["user_role"].(string);

	if userRole != "admin" && userId != gubar.Id {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;		
	}

	balance, err := h.transactionService.GetUserBalanceAtDate(gubar.Id, gubar.Date);;
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()});

		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "balance": balance});
}