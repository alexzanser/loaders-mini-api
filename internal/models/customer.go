package models

//Customer представляет структуру для заказчика
type Customer struct {
	ID			int64
	Username	string
	Passwd		string
	PasswdHash	string
	Balance		int
	Tasks		[]Task
}

//NewCustomer создает нового заказчика
func NewCustomer() *Customer {
	return &Customer{}
}

//UpdateBalance обновляет баланс после выполнения задания
func (c * Customer) UpdateBalance(cost int) {
	c.Balance -= cost
}
