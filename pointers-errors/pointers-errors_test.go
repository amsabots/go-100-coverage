package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T){
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(15))
	})

	t.Run("withdraw insufficient fund", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(50))
		assertError(t, err, ErrInsufficientBalance)
        assertBalance(t, wallet, Bitcoin(20)) 

	})	
}

func assertBalance (t testing.TB, w Wallet, want Bitcoin)  {
	t.Helper()
  got := w.balance

  if got != want {
	t.Errorf("got %s want %s", got, want)
} 

}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("wanted an error but didn't get one")
	}
	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}

}



