package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Foo",
			Address: "Mon Adresse",
			Phone:   "07888888",
		},
		Number:  245,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("Nom du client vide")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Foo",
			Address: "Mon Adresse",
			Phone:   "07888888",
		},
		Number:  245,
		Balance: 0,
	}

	account.Deposit(100)

	if account.Balance != 100 {
		t.Error("Solde invalide")
	}
}

func TestDepositNegative(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Foo",
			Address: "Mon Adresse",
			Phone:   "07888888",
		},
		Number:  245,
		Balance: 0,
	}

	err := account.Deposit(-100)
	if err == nil {
		t.Error("Erreur non retournée")
	}
}
