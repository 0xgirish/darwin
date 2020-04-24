package darwin

// Env is environment in which chromosomes fitness is evaluated
type Env interface {
	// Fit reports fitness of the chromosome in the environment
	//
	// e.g. Fit can start a game simulation using chromosome
	// and report score as fitness
	Fit(ch Chromosome) float64
}
