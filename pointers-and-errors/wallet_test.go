package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		balance := wallet.Balance()
		expected := Bitcoin(10)

		if expected != balance {
			t.Errorf("Expected '%s' but received '%s'", expected, balance)
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(50))

		balance := wallet.Balance()
		expected := Bitcoin(50)

		if expected != balance {
			t.Errorf("Expected '%s' but received '%s'", expected, balance)
		}
		if err != nil {
			t.Errorf("Expected nil but received '%s'", err)
		}
	})
	t.Run("withdraw more than balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(150))

		balance := wallet.Balance()
		expected := Bitcoin(100)
		expectedError := ErrInsufficientFunds

		if expected != balance {
			t.Errorf("Expected '%s' but received '%s'", expected, balance)
		}

		if err == nil {
			// Fatal stops the test if this one is failed to prevent run next assertions
			t.Fatal("Expected an error but didn't get one")
		}
		if err != expectedError {
			t.Errorf("Expected an error message %q but received %q", err.Error(), expectedError)
		}
	})
}
