package main

import "fmt"

func main() {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
	// TODO: create env for ga
	// TODO: create mutation and crossover funcs
}
