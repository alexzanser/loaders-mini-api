package models

import (
	"math/rand"
	"time"
)

type Loader struct {
	ID				int64
	Username		string
	Passwd			string
	PasswdHash		string
	MaxWeight		int
	Drunk			bool
	Fatigue			int
	Salary			int
	Balance			int
	CompletedTasks	[]Task
}

func NewLoader() *Loader {
	return &Loader{
		MaxWeight: randomLoaderWeight(),
		Drunk: drunk(),
		Fatigue: randomFatigue(),
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

func (l *Loader) UpdateWeight() {
	drunkAffect := 0
	if l.Drunk {
		drunkAffect  += 50 
	}
	l.MaxWeight = l.MaxWeight * (100 - l.Fatigue / 100) * (100 - drunkAffect/ 100)
}