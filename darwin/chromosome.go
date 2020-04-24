package darwin

// Chromosome represents an solution candidate of problem
type Chromosome interface {
	// Crossover create new chromosome using crossover with ch chromosome
	Crossover(ch Chromosome) Chromosome

	// Mutation mutates current chromosome with some mutation probability
	Mutation(mutationProb float64) Chromosome
}

type chromosomeWithFitness struct {
	chromosome Chromosome
	fitness    float64
}

type chromosomesWithFitness []chromosomeWithFitness

// Len is the number of chromosomes in the collection
func (cwf chromosomesWithFitness) Len() int {
	return len(cwf)
}

// Less reports whether the chromosome with
// ith index is less fitter than jth index chromosome.
func (cwf chromosomesWithFitness) Less(i, j int) bool {
	return cwf[i].fitness < cwf[j].fitness
}

// Swap swaps two chromosomes with index i, j
func (cwf chromosomesWithFitness) Swap(i, j int) {
	cwf[i], cwf[j] = cwf[j], cwf[i]
}

func (cwf chromosomesWithFitness) chromosomes() []Chromosome {
	size := cwf.Len()
	chromosomes := make([]Chromosome, size)
	for i := 0; i < size; i++ {
		chromosomes[i] = cwf[i].chromosome
	}
	return chromosomes
}
