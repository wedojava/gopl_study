package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	// Alice
	go func() {
		Deposit(200)
		WithdrawIt(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(50)
		WithdrawIt(50)
		Deposit(100)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Wait for both transactions
	<-done
	<-done

	if got, want := Balance(), 100; got != want {
		t.Errorf("Balance: %d, want: %d", got, want)
	}
}

func TestWithdrawal(t *testing.T) {
	b1 := Balance()
	ok := WithdrawIt(50)
	if !ok {
		t.Errorf("!ok. balance: %d", Balance())
	}
	expects := b1 - 50
	if b2 := Balance(); b2 != expects {
		t.Errorf("balance: %d, want: %d", b2, expects)
	}
}

func TestWithdrawalInsufficient(t *testing.T) {
	b1 := Balance()
	ok := WithdrawIt(b1 + 1) // there must be false, else balance here is wrong, and go continue in teller()
	b2 := Balance()
	if ok {
		t.Errorf("!ok. balance: %d", b2)
	}
	if b2 != b1 {
		t.Errorf("balance: %d, want: %d", b2, b1)
	}
}
