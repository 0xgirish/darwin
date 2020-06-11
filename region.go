package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

// Shop coordinates in the map
type Shop struct {
	Lat float64 `csv:"lat"`
	Lon float64 `csv:"lon"`
}

// Shops return shops from csv file
func Shops(region string) (shops []*Shop, err error) {
	// open shops.csv file for corresponding region
	csvfile, err := os.Open(fmt.Sprintf("csv/%s/shops.csv", region))
	if err != nil {
		return
	}
	defer csvfile.Close()

	err = gocsv.UnmarshalFile(csvfile, &shops)
	return
}

// Head is equvalent to padas head
func Head(shops []*Shop) {
	size := len(shops)

	fmt.Println("  lat, lon")
	for i := 0; i < 4 && i < size; i++ {
		fmt.Printf("%d %f, %f\n", i+1, shops[i].Lat, shops[i].Lon)
	}
}
