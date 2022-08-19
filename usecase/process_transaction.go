package usecase

import "github.com/ghindirestraits/imersao-full-stack-full-cycle/domain"

type UseCaseTransaction struc {
	TransactionRepository domain.TransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	
}
