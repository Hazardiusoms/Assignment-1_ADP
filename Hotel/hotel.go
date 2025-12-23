package Hotel

import _ "fmt"

type Room struct {
	Number        int
	Type          string
	PricePerNight float64
	IsOccupied    bool
}

type Hotel struct {
	Name  string
	Rooms map[int]Room
}
