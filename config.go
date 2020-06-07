package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

// Config for GA parameters
type Config struct {
	GA struct {
		TopK           []int     `yaml:"topk"`
		MutationProb   []float64 `yaml:"mutation_prob"`
		CrossoverProb  []float64 `yaml:"crossover_prob"`
		PopulationSize []uint    `yaml:"population"`
	} `yaml:"ga"`
	Extra struct {
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
		topk: %v
		mutation_prob: %v
		crossover_prob: %v
		population: %v
	extra:
		radius: %v
		n_circles: %v
		mutation_range: %v
	`, c.GA.TopK, c.GA.MutationProb, c.GA.CrossoverProb, c.GA.PopulationSize,
		c.Extra.Radius, c.Extra.NCircles, c.Extra.MutationRange)
}
