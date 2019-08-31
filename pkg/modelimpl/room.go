package modelimpl

import (
	"github.com/iamdejan/ghost-racer-game-server/pkg/model"
	"sync"
	"time"
)

type room struct {
	roomMutex sync.RWMutex

	roomID uint64
	players map[uint64]model.Player
	capacity int
	circuit model.Circuit
}

func (r *room) ID() uint64 {
	return r.roomID
}

func (r *room) Capacity() int {
	return r.capacity
}

func (r *room) Circuit() model.Circuit {
	return r.circuit
}

func buildPlayerWithPayload(payload model.PlayerPayload) *player {
	return &player{
		playerID:                  payload.PlayerID,
		name:                      payload.PlayerName,
		lapsCompleted:             0,
		checkpointsCompleted:      0.0,
		latestCheckpointTimestamp: time.Now().UnixNano(),
	}
}

func (r *room) InsertPlayer(payload model.PlayerPayload) bool {
	r.roomMutex.Lock()
	defer r.roomMutex.Unlock()

	player := buildPlayerWithPayload(payload)
	
	r.players[player.PlayerID()] = player

	return true
}

func (r *room) RemovePlayer(playerID uint64) bool {
	r.roomMutex.Lock()
	defer r.roomMutex.Unlock()

	delete(r.players, playerID)

	return true
}

func (r *room) QueryPlayer(playerID uint64) model.Player {
	return r.players[playerID]
}

func (r *room) Events() model.RoomEventFeed {
	return nil
}