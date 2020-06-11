package main

import "flag"

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config/default.yaml", "path to config file for GA parameters")
	flag.Parse()
}
