package tmod

import (
	"fmt"
	"log"
)

// modular
var modulars = make(map[string]Modular)

type Modular interface {
	ModRun(c chan interface{})
}

func Run() {
	if len(modulars) <= 0 {
		log.Fatal("No Modular Registered")
		return
	}

	c := make(chan interface{})
	for _, v := range modulars {
		go v.ModRun(c)
	}

	for i := 0; i < len(modulars); i++ {
		fmt.Println(<-c)
	}
}

func Register(modName string, modular Modular) {
	if _, ok := modulars[modName]; ok {
		log.Fatalln(modName, "Modular already registered")
		return
	}

	log.Println("Register", modName, "Modular")
	modulars[modName] = modular
}
