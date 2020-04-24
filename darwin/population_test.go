package darwin

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestPopulationSelect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	chromosomes := []Chromosome{
		newMockChromosome("ch1"),
		newMockChromosome("ch2"),
		newMockChromosome("ch3"),
		newMockChromosome("ch4"),
		newMockChromosome("ch5"),
		newMockChromosome("ch6"),
	}

	results := []Chromosome{
		newMockChromosome("ch5"),
		newMockChromosome("ch6"),
		newMockChromosome("ch2"),
		newMockChromosome("ch3"),
		newMockChromosome("ch4"),
		newMockChromosome("ch1"),
	}

	env := NewmockEnv(ctrl)
	env.EXPECT().Fit(chromosomes[0]).Return(1.0).Times(1)
	env.EXPECT().Fit(chromosomes[1]).Return(10.0).Times(1)
	env.EXPECT().Fit(chromosomes[2]).Return(5.0).Times(1)
	env.EXPECT().Fit(chromosomes[3]).Return(1.5).Times(1)
	env.EXPECT().Fit(chromosomes[4]).Return(23.0).Times(1)
	env.EXPECT().Fit(chromosomes[5]).Return(11.0).Times(1)

	population := NewPopulation(chromosomes, env)
	selection := population.Select(uint(len(chromosomes)) + 1)
	for i := 0; i < len(chromosomes); i++ {
		if results[i] != selection[i] {
			t.Errorf("%v failed [%v]: expected %v, got %v", t.Name(),
				i+1, results[i], selection[i])
		}
	}
}
