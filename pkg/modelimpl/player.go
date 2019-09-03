package modelimpl

import "github.com/iamdejan/ghost-racer-game-server/pkg/model"

type player struct {
	playerID uint64
	name string
	lapsCompleted int
	checkpointsCompleted float64
	latestCheckpointTimestamp int64

	position model.Position
}

func (p *player) PlayerID() uint64 {
	return p.playerID
}

func (p *player) Name() string {
	return p.name
}

func (p *player) LapsCompleted() int {
	return p.lapsCompleted
}

func (p *player) CheckpointsCompleted() float64 {
	return p.checkpointsCompleted
}

func (p *player) LatestCheckpointTimestamp() int64 {
	return p.latestCheckpointTimestamp
}

func (p *player) Position() model.Position {
	return p.position
}

func (p *player) SetPosition(position model.Position) {
	p.position = position
}