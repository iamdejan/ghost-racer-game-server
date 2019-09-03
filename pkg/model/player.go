package model

type Player interface {
	PlayerID() uint64

	Name() string
	LapsCompleted() int
	CheckpointsCompleted() float64
	LatestCheckpointTimestamp() int64

	Position() Position
	SetPosition(position Position)
}

type Position struct {
	X float64
	Y float64
}

type PlayerPayload struct {
	PlayerID uint64
	PlayerName string
}