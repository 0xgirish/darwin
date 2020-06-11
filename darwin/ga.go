// Package darwin provides basic functionalities
// and interface for genetic algorithm implementation
package darwin

// GeneticAlgorithm implementation
type GeneticAlgorithm struct {
	// TopK selection of fittest chromosomes in population
	TopK uint
	// MutationProb is mutation probability for chromosome
	// should be in range (0, 1]
	MutationProb float64
	// CrossoverProb is crossover probability for chromosome
	// should be in range (0, 1]
	CrossoverProb float64
	// PopulationSize is number of chromosomes in the population
	PopulationSize uint
	// Mutate generate mutated population
	Mutate func(p []Chromosome, prob float64) []Chromosome
	// Corossover generated new populations with crossover
	Crossover func(p []Chromosome, prob float64, nchilds uint) []Chromosome
}

// Iterate does a single iteration of selection, corssover and mutation on the population
func (ga GeneticAlgorithm) Iterate(p Population, env Env, selection SelectionMethod) Population {
	eliteChromosomes := p.Select(ga.TopK, env, selection)

	// create nchilds using crossover
	// nchilds := uint(p.size - ga.TopK)
	offspring := ga.Crossover(eliteChromosomes, ga.CrossoverProb, p.size)

	// newChromosomes := append(eliteChromosomes, offspring...)
	newGeneration := ga.Mutate(offspring, ga.MutationProb)

	return NewPopulation(newGeneration)
}
