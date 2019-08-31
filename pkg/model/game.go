package model

type Game interface {
	CreateRoom(capacity int, circuitID uint64) (Room, int, string)
	QueryRoom(roomID uint64) Room
}