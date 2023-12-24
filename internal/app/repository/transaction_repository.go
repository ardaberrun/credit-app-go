package repository

import (
	"database/sql"
	"fmt"
	"github/ardaberrun/credit-app-go/internal/app/model"
	"time"
)

type ITransactionRepository interface {
	GiveCredit(id int, amount int64) error
	TransferUserCredit(sender, receiver int, amount int64) error
	WithdrawMoney(id int, amount int64) error
	GetTransactionsByUserId(id int) ([]*model.Transaction, error)
	GetUserBalanceAtDate(id int, d time.Time) (int64, error)
}

type TransactionRepository struct {
	db *sql.DB
}

func InitializeTransactionRepository(db *sql.DB) ITransactionRepository {
	return &TransactionRepository{db: db}
}


func (tr *TransactionRepository) GiveCredit(id int, amount int64) error {
	tx, err := tr.db.Begin();

	_, err = tx.Exec("UPDATE USERS SET balance = balance + $2 WHERE id = $1", id, amount);
	if err != nil {
		tx.Rollback();

		return err;
	}

	_, err = tx.Exec("INSERT INTO TRANSACTIONS (user_id, amount, created_at) VALUES($1, $2, $3)", id, amount, time.Now().UTC())
	if err != nil {
		tx.Rollback();

		return err;
	}

	err = tx.Commit();
	if err != nil {
		return err;
	}

	return nil;
}

func (tr *TransactionRepository) TransferUserCredit(sender, receiver int, amount int64) error {
	var senderBalance int64;

	err := tr.db.QueryRow("SELECT balance FROM USERS WHERE id = $1", sender).Scan(&senderBalance);
	if err != nil {
		return err;
	}

	if senderBalance < amount {
		return fmt.Errorf("You do not have sufficient funds to perform this transaction.");
	}

	tx, err := tr.db.Begin();
	if err != nil {
		return err;
	}

	_, err = tx.Exec("UPDATE USERS SET balance = balance - $1 WHERE id = $2", amount, sender)
	if err != nil {
		tx.Rollback();

		return err;
	}

	_, err = tx.Exec("UPDATE USERS SET balance = balance + $1 WHERE id = $2", amount, receiver);
	if err != nil {
		tx.Rollback();

		return err;
	}

	_, err = tx.Exec("INSERT INTO TRANSACTIONS (user_id, amount, created_at) VALUES($1, $2, $3)", sender, amount * -1, time.Now().UTC())
	if err != nil {
		tx.Rollback();

		return err;
	}

	_, err = tx.Exec("INSERT INTO TRANSACTIONS (user_id, amount, created_at) VALUES($1, $2, $3)", receiver, amount, time.Now().UTC())
	if err != nil {
		tx.Rollback();

		return err;
	}

	err = tx.Commit();
	if err != nil {
		return err;
	}

	return nil;
}

func (tr *TransactionRepository) WithdrawMoney(id int, amount int64) error {
	var balance int64;

	err := tr.db.QueryRow("SELECT balance FROM USERS WHERE id = $1", id).Scan(&balance);
	if err != nil {
		return err;
	}

	if balance < amount {
		return fmt.Errorf("You do not have sufficient funds to perform this transaction.");
	}

	tx, err := tr.db.Begin();
	if err != nil {
		return err;
	}

	_, err = tx.Exec("UPDATE USERS SET balance = balance - $2 WHERE id = $1", id, amount);
	if err != nil {
		tx.Rollback();

		return err;
	}

	_, err = tx.Exec("INSERT INTO TRANSACTIONS (user_id, amount, created_at) VALUES($1, $2, $3)", id, amount * -1, time.Now().UTC())
	if err != nil {
		tx.Rollback();

		return err;
	}

	err = tx.Commit();
	if err != nil {
		return err;
	}

	return nil;
}

func (tr *TransactionRepository) GetTransactionsByUserId(id int) ([]*model.Transaction, error) {
	rows, err := tr.db.Query("SELECT * FROM TRANSACTIONS WHERE USER_ID = $1", id);
	if err != nil {
		return nil, err
	}
	defer rows.Close();

	txs := []*model.Transaction{}
	for rows.Next() {
		tx := new(model.Transaction)

		if err := rows.Scan(&tx.Id, &tx.UserId, &tx.Amount, &tx.CreatedAt); err != nil {
			return nil, err
		}

		txs = append(txs, tx);
	}

	return txs, nil;
}

func (tr *TransactionRepository) GetUserBalanceAtDate(id int, d time.Time) (int64, error) {
	var balance int64;

	query :=
	`SELECT COALESCE(SUM(amount), 0) AS account_balance FROM TRANSACTIONS
		WHERE user_id = $1 AND created_at <= $2`;

	err := tr.db.QueryRow(query, id, d).Scan(&balance);
	if err != nil {
		return 0, err
	}

	return balance, nil;
}