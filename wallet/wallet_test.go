package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Egold(10))

		assertBalance(t, wallet, Egold(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{ballance: 60}
		err := wallet.Withdraw(Egold(50))

		assertNoError(t, err)
		assertBalance(t, wallet, Egold(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Egold(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Egold(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance (t *testing.T, wallet Wallet, expectedBalance Egold) {
	t.Helper()
	got := wallet.Balance()

	if got != expectedBalance {
		t.Errorf("got %s but expected %s", got, expectedBalance)
	}
}

func assertError (t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError (t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}