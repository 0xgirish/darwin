package main

import (
	"github.com/zkmrgirish/gobtp/darwin"
)

type meta struct {
	Covered   float64
	Overlaped float64
}

// Env for btp
type Env struct {
	radius float64
	shops  []*Shop
}

// Fit reports fitness of chromosome and metadata
func (env *Env) Fit(ch darwin.Chromosome) (float64, darwin.MetaData) {
	nshops := len(env.shops)

	ins, _ := ch.(InstanceN)
	covered := make([]uint, nshops)
	for c := 0; c < int(ins.ncircles); c++ {
		env.coverage(ins.lats[c], ins.lons[c], covered)
	}

	coveredNodes, overlapedNodes := 0, 0
	for _, coverage := range covered {
		if coverage != 0 {
			coveredNodes++
		}

		if coverage > 1 {
			overlapedNodes = overlapedNodes + int(coverage) - 1
		}
	}

	coveredPercentage := float64(100*coveredNodes) / float64(nshops)
	overlapedPercentage := float64(100*overlapedNodes) / float64(nshops)

	fitness := (coveredPercentage - 8*overlapedPercentage)
	// fitness := coveredPercentage / (overlapedPercentage + 1)

	metadata := meta{Covered: coveredPercentage, Overlaped: overlapedPercentage}
	return fitness, metadata
}

func (env *Env) coverage(lat, lon float64, covered []uint) {
	nshops := len(env.shops)

	for i := 0; i < nshops; i++ {
		slat, slon := env.shops[i].Lat, env.shops[i].Lon
		if Geodisc(lat, lon, slat, slon) <= env.radius {
			covered[i]++
		}
	}
}
