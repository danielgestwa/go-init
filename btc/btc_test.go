package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))

		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10.0)}
		err := wallet.Withdraw(Bitcoin(3.0))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(7.0))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10.0)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(20.0))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error("didn't expect error but got one")
	}
}

func assertError(t testing.TB, err error, wanted error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}

	if err != wanted {
		t.Errorf("Expected error %s, got %s", err, wanted)
	}
}
