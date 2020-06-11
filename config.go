package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

// Config for GA parameters
type Config struct {
	GA struct {
		SelectionPercentage float64   `yaml:"selection_percentage"`
		MutationProb        []float64 `yaml:"mutation_prob"`
		CrossoverProb       []float64 `yaml:"crossover_prob"`
		PopulationSize      []uint    `yaml:"population"`
	} `yaml:"ga"`
	Extra struct {
		Iterations    uint      `yaml:"iterations"`
		Region        string    `yaml:"region"`
		LatsRange     []float64 `yaml:"lats_range"`
		LonsRange     []float64 `yaml:"lons_range"`
		Radius        float64   `yaml:"radius"`
		NCircles      []uint    `yaml:"n_circles"`
		MutationRange []float64 `yaml:"mutation_range"`
	} `yaml:"extra"`
}

// NewConfig from configPath
func NewConfig() (c Config, err error) {
	// open file for decoding
	file, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer file.Close()

	// unmarshal yaml to config
	err = yaml.NewDecoder(file).Decode(&c)
	return
}

func (c Config) String() string {
	return fmt.Sprintf(
		`Config
	ga:
		selection_percentage: %v
		mutation_prob: %v
		crossover_prob: %v
		population: %v
	extra:
		iterations: %v
		region: %v
		lats_range: %v
		lons_range: %v
		radius: %v
		n_circles: %v
		mutation_range: %v
	`, c.GA.SelectionPercentage, c.GA.MutationProb, c.GA.CrossoverProb, c.GA.PopulationSize,
		c.Extra.Iterations, c.Extra.Region, c.Extra.LatsRange, c.Extra.LonsRange,
		c.Extra.Radius, c.Extra.NCircles, c.Extra.MutationRange)
}
