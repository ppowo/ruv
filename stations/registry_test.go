package stations

import "testing"

func TestGetStationSupportsCanonicalCodesAndLegacyAliases(t *testing.T) {
	tests := []struct {
		input    string
		wantCode string
	}{
		{input: "nood", wantCode: "nood"},
		{input: "drmm", wantCode: "drmm"},
		{input: "9128", wantCode: "9128"},
		{input: "lyll", wantCode: "nood"},
		{input: "cash", wantCode: "drmm"},
		{input: "lake", wantCode: "9128"},
		{input: "NOOD", wantCode: "nood"},
		{input: "  cash  ", wantCode: "drmm"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			station, err := GetStation(tt.input)
			if err != nil {
				t.Fatalf("GetStation(%q) returned error: %v", tt.input, err)
			}

			if station.Name != tt.wantCode {
				t.Fatalf("GetStation(%q) code = %q, want %q", tt.input, station.Name, tt.wantCode)
			}
		})
	}
}

func TestGetStationsReturnsCanonicalStationsOnly(t *testing.T) {
	got := GetStations()
	want := []string{"reso", "rese", "ntso", "ntst", "nood", "drmm", "9128", "alha"}

	if len(got) != len(want) {
		t.Fatalf("GetStations() len = %d, want %d", len(got), len(want))
	}

	for i, wantCode := range want {
		if got[i].Name != wantCode {
			t.Fatalf("GetStations()[%d] = %q, want %q", i, got[i].Name, wantCode)
		}
	}
}

func TestGetStationsReturnsCopy(t *testing.T) {
	copied := GetStations()
	copied[0].Name = "changed"

	fresh := GetStations()
	if fresh[0].Name != "reso" {
		t.Fatalf("GetStations() returned mutable backing data, first station = %q", fresh[0].Name)
	}
}
