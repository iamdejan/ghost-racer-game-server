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

	events *roomEventFeed
}

func newRoom(roomID uint64, capacity int, circuitID uint64) *room {
	return &room{
		players:  make(map[uint64]model.Player),
		roomID:   roomID,
		capacity: capacity,
		circuit: &circuit{
			circuitID: circuitID,
		},
		events: newRoomEventFeed(),
	}
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

func (r *room) isFull() bool {
	return r.capacity == len(r.players)
}

func (r *room) doesGameStart() bool {
	return r.isFull()
}

func (r *room) startRace() {
	r.events.put(model.RoomEventRaceStarts())
}

func (r *room) InsertPlayer(payload model.PlayerPayload) bool {
	r.roomMutex.Lock()
	defer r.roomMutex.Unlock()

	if r.doesGameStart() {
		return false
	}

	player := buildPlayerWithPayload(payload)
	
	r.players[player.PlayerID()] = player

	r.events.put(model.RoomEventInsertPlayer(&model.PlayerJoinRoomEventPayload{
		PlayerID:   player.playerID,
		PlayerName: player.name,
	}))

	if r.isFull() {
		r.startRace()
	}

	return true
}

func (r *room) RemovePlayer(playerID uint64) bool {
	r.roomMutex.Lock()
	defer r.roomMutex.Unlock()

	if r.doesGameStart() {
		return false
	}

	delete(r.players, playerID)

	r.events.put(model.RoomEventRemovePlayer(&model.PlayerLeaveRoomEventPayload{
		PlayerID: playerID,
	}))

	return true
}

func (r *room) QueryPlayer(playerID uint64) model.Player {
	return r.players[playerID]
}

func (r *room) EventFeed() model.RoomEventFeed {
	return r.events
}