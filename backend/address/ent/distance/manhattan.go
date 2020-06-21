package distance

func Manhattan(a, b Coord) float64 {
	lata, latb := Coord{Lat: a.Lat}, Coord{Lat: b.Lat}
	lona, lonb := Coord{Lon: a.Lon}, Coord{Lon: b.Lon}
	return Haversine(lata, latb) + Haversine(lona, lonb)
}
