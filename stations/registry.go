package stations

import (
	"fmt"
	"strings"

	"github.com/ppowo/ruv/models"
)

// stations is a private slice containing all available radio stations.
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
		Name:        "nood",
		Description: "Noods Radio - Music-heavy community radio from Bristol",
		URL:         "https://noods-radio.radiocult.fm/stream",
	},
	{
		Name:        "drmm",
		Description: "Intergalactic FM Dream Machine - Experimental music from The Hague",
		URL:         "https://radio.intergalactic.fm/3A",
	},
	{
		Name:        "9128",
		Description: "9128.live - Curated ambient/drone stream with zero talk",
		URL:         "https://streams.radio.co/s0aa1e6f4a/listen",
	},
	{
		Name:        "alha",
		Description: "Radio Alhara - Palestinian online radio",
		URL:         "https://stream.radiojar.com/78cxy6wkxtzuv",
	},
}

// GetStations returns a copy of all available stations.
func GetStations() []models.Station {
	// Return a copy to prevent external modification.
	result := make([]models.Station, len(stations))
	copy(result, stations)
	return result
}

// GetStation looks up a station by name (case-insensitive).
// Returns a copy — mutations do not affect the internal registry.
func GetStation(name string) (models.Station, error) {
	if name == "" {
		return models.Station{}, fmt.Errorf("station name cannot be empty")
	}

	normalized := strings.ToLower(strings.TrimSpace(name))
	for i, station := range stations {
		if strings.EqualFold(station.Name, normalized) {
			return stations[i], nil
		}
	}

	return models.Station{}, fmt.Errorf("station %q not found", name)
}
