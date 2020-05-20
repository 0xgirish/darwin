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
	Mutate func(p Population) Population
	// Corossover generated new populations with crossover
	Crossover func(p Population) Population
}

// Iterate does a single iteration of selection, corssover and mutation on the population
func (ga GeneticAlgorithm) Iterate(p Population, env Env) Population {
	elitePopulation := NewPopulation(p.Select(ga.TopK, env))
	offspring := ga.Crossover(elitePopulation)
	newGeneration := ga.Mutate(offspring)

	return newGeneration
}
