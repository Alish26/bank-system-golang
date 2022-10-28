package presenter

import (
	"time"

	"github.com/Alish26/bank-system-golang/domain"
	"github.com/Alish26/bank-system-golang/usecase"
)

type createAccountPresenter struct{}

func NewCreateAccountPresenter() usecase.CreateAccountPresenter {
	return createAccountPresenter{}
}

func (a createAccountPresenter) Output(account domain.Account) usecase.CreateAccountOutput {
	return usecase.CreateAccountOutput{
		ID:        account.ID().String(),
		Name:      account.Name(),
		CPF:       account.CPF(),
		Balance:   account.Balance().Float64(),
		CreatedAt: account.CreatedAt().Format(time.RFC3339),
	}
}
