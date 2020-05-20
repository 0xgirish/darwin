package darwin

import (
	"log"
	"sort"
)

// Population of the chromosomes in the environment
type Population struct {
	size        uint
	Chromosomes []Chromosome
}

// NewPopulation of chromosomes in the env
func NewPopulation(chromosomes []Chromosome) Population {
	size := len(chromosomes)
	return Population{
		size:        uint(size),
		Chromosomes: chromosomes,
	}
}

// Select topK fittest chromosomes in the environment
func (p Population) Select(topK uint, env Env) []Chromosome {
	if topK > p.size {
		log.Printf("topK = %v, is greater than population size %v", topK, p.size)
		topK = p.size
	}

	// calculate fitness of all the chromosomes in env
	_cwf := make(chan cwfpair, p.size)
	for i := 0; i < int(p.size); i++ {
		go p.fitness(i, env, _cwf)
	}

	cwf := make(cwfpairs, p.size)
	for i := 0; i < int(p.size); i++ {
		cwf[i] = <-_cwf
	}

	// sort chromosomes in descending order with fitness
	sort.Sort(sort.Reverse(cwf))
	return cwf.chromosomes()[:topK]
}

// get fitness of the chromosome with index i relative to the environment
func (p Population) fitness(i int, env Env, _cwf chan<- cwfpair) {
	fitness := env.Fit(p.Chromosomes[i])
	_cwf <- cwfpair{chromosome: p.Chromosomes[i], fitness: fitness}
}
