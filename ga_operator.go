package main

import (
	"github.com/zkmrgirish/gobtp/darwin"
)

// Crossover of elite chromosomes and create nchilds
func Crossover(p []darwin.Chromosome, prob float64, nchilds uint) []darwin.Chromosome {
	size := len(p)
	children := make([]darwin.Chromosome, nchilds)

	for i := 0; i < int(nchilds); i++ {
		r1, r2 := darwin.RandIntn(size), darwin.RandIntn(size)
		parent1, parent2 := p[r1], p[r2]
		children[i] = parent1.Crossover(parent2, prob)
	}
	return children
}

// Mutate children from crossover to take over the world
func Mutate(p []darwin.Chromosome, prob float64) []darwin.Chromosome {
	mutants := make([]darwin.Chromosome, len(p))

	for i, ch := range p {
		mutants[i] = ch.Mutation(prob)
	}
	return mutants
}
