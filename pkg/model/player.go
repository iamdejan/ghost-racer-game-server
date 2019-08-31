package model

type Player interface {
	PlayerID() uint64

	Name() string
	LapsCompleted() int
	CheckpointsCompleted() float64
	LatestCheckpointTimestamp() int64
}

type PlayerPayload struct {
	PlayerID uint64
	PlayerName string
}