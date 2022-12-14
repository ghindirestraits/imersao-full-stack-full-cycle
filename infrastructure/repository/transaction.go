package repository

import "database/sql"

type TransactionRepositoryDb struc {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (t *TransactionRepositoryDb) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err := t.db.Prepare(query=`insert into transactions(id, credit_card_id, status, description, store, created_at value($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardID,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt,
	)

	if err != nil {
		return err
	}

	if transaction.Status == "approved" {
		err = t.updateBalance(creditCard)

		if err != nil {
			return err
		}
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDb) updateBalance(creditCard domain.CreditCard) error {
	_, err := t.db.Exec("udpate credit_cards set balance = $1 where id = $2",
		creditCard.Balance, creditCard.ID)
	
	if err != nil {
		return err
	}
	
	return nil
}

// GetCreditCard(creditCard CreditCard) (CreditCard, error)
func (t *TransactionRepositoryDb) CreateCreditCard(creditCard CreditCard) error {
	stmt, err := t.db.Prepare(`insert into credit_cards(id, name, number, expiration_month,expiration_year, CVV,balance, balance_limit)	values($1,$2,$3,$4,$5,$6,$7,$8)`)
	
	if err != nil {
		return err
	}
	
	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit,
	)
	
	if err != nil {
		return err
	}
	
	err = stmt.Close()
	
	if err != nil {
		return err
	}
	
	return nil
}
