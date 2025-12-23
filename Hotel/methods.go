package Hotel

import "fmt"

// Constructor
func NewHotel(name string) *Hotel {
	return &Hotel{
		Name:  name,
		Rooms: make(map[int]Room),
	}
}

// Add room to hotel
func (h *Hotel) AddRoom(room Room) {
	h.Rooms[room.Number] = room
}

// Book a room
func (h *Hotel) BookRoom(roomNumber int) {
	room, exists := h.Rooms[roomNumber]
	if !exists {
		return
	}
	room.IsOccupied = true
	h.Rooms[roomNumber] = room
}

// Print all rooms
func (h *Hotel) PrintRooms() {
	for _, room := range h.Rooms {
		fmt.Printf(
			"Room %d | Type: %s | Price: %.2f | Occupied: %v\n",
			room.Number,
			room.Type,
			room.PricePerNight,
			room.IsOccupied,
		)
	}
}
