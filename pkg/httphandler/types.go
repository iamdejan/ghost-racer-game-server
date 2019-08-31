package httphandler

type createRoomRequest struct {
	Capacity int
	CircuitID uint64
}

type createRoomResponse struct {
	RoomID uint64
	Capacity int
	CircuitID uint64
}

type joinRoomRequest struct {
	PlayerID uint64
	PlayerName string
}

type joinRoomResponse struct {
	RoomID uint64
	Capacity int
	CircuitID uint64
}