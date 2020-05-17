package darwin

import (
	"sort"
	"testing"
)

// mockChromosome for testing
type mockChromosome struct {
	name string
}

// Crossover create new chromosome using crossover with ch chromosome
func (mc mockChromosome) Crossover(ch Chromosome, crossoverPorb float64) Chromosome {
	return mockChromosome{mc.name + "_corssed"}
}

// Mutation mutates current chromosome with some mutation probability
func (mc mockChromosome) Mutation(mutationProb float64) Chromosome {
	return mockChromosome{mc.name + "_mutated"}
}

func TestSortChromosome(t *testing.T) {
	cwfs := cwfpairs{
		cwfpair{mockChromosome{"ch1"}, 1.0},
		cwfpair{mockChromosome{"ch2"}, 10.0},
		cwfpair{mockChromosome{"ch3"}, 5.0},
		cwfpair{mockChromosome{"ch4"}, 1.5},
		cwfpair{mockChromosome{"ch5"}, 23.0},
		cwfpair{mockChromosome{"ch6"}, 11.0},
	}

	results := []mockChromosome{
		mockChromosome{"ch5"},
		mockChromosome{"ch6"},
		mockChromosome{"ch2"},
		mockChromosome{"ch3"},
		mockChromosome{"ch4"},
		mockChromosome{"ch1"},
	}

	sort.Sort(sort.Reverse(cwfs))
	chromosomes := cwfs.chromosomes()
	for i, chromosome := range chromosomes {
		mock, _ := chromosome.(mockChromosome)
		if results[i] != mock {
			t.Fatalf("%v failed [%v]: expected %v got %v", t.Name(), i+1, results[i], mock)
		}
	}
}
