package models

type Customer struct {
	ID			int64
	Username	string
	Passwd		string
	PasswdHash	string
	Balance		int
	Tasks		[]Task
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (c * Customer) UpdateBalance(cost int) {
	c.Balance -= cost
}
