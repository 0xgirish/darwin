package main

import (
	"flag"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var ConfigPath string

func init() {
	flag.StringVar(&ConfigPath, "config", "config/default.yaml", "path to config file for GA parameters")
	flag.Parse()
}

// Config for GA parameters
type Config struct {
	TopK           int     `yaml:"topk"`
	MutationProb   float64 `yaml:"mutation"`
	CrossoverProb  float64 `yaml:"crossover"`
	PopulationSize uint    `yaml:"population"`
}

// NewConfig from ConfigPath
func NewConfig() (c Config, err error) {
	// open file for decoding
	file, err := os.Open(ConfigPath)
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
		`Config {
    TopK:           %v,
    MutationProb:   %v,
    CrossoverProb:  %v,
    PopulationSize: %v,
}`, c.TopK, c.MutationProb, c.CrossoverProb, c.PopulationSize)
}
