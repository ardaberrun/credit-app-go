package service

import (
	"time"
	"github/ardaberrun/credit-app-go/internal/app/model"
	"github/ardaberrun/credit-app-go/internal/app/repository"
)


type TransactionService struct {
	transactionRepository repository.ITransactionRepository
}

type ITransactionService interface {
	GiveCredit(id int, amount int64) error
	TransferUserCredit(sender, receiver int, amount int64) error
	WithdrawMoney(id int, amount int64) error
	GetTransactionsByUserId(id int) ([]*model.Transaction, error)
	GetUserBalanceAtDate(id int, d time.Time) (int64, error)
}

func InitializeTransactionService(transactionRepository repository.ITransactionRepository) *TransactionService {
	return &TransactionService{transactionRepository: transactionRepository};
}

func (ts *TransactionService) GiveCredit(id int, amount int64) error {
	return ts.transactionRepository.GiveCredit(id, amount);
}

func (ts *TransactionService) TransferUserCredit(sender, receiver int, amount int64) error {
	return ts.transactionRepository.TransferUserCredit(sender, receiver, amount);
}

func (ts *TransactionService) WithdrawMoney(id int, amount int64) error {
	return ts.transactionRepository.WithdrawMoney(id, amount);
}

func (ts *TransactionService) GetTransactionsByUserId(id int) ([]*model.Transaction, error) {
	return ts.transactionRepository.GetTransactionsByUserId(id);
}

func (ts *TransactionService) GetUserBalanceAtDate(id int, d time.Time) (int64, error) {
	return ts.transactionRepository.GetUserBalanceAtDate(id, d);
}