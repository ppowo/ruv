package stations

import (
	"fmt"

	"github.com/ppowo/ruv/models"
)

// stations is a private slice containing all available radio stations
var stations = []models.Station{
	{
		Name:        "reso",
		Description: "Resonance FM - London-based art radio station",
		URL:         "https://stream.resonance.fm/resonance",
	},
	{
		Name:        "rese",
		Description: "Resonance Extra - Alternative Resonance FM stream",
		URL:         "https://stream.resonance.fm/resonance-extra",
	},
	{
		Name:        "ntso",
		Description: "NTS Radio 1 - Global online radio station",
		URL:         "https://streams.radiomast.io/nts1/hls.m3u8",
	},
	{
		Name:        "ntst",
		Description: "NTS Radio 2 - Global online radio station",
		URL:         "https://streams.radiomast.io/nts2/hls.m3u8",
	},
	{
		Name:        "lyll",
		Description: "LYL Radio - Community radio from London",
		URL:         "https://radio.lyl.live/hls/aac_hifi.m3u8",
	},
	{
		Name:        "cash",
		Description: "Cashmere Radio - Berlin-based online radio",
		URL:         "https://cashmereradio.out.airtime.pro/cashmereradio_b",
	},
	{
		Name:        "lake",
		Description: "The Lake Radio - Online radio station",
		URL:         "http://hyades.shoutca.st:8627/stream",
	},
	{
		Name:        "alha",
		Description: "Radio Alhara - Palestinian online radio",
		URL:         "https://stream.radiojar.com/78cxy6wkxtzuv",
	},
}

// GetStations returns a copy of all available stations
func GetStations() []models.Station {
	// Return a copy to prevent external modification
	result := make([]models.Station, len(stations))
	copy(result, stations)
	return result
}

// GetStation looks up a station by name (case-insensitive)
func GetStation(name string) (*models.Station, error) {
	if name == "" {
		return nil, fmt.Errorf("station name cannot be empty")
	}

	// Search for station
	for i, station := range stations {
		if station.Name == name {
			return &stations[i], nil
		}
	}

	return nil, fmt.Errorf("station %q not found", name)
}
