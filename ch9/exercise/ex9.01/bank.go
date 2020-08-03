package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawals:
			if w.amount > balance {
				w.success <- false
				continue
			}
			balance -= w.amount
			w.success <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

type Withdrawl struct {
	amount  int
	success chan bool
}

var withdrawals = make(chan Withdrawl)

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- Withdrawl{amount, ch}
	return <-ch
}
