package distance_test

import (
	"testing"

	haversine "github.com/athomecomar/athome/backend/address/ent/distance"
)

var haversineTests = []struct {
	p     haversine.Coord
	q     haversine.Coord
	outKm float64
}{
	{
		haversine.Coord{Lat: 22.55, Lon: 43.12},  // Rio de Janeiro, Brazil
		haversine.Coord{Lat: 13.45, Lon: 100.28}, // Bangkok, Thailand
		6094.544408786774,
	},
	{
		haversine.Coord{Lat: 20.10, Lon: 57.30}, // Port Louis, Mauritius
		haversine.Coord{Lat: 0.57, Lon: 100.21}, // Padang, Indonesia
		5145.525771394785,
	},
	{
		haversine.Coord{Lat: 51.45, Lon: 1.15},  // Oxford, United Kingdom
		haversine.Coord{Lat: 41.54, Lon: 12.27}, // Vatican, City Vatican City
		1389.1793118293067,
	},
	{
		haversine.Coord{Lat: 22.34, Lon: 17.05}, // Windhoek, Namibia
		haversine.Coord{Lat: 51.56, Lon: 4.29},  // Rotterdam, Netherlands
		3429.89310043882,
	},
	{
		haversine.Coord{Lat: 63.24, Lon: 56.59}, // Esperanza, Argentina
		haversine.Coord{Lat: 8.50, Lon: 13.14},  // Luanda, Angola
		6996.18595539861,
	},
	{
		haversine.Coord{Lat: 90.00, Lon: 0.00}, // North/South Poles
		haversine.Coord{Lat: 48.51, Lon: 2.21}, // Paris,  France
		4613.477506482742,
	},
	{
		haversine.Coord{Lat: 45.04, Lon: 7.42},  // Turin, Italy
		haversine.Coord{Lat: 3.09, Lon: 101.42}, // Kuala Lumpur, Malaysia
		10078.111954385415,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range haversineTests {
		km := haversine.Haversine(input.p, input.q)

		if input.outKm != km {
			t.Errorf("fail: want %v %v -> %v got %v",
				input.p,
				input.q,
				input.outKm,
				km,
			)
		}
	}
}

var manhattanTests = []struct {
	p     haversine.Coord
	q     haversine.Coord
	outKm float64
}{
	{
		haversine.Coord{Lat: 22.55, Lon: 43.12},  // Rio de Janeiro, Brazil
		haversine.Coord{Lat: 13.45, Lon: 100.28}, // Bangkok, Thailand
		6094.544408786774,
	},
	{
		haversine.Coord{Lat: 20.10, Lon: 57.30}, // Port Louis, Mauritius
		haversine.Coord{Lat: 0.57, Lon: 100.21}, // Padang, Indonesia
		5145.525771394785,
	},
	{
		haversine.Coord{Lat: 51.45, Lon: 1.15},  // Oxford, United Kingdom
		haversine.Coord{Lat: 41.54, Lon: 12.27}, // Vatican, City Vatican City
		1389.1793118293067,
	},
	{
		haversine.Coord{Lat: 22.34, Lon: 17.05}, // Windhoek, Namibia
		haversine.Coord{Lat: 51.56, Lon: 4.29},  // Rotterdam, Netherlands
		3429.89310043882,
	},
	{
		haversine.Coord{Lat: 63.24, Lon: 56.59}, // Esperanza, Argentina
		haversine.Coord{Lat: 8.50, Lon: 13.14},  // Luanda, Angola
		6996.18595539861,
	},
	{
		haversine.Coord{Lat: 90.00, Lon: 0.00}, // North/South Poles
		haversine.Coord{Lat: 48.51, Lon: 2.21}, // Paris,  France
		4613.477506482742,
	},
	{
		haversine.Coord{Lat: 45.04, Lon: 7.42},  // Turin, Italy
		haversine.Coord{Lat: 3.09, Lon: 101.42}, // Kuala Lumpur, Malaysia
		10078.111954385415,
	},
}

func TestManhattanDistance(t *testing.T) {
	for _, input := range manhattanTests {
		km := haversine.Manhattan(input.p, input.q)

		if input.outKm != km {
			t.Errorf("fail: want %v %v -> %v got %v",
				input.p,
				input.q,
				input.outKm,
				km,
			)
		}
	}
}
