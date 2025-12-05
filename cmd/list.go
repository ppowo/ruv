package cmd

import (
	"fmt"
	"strings"

	"github.com/ppowo/ruv/stations"
)

func runList() error {
	// Get all stations from registry
	stationList := stations.GetStations()

	// Print header
	fmt.Println("Available Radio Stations:")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Printf("%-20s | %s\n", "NAME", "DESCRIPTION")
	fmt.Println(strings.Repeat("-", 70))

	// Print each station
	for _, station := range stationList {
		fmt.Printf("%-20s | %s\n",
			station.Name,
			station.Description)
	}

	fmt.Println(strings.Repeat("-", 70))
	fmt.Printf("\nTotal: %d station(s)\n", len(stationList))

	return nil
}
