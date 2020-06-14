package main

import (
	"fmt"

	"github.com/zkmrgirish/gobtp/darwin"
)

// InstanceN for btp
type InstanceN struct {
	ncircles      uint
	mutationRange float64
	lats          []float64
	lons          []float64
	coordRange    [2][2]float64
}

// NewInstanceN with coordRange as range of lats and lons
func NewInstanceN(maxNcircles uint, mutationRange float64, coordRange [2][2]float64) InstanceN {
	ncircles := uint(8) + uint(darwin.RandIntn(int(maxNcircles-7)))

	ins := InstanceN{ncircles: ncircles, mutationRange: mutationRange, coordRange: coordRange}
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
func (ins InstanceN) Crossover(ch darwin.Chromosome, crossoverProb float64) darwin.Chromosome {
	mate, _ := ch.(InstanceN)

	// with corssoverProb do corssover of ins and mate
	if darwin.Flip(crossoverProb) {
		// shuffle both instance and mate coordinates
		darwin.RandSource.Shuffle(int(ins.ncircles), func(i, j int) {
			ins.lats[i], ins.lats[j] = ins.lats[j], ins.lats[i]
			ins.lons[i], ins.lons[j] = ins.lons[j], ins.lons[i]
		})

		darwin.RandSource.Shuffle(int(mate.ncircles), func(i, j int) {
			mate.lats[i], mate.lats[j] = mate.lats[j], mate.lats[i]
			mate.lons[i], mate.lons[j] = mate.lons[j], mate.lons[i]
		})

		child, _ := ins.Copy().(InstanceN)
		crossover_point := darwin.RandIntn(int(ins.ncircles))
		for i := crossover_point; i < int(ins.ncircles); i++ {
			child.lats[i], child.lons[i] = mate.lats[i], mate.lons[i]
		}

		// add some random points to the child with ncircles = random(ins.size, mate.size)
		new_ncircles := ins.ncircles + uint(darwin.RandIntn(int(mate.ncircles-ins.ncircles+1)))
		dlat := (child.coordRange[0][1] - child.coordRange[0][0])
		dlon := (child.coordRange[1][1] - child.coordRange[1][0])
		for i := child.ncircles; i < new_ncircles; i++ {
			child.lats = append(child.lats, darwin.Random()*dlat+child.coordRange[0][0])
			child.lons = append(child.lons, darwin.Random()*dlon+child.coordRange[0][0])
		}
		child.ncircles = new_ncircles

		return child
	}

	// return ins or mate with 1/2 probability
	if darwin.Flip(0.5) {
		return ins
	}
	return mate
}

// Mutation of the instance
func (ins InstanceN) Mutation(mutationProb float64) darwin.Chromosome {

	// create a new mutant like Logan
	mutant, _ := ins.Copy().(InstanceN)
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
func (ins InstanceN) Copy() darwin.Chromosome {
	lats := make([]float64, ins.ncircles)
	lons := make([]float64, ins.ncircles)

	// copy lats and lons
	copy(lats, ins.lats)
	copy(lons, ins.lons)

	return InstanceN{
		ncircles:      ins.ncircles,
		mutationRange: ins.mutationRange,
		lats:          lats,
		lons:          lons,
	}
}

// String makes instance printable
func (ins InstanceN) String() string {
	lats := fmt.Sprintf("Lats: %v", ins.lats)
	lons := fmt.Sprintf("Lons: %v", ins.lons)
	return fmt.Sprintf("%s\n%s", lats, lons)
}
