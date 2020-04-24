package darwin

import (
	"sort"
	"testing"
)

type mockChromosome struct {
	name string
}

func newMockChromosome(name string) mockChromosome {
	return mockChromosome{name}
}

func (mc mockChromosome) Crossover(ch Chromosome) Chromosome {
	return mockChromosome{mc.name + "_crossed"}
}

func (mc mockChromosome) Mutation(mutationProb float64) Chromosome {
	return mockChromosome{mc.name + "_mutated"}
}

func TestSortChromosome(t *testing.T) {
	cwfs := chromosomesWithFitness{
		chromosomeWithFitness{newMockChromosome("ch1"), 1.0},
		chromosomeWithFitness{newMockChromosome("ch2"), 10.0},
		chromosomeWithFitness{newMockChromosome("ch3"), 5.0},
		chromosomeWithFitness{newMockChromosome("ch4"), 1.5},
		chromosomeWithFitness{newMockChromosome("ch5"), 23.0},
		chromosomeWithFitness{newMockChromosome("ch6"), 11.0},
	}

	results := []mockChromosome{
		newMockChromosome("ch5"),
		newMockChromosome("ch6"),
		newMockChromosome("ch2"),
		newMockChromosome("ch3"),
		newMockChromosome("ch4"),
		newMockChromosome("ch1"),
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
