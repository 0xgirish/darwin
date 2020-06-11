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

// SelectionMethod for selection operator ga
type SelectionMethod func(cwfs cwfpairs, topK uint) []Chromosome

// Select topK fittest chromosomes in the environment
func (p Population) Select(topK uint, env Env, selection SelectionMethod) []Chromosome {
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

	return selection(cwf, topK)
}

// Fittest report fittest chromosome int the environment
func (p Population) Fittest(env Env) Chromosome {
	return p.Select(1, env, RankSelection)[0]
}

// WheelSelection implement Roulette wheel selection algorithm
func WheelSelection(cwfs cwfpairs, topK uint) []Chromosome {
	sort.Sort(cwfs)
	popsize, sumfitness := len(cwfs), 0.0
	for i := 0; i < popsize; i++ {
		sumfitness += cwfs[i].fitness
	}

	chromosomes := make([]Chromosome, topK)
	for i := 0; i < int(topK); i++ {
		rand, partsum := Random()*sumfitness, 0.0
		for j := 0; j < popsize; j++ {
			partsum += cwfs[j].fitness
			if partsum >= rand {
				chromosomes[i] = cwfs[j].chromosome
				break
			}
		}
	}
	return chromosomes
}

// RankSelection select topK chromosomes according to fitness
func RankSelection(cwfs cwfpairs, topK uint) []Chromosome {
	sort.Sort(sort.Reverse(cwfs))
	return cwfs.chromosomes()[:topK]
}

// get fitness of the chromosome with index i relative to the environment
func (p Population) fitness(i int, env Env, _cwf chan<- cwfpair) {
	fitness, _ := env.Fit(p.Chromosomes[i]) // ignore metadata

	log.Printf("chromosome: %v, fitness: %v", i, fitness)
	_cwf <- cwfpair{chromosome: p.Chromosomes[i], fitness: fitness}
}
