package tmod2

import (
	"tmod"
)

type task2 struct{}

type thisStructForTask2 struct {
	Name string `json:"name"`
}

func (t task2) ModRun(c chan interface{}) {
	c <- doTask2()
	return
}

func doTask2() (td thisStructForTask2) {
	td.Name = "Tmod 2 Task 2"
	return
}

func init() {
	var mod task2
	tmod.Register("tmod2task2", mod)
}
