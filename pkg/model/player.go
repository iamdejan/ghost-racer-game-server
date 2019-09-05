package model

import (
	"log"
	"strconv"
)

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

func NewPosition(data []string) (p Position) {
	x, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		log.Fatal("Fail to parse! Error:", err)
	}
	y, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		log.Fatal("Fail to parse! Error:", err)
	}
	p = Position{
		X: x,
		Y: y,
	}
	return p
}

type PlayerPayload struct {
	PlayerID uint64
	PlayerName string
}