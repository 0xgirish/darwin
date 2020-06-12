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
	for _, ncircles := range config.Extra.NCircles {
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
							chromosomes[i] = NewInstance(ncircles, mrange, coordRange)
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
							ncircles, popsize, cprob, mprob, mrange, selectionPercentage,
							fitness, mdata.Covered, mdata.Overlaped)
						fmt.Println(fittest)
					}
				}
				fmt.Printf("-------------------------------------------\n")
			}
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

// Instance for btp
type Instance struct {
	ncircles      uint
	mutationRange float64
	lats          []float64
	lons          []float64
}

// NewInstance with coordRange as range of lats and lons
func NewInstance(ncircles uint, mutationRange float64, coordRange [2][2]float64) Instance {
	ins := Instance{ncircles: ncircles, mutationRange: mutationRange}
	ins.lats, ins.lons = make([]float64, ncircles), make([]float64, ncircles)

	dlat := (coordRange[0][1] - coordRange[0][0])
	dlon := (coordRange[1][1] - coordRange[1][0])
	for i := 0; i < int(ncircles); i++ {
		ins.lats[i] = darwin.Random()*dlat + coordRange[0][0]
		ins.lons[i] = darwin.Random()*dlon + coordRange[1][0]
	}

	return ins
}

// Crossover of the instance with ch mate
func (ins Instance) Crossover(ch darwin.Chromosome, crossoverProb float64) darwin.Chromosome {
	mate, _ := ch.(Instance)

	// with corssoverProb do corssover of ins and mate
	if darwin.Flip(crossoverProb) {
		child, _ := ins.Copy().(Instance)
		for i := 0; i < int(ins.ncircles); i++ {
			// choose mates gnome with 1/2 probability
			if darwin.Flip(crossoverProb) {
				child.lats[i], child.lons[i] = mate.lats[i], mate.lons[i]
			}
		}
		return child
	}

	// return ins or mate with 1/2 probability
	if darwin.Flip(0.5) {
		return ins
	}
	return mate
}

// Mutation of the instance
func (ins Instance) Mutation(mutationProb float64) darwin.Chromosome {

	// create a new mutant like Logan
	mutant, _ := ins.Copy().(Instance)
	for i := 0; i < int(mutant.ncircles); i++ {
		if darwin.Flip(mutationProb) {
			mutatelat := 2 * (darwin.Random() - 0.5) * mutant.mutationRange
			mutatelon := 2 * (darwin.Random() - 0.5) * mutant.mutationRange
			mutant.lats[i] = mutant.lats[i] + mutatelat
			mutant.lons[i] = mutant.lons[i] + mutatelon
		}
	}
	return mutant
}

// Copy instance to another
func (ins Instance) Copy() darwin.Chromosome {
	lats := make([]float64, ins.ncircles)
	lons := make([]float64, ins.ncircles)

	// copy lats and lons
	copy(lats, ins.lats)
	copy(lons, ins.lons)

	return Instance{
		ncircles:      ins.ncircles,
		mutationRange: ins.mutationRange,
		lats:          lats,
		lons:          lons,
	}
}

// String makes instance printable
func (ins Instance) String() string {
	lats := fmt.Sprintf("Lats: %v", ins.lats)
	lons := fmt.Sprintf("Lons: %v", ins.lons)
	return fmt.Sprintf("%s\n%s", lats, lons)
}
