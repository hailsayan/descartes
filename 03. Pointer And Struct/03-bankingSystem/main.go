package main

type SavingsAccount struct {
	balance int
}

func (sa *SavingsAccount) MonthlyInterest() int {
	if sa.CheckBalance() <= 0 {
		return 0
	}
	return int((float64(sa.CheckBalance())*5)/100) / 12
}

func (sa *SavingsAccount) Transfer(receiver Account, amount int) string {
	return DoTransfer(sa, receiver, amount)
}

func (sa *SavingsAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	sa.balance += amount
	return "Success"
}

func (sa *SavingsAccount) Withdraw(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	if sa.CheckBalance() < amount {
		return "Account balance is not enough"
	}

	sa.balance -= amount
	return "Success"
}

func (sa *SavingsAccount) CheckBalance() int {
	return sa.balance
}

type CheckingAccount struct {
	balance int
}

func (sa *CheckingAccount) MonthlyInterest() int {
	if sa.CheckBalance() <= 0 {
		return 0
	}
	return int((float64(sa.CheckBalance())*2)/100) / 12
}

func (sa *CheckingAccount) Transfer(receiver Account, amount int) string {
	return DoTransfer(sa, receiver, amount)
}

func (sa *CheckingAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	sa.balance += amount
	return "Success"
}

func (sa *CheckingAccount) Withdraw(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	if sa.CheckBalance() < amount {
		return "Account balance is not enough"
	}

	sa.balance -= amount
	return "Success"
}

func (sa *CheckingAccount) CheckBalance() int {
	return sa.balance
}

type InvestmentAccount struct {
	balance int
}

func (sa *InvestmentAccount) MonthlyInterest() int {
	if sa.CheckBalance() <= 0 {
		return 0
	}
	return int((float64(sa.CheckBalance())*1)/100) / 12
}

func (sa *InvestmentAccount) Transfer(receiver Account, amount int) string {
	return DoTransfer(sa, receiver, amount)
}

func (sa *InvestmentAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	sa.balance += amount
	return "Success"
}

func (sa *InvestmentAccount) Withdraw(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	if sa.CheckBalance() < amount {
		return "Account balance is not enough"
	}

	sa.balance -= amount
	return "Success"
}

func (sa *InvestmentAccount) CheckBalance() int {
	return sa.balance
}

func DoTransfer(s Account, r Account, amount int) string {
	switch r.(type) {
	case *SavingsAccount:
	case *InvestmentAccount:
	case *CheckingAccount:
	default:
		return "Invalid receiver account"
	}

	result := s.Withdraw(amount)
	if result != "Success" {
		return result
	}
	result = r.Deposit(amount)
	if result != "Success" {
		return result
	}

	return "Success"
}

func NewSavingsAccount() *SavingsAccount {
	return &SavingsAccount{}
}

func NewCheckingAccount() *CheckingAccount {
	return &CheckingAccount{}
}

func NewInvestmentAccount() *InvestmentAccount {
	return &InvestmentAccount{}
}
