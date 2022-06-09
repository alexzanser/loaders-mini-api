package models

import (
	"math/rand"
	"time"
)

//Task представляет структуру для задачи
type Task struct {
	ID		int64
	Name	string
	Weight	int
}

//NewTask создает новую задачу
func NewTask() *Task {
	return &Task{
		Name: randomTask(),
		Weight: randomTaskWeight(),
	}
}

var taskNames = []string{
	"Furniture",
	"Clothes",
	"Tools", 
	"Appliances",
	"Food", 
	"Auto parts",
}

var (
	minTaskWeight, maxTaskWeight = 10, 80
)

func randomTask() string {
	rand.Seed(time.Now().UnixNano())
	return  taskNames[rand.Intn(len(taskNames))]
}

func randomTaskWeight() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxTaskWeight - minTaskWeight) + minTaskWeight
}
