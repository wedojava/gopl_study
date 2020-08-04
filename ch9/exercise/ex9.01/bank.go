package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawals = make(chan Withdrawal)

type Withdrawal struct {
	amount  int
	success chan bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func WithdrawIt(amount int) bool {
	ifsuccess := make(chan bool)
	withdrawals <- Withdrawal{amount, ifsuccess}
	return <-ifsuccess
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case wd := <-withdrawals:
			if wd.amount > balance {
				wd.success <- false
				continue
			}
			balance -= wd.amount
			wd.success <- true
		case dp := <-deposits:
			balance += dp
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
