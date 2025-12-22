package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/ppowo/ruv/stations"
)

func runList() error {
	stationList := stations.GetStations()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Name", "Description"})
	for _, station := range stationList {
		table.Append([]string{station.Name, station.Description})
	}
	table.Footer([]string{"Total", fmt.Sprintf("%d station(s)", len(stationList))})
	table.Render()

	return nil
}
