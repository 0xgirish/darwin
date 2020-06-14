package darwin

import "fmt"

// Chromosome represents an solution candidate of problem
type Chromosome interface {
	// Crossover create new chromosome using crossover with ch chromosome
	Crossover(ch Chromosome, crossoverPorb float64) Chromosome

	// Mutation mutates current chromosome with some mutation probability
	Mutation(mutationProb float64) Chromosome

	// Copy chromosome to another
	Copy() Chromosome

	// String makes chromosome printable
	String() string
}

// cwfpair is a pair of chromosome and its fitness in the environment
type cwfpair struct {
	chromosome Chromosome
	fitness    float64
	metaData   MetaData
}

type cwfpairs []cwfpair

// Len is the number of chromosomes in the collection
func (cwf cwfpairs) Len() int {
	return len(cwf)
}

// Less reports whether the chromosome with
// ith index is less fitter than jth index chromosome.
func (cwf cwfpairs) Less(i, j int) bool {
	return cwf[i].fitness < cwf[j].fitness
}

// Swap swaps two chromosomes with index i, j
func (cwf cwfpairs) Swap(i, j int) {
	cwf[i], cwf[j] = cwf[j], cwf[i]
}

func (cwf cwfpairs) Print() {
	for i := 0; i < cwf.Len(); i++ {
		fmt.Printf("%f ", cwf[i].fitness)
	}
	fmt.Println()
}

// extract chromosomes from cwfpairs
func (cwf cwfpairs) chromosomes() []Chromosome {
	size := cwf.Len()
	chromosomes := make([]Chromosome, size)
	for i := 0; i < size; i++ {
		chromosomes[i] = cwf[i].chromosome
	}
	return chromosomes
}
