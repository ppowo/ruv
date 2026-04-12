package stations

import "testing"

func TestGetStation(t *testing.T) {
	tests := []struct {
		input    string
		wantCode string
	}{
		{input: "nood", wantCode: "nood"},
		{input: "drmm", wantCode: "drmm"},
		{input: "9128", wantCode: "9128"},
		{input: "NOOD", wantCode: "nood"},
		{input: "  nood  ", wantCode: "nood"},
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

func TestGetStationReturnsCopy(t *testing.T) {
	station, err := GetStation("reso")
	if err != nil {
		t.Fatalf("GetStation(\"reso\") returned error: %v", err)
	}

	station.Name = "changed"

	fresh, _ := GetStation("reso")
	if fresh.Name != "reso" {
		t.Fatalf("GetStation returned mutable backing data, first station = %q", fresh.Name)
	}
}
