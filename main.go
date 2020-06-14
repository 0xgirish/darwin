package main

import (
	"fmt"
	"log"

	"github.com/zkmrgirish/gobtp/darwin"
)

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	shops, err := Shops(config.Extra.Region)
	if err != nil {
		log.Fatal(err)
	}

	coordRange := getCoordRange(config)
	selectionPercentage := config.GA.SelectionPercentage

	env := Env{radius: config.Extra.Radius, shops: shops}
	fmt.Printf("ncircles,popsize,cprob,mprob,mrange,sp,fitness,covered,overlaped\n")
	for _, popsize := range config.GA.PopulationSize {
		topk := uint(float64(popsize) * selectionPercentage)

		for _, cprob := range config.GA.CrossoverProb {
			for _, mprob := range config.GA.MutationProb {
				for _, mrange := range config.Extra.MutationRange {
					ga := darwin.GeneticAlgorithm{
						TopK:           topk,
						MutationProb:   mprob,
						CrossoverProb:  cprob,
						PopulationSize: popsize,
						Mutate:         Mutate,
						Crossover:      Crossover,
					}
					chromosomes := make([]darwin.Chromosome, popsize)
					for i := 0; i < int(popsize); i++ {
						chromosomes[i] = NewInstanceN(uint(14), mrange, coordRange)
					}

					p := darwin.NewPopulation(chromosomes)
					for i := 0; i < int(config.Extra.Iterations); i++ {
						p = ga.Iterate(p, &env, darwin.RankSelection)
						log.Printf("-------------------------------------------\n")
					}
					fittest := p.Fittest(&env)
					fitness, metadata := env.Fit(fittest)
					mdata, _ := metadata.(meta)
					fmt.Printf("%d,%d,%.3f,%.3f,%.3f,%.2f,%.4f,%.4f,%.4f\n",
						fittest.(InstanceN).ncircles, popsize, cprob, mprob, mrange, selectionPercentage,
						fitness, mdata.Covered, mdata.Overlaped)
					fmt.Println(fittest)
				}
			}
			fmt.Printf("-------------------------------------------\n")
		}
	}
}

func getCoordRange(config Config) (coordRange [2][2]float64) {
	coordRange[0][0] = config.Extra.LatsRange[0]
	coordRange[0][1] = config.Extra.LatsRange[1]
	coordRange[1][0] = config.Extra.LonsRange[0]
	coordRange[1][1] = config.Extra.LonsRange[1]
	return
}
