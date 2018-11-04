package tmod3

import (
	"tmod"
)

type task1 struct{}

type thisStructForTask1 struct {
	Name string `json:"name"`
}

func (t task1) ModRun(c chan interface{}) {
	c <- doTask1()
	return
}

func doTask1() (td thisStructForTask1) {
	td.Name = "Tmod 3 Task 1"
	return
}

func init() {
	var mod task1
	tmod.Register("tmod3task1", mod)
}
