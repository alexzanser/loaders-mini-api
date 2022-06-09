package models

import (
	"math/rand"
	"time"
)

type Loader struct {
	ID				int64
	Username		string
	Passwd			string	`json:"password,omitempty"`
	PasswdHash		string	`json:"password_hash,omitempty"`
	MaxWeight		int
	Drunk			bool
	Fatigue			int		`json:"fatigue,omitempty"`
	Salary			int
	Balance			int
	CompletedTasks	[]Task
}

func NewLoader() *Loader {
	return &Loader{
		MaxWeight: randomLoaderWeight(),
		Drunk: drunk(),
		Fatigue: randomFatigue(),
		Balance: 0,
		Salary: randomSalary(),
		CompletedTasks: make([]Task, 0),
	}
}

var MinLoaderWeight, MaxLoaderWeight = 5, 30
var MinSalary, MaxSalary = 10000, 30000

func randomLoaderWeight() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MaxLoaderWeight - MinLoaderWeight) + MinLoaderWeight
}

func drunk() bool {
	if time.Now().UnixNano() % 2 == 0 {
		return true
	}
	return false
}

func randomFatigue() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func randomSalary() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MaxSalary - MinSalary) + MinSalary
}

func (l *Loader) Update(){
	drunkAffect := 0
	if l.Drunk {
		drunkAffect  += 50 
	}
	l.MaxWeight = l.MaxWeight * (1 - l.Fatigue / 100) * (1 - drunkAffect/ 100)
	if l.MaxWeight < 0 {
		l.MaxWeight = 0
	}
	l.Balance += l.Salary
	l.Fatigue += 20
}