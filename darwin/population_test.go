package darwin

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestPopulationSelect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	chromosomes := []Chromosome{
		mockChromosome{"ch1"},
		mockChromosome{"ch2"},
		mockChromosome{"ch3"},
		mockChromosome{"ch4"},
		mockChromosome{"ch5"},
		mockChromosome{"ch6"},
	}

	results := []Chromosome{
		mockChromosome{"ch5"},
		mockChromosome{"ch6"},
		mockChromosome{"ch2"},
		mockChromosome{"ch3"},
		mockChromosome{"ch4"},
		mockChromosome{"ch1"},
	}

	env := NewMockEnv(ctrl)
	env.EXPECT().Fit(chromosomes[0]).Return(1.0, struct{}{}).Times(1)
	env.EXPECT().Fit(chromosomes[1]).Return(10.0, struct{}{}).Times(1)
	env.EXPECT().Fit(chromosomes[2]).Return(5.0, struct{}{}).Times(1)
	env.EXPECT().Fit(chromosomes[3]).Return(1.5, struct{}{}).Times(1)
	env.EXPECT().Fit(chromosomes[4]).Return(23.0, struct{}{}).Times(1)
	env.EXPECT().Fit(chromosomes[5]).Return(11.0, struct{}{}).Times(1)

	population := NewPopulation(chromosomes)
	selection := population.Select(uint(len(chromosomes))+1, env, RankSelection)
	for i := 0; i < len(chromosomes); i++ {
		if results[i] != selection[i] {
			t.Errorf("%v failed [%v]: expected %v, got %v", t.Name(),
				i+1, results[i], selection[i])
		}
	}
}
