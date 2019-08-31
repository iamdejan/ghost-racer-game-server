package modelimpl

type player struct {
	playerID uint64
	name string
	lapsCompleted int
	checkpointsCompleted float64
	latestCheckpointTimestamp int64
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