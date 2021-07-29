// https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial
package funding

type FundServer struct {
	commands chan TransactionCommand
	fund     Fund
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}

type Transactor func(f *Fund)

type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

func (s *FundServer) Transact(transactor Transactor) {
	command := TransactionCommand{
		transactor,
		make(chan bool),
	}
	s.commands <- command
	<-command.Done
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan TransactionCommand),
		fund:     *NewFund(initialBalance),
	}

	go server.loop()
	return server
}

func (s *FundServer) Balance() int {
	var balance int
	s.Transact(func(f *Fund) {
		balance = f.Balance()
	})
	return balance
}

func (s *FundServer) Withdrawal(amount int) {
	s.Transact(func(f *Fund) {
		f.Withdraw(amount)
	})
}

func (s *FundServer) loop() {
	for transaction := range s.commands {
		transaction.Transactor(&s.fund)
		transaction.Done <- true
	}
}
