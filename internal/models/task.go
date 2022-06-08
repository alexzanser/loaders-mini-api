package models

import (
	"math/rand"
	"time"
)

type Task struct {
	ID		int64
	Name	string
	Weight	int
}

func NewTask() *Task {
	return &Task{
		Name: randomTask(),
		Weight: randomTaskWeight(),
	}
}

var TaskNames = []string{
	"Furniture",
	"Clothes",
	"Tools", 
	"Appliances",
	"Food", 
	"Auto parts",
}

var MinTaskWeight, MaxTaskWeight = 10, 80

func randomTask() string {
	rand.Seed(time.Now().UnixNano())
	return  TaskNames[rand.Intn(len(TaskNames))]
}

func randomTaskWeight() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MaxTaskWeight - MinTaskWeight) + MinTaskWeight
}
