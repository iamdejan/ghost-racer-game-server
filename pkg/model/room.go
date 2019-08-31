package model

type Room interface {
	//HTTP Protocol
	ID() uint64
	Capacity() int
	Circuit() Circuit
	InsertPlayer(payload PlayerPayload) bool
	RemovePlayer(playerID uint64) bool
	QueryPlayer(playerID uint64) Player
	Events() RoomEventFeed
}