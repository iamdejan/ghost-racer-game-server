package modelimpl

import (
	"github.com/iamdejan/ghost-racer-game-server/internal/utility"
	"github.com/iamdejan/ghost-racer-game-server/pkg/model"
	"net/http"
	"sync"
	"testing"
)

var game Game

func newGame() Game {
	game := Game{
		gameMutex:  sync.RWMutex{},
		rooms:      make(map[uint64]model.Room),
		nextRoomID: 0,
	}
	return game
}

func TestGame_CreateRoom(t *testing.T) {
	game = newGame()
	room, statusCode, _ := game.CreateRoom(2,1)

	utility.AssertEquals(statusCode, http.StatusOK, "Status Code is not 200!", t)

	utility.AssertNotNil(room, "room is null!", t)
	utility.AssertNotNil(room.Circuit(), "circuit is null!", t)
}

func TestGame_QueryRoom(t *testing.T) {
	TestGame_CreateRoom(t)

	room := game.QueryRoom(1)
	utility.AssertNotNil(room, "room is null!", t)
}