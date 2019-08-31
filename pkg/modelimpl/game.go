package modelimpl

import (
	"github.com/iamdejan/ghost-racer-game-server/pkg/model"
	"net/http"
	"sync"
)

type Game struct {
	gameMutex sync.RWMutex

	rooms map[uint64]model.Room
	nextRoomID uint64
}

func NewGame() model.Game {
	game := Game{
		rooms:      make(map[uint64]model.Room),
		nextRoomID: 0,
	}

	return &game
}

func (game *Game) CreateRoom(capacity int, circuitID uint64) (model.Room, int, string) {
	if capacity < 2 {
		return nil, http.StatusForbidden, "Capacity must be 2 or greater!"
	}

	game.gameMutex.Lock()
	defer game.gameMutex.Unlock()

	game.nextRoomID += 1
	roomID := game.nextRoomID

	var room = newRoom(roomID, capacity, circuitID)

	game.rooms[roomID] = room

	return room, 200, ""
}

func newRoom(roomID uint64, capacity int, circuitID uint64) *room {
	return &room{
		players:  make(map[uint64]model.Player),
		roomID:   roomID,
		capacity: capacity,
		circuit: &circuit{
			circuitID: circuitID,
		},
	}
}

func (game *Game) QueryRoom(roomID uint64) model.Room {
	return game.rooms[roomID]
}