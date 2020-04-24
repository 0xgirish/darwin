package darwin

import (
	"log"
	"sort"
)

// Population of the chromosomes in the environment
type Population struct {
	env         Env
	size        uint
	chromosomes []Chromosome
}

// NewPopulation of chromosomes in the env
func NewPopulation(chromosomes []Chromosome, env Env) Population {
	size := len(chromosomes)
	return Population{
		env:         env,
		size:        uint(size),
		chromosomes: chromosomes,
	}
}

// Select topK fittest chromosomes in the environment
func (p Population) Select(topK uint) []Chromosome {
	if topK > p.size {
		log.Printf("topK = %v, is greater than population size %v", topK, p.size)
		topK = p.size
	}

	// calculate fitness of all the chromosomes in env
	_cwf := make(chan chromosomeWithFitness, p.size)
	for i := 0; i < int(p.size); i++ {
		go p.fitness(i, _cwf)
	}

	cwf := make(chromosomesWithFitness, p.size)
	for i := 0; i < int(p.size); i++ {
		cwf[i] = <-_cwf
	}

	// sort chromosomes in descending order with fitness
	sort.Sort(sort.Reverse(cwf))
	return cwf.chromosomes()[:topK]
}

func (p Population) fitness(i int, _cwf chan<- chromosomeWithFitness) {
	fitness := p.env.Fit(p.chromosomes[i])
	_cwf <- chromosomeWithFitness{chromosome: p.chromosomes[i], fitness: fitness}
}
