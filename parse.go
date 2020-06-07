package main

import "flag"

var (
	// command line options
	configPath string
	region     string
)

func init() {
	flag.StringVar(&configPath, "config", "config/default.yaml", "path to config file for GA parameters")
	flag.StringVar(&region, "region", "dlh", "name of the region to process data, e.g. blr, dlh, cdg")
	flag.Parse()
}
