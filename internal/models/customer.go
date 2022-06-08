package models

type Customer struct {
	Username	string
	Passwd		string
	PasswdHash	string
	Balance		int
	Tasks		[]Task
}

func NewCustomer() *Customer {
	return &Customer{}
}
