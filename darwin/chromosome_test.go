package darwin

import (
	"sort"
	"testing"
)

type mockChromosome struct {
	name  string
	index int
}

func newMock(name string, index int) mockChromosome {
	return mockChromosome{name, index}
}

func (mc mockChromosome) Crossover(ch Chromosome) Chromosome {
	return mockChromosome{mc.name + "_crossed", mc.index}
}

func (mc mockChromosome) Mutation(mutationProb float64) Chromosome {
	return mockChromosome{mc.name + "_mutated", mc.index}
}

func TestSortChromosome(t *testing.T) {
	cwfs := chromosomesWithFitness{
		chromosomeWithFitness{newMock("ch1", 1), 1.0},
		chromosomeWithFitness{newMock("ch2", 2), 10.0},
		chromosomeWithFitness{newMock("ch3", 3), 5.0},
		chromosomeWithFitness{newMock("ch4", 4), 1.5},
		chromosomeWithFitness{newMock("ch5", 5), 23.0},
		chromosomeWithFitness{newMock("ch6", 6), 11.0},
	}

	results := []mockChromosome{
		newMock("ch5", 5),
		newMock("ch6", 6),
		newMock("ch2", 2),
		newMock("ch3", 3),
		newMock("ch4", 4),
		newMock("ch1", 1),
	}

	sort.Sort(sort.Reverse(cwfs))
	chromosomes := cwfs.chromosomes()
	for i, chromosome := range chromosomes {
		mock, _:= chromosome.(mockChromosome)
		if results[i] != mock {
			t.Fatalf("%v failed [%v]: expected %v got %v", t.Name(), i+1, results[i], mock)
		}
	}
}
