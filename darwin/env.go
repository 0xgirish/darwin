package darwin

// MetaData from env.Fit
type MetaData interface{}

// Env is environment in which chromosomes fitness is evaluated
type Env interface {
	// Fit reports fitness of the chromosome in the environment
	//
	// e.g. Fit can start a game simulation using chromosome
	// and report score as fitness, and some metadata
	Fit(ch Chromosome) (float64, MetaData)
}
