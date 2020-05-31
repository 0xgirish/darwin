package main

import "fmt"

func main() {
	config := NewConfig()
	config.Parse()

	fmt.Println(config)
	// TODO: create env for ga
	// TODO: create mutation and crossover funcs
}
