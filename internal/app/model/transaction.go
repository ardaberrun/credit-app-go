package model

import "time"

type Transaction struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type GiveCreditRequest struct {
	Id     int   `json:"id" bind:"required"`
	Amount int64 `json:"amount" bind:"required"`
}

type TransferUserCreditRequest struct {
	From   int   `json:"from" bind:"required"`
	To     int   `json:"to" bind:"required"`
	Amount int64 `json:"amount" bind:"required"`
}

type WithdrawRequest struct {
	Id     int   `json:"id" bind:"required"`
	Amount int64 `json:"amount" bind:"required"`
}

type GetUserBalanceAtDateRequest struct {
	Id   int       `json:"id" bind:"required"`
	Date time.Time `json:"date" bind:"required"`
}