package modelimpl

import (
	"github.com/iamdejan/ghost-racer-game-server/internal/utility"
	"github.com/iamdejan/ghost-racer-game-server/pkg/model"
	"testing"
)

func buildNewRoom(c circuit) room {
	return room{
		roomID:0,
		players:make(map[uint64]model.Player),
		capacity:0,
		circuit: &c,
	}
}

func buildNewCircuit(circuitID uint64) circuit {
	return circuit{circuitID:circuitID}
}

func initRoomAndCircuit() room {
	var c circuit = buildNewCircuit(1)
	var r room = buildNewRoom(c)
	return r
}

func buildNewPlayer(playerID uint64, name string) model.PlayerPayload {
	return model.PlayerPayload{
		PlayerID: playerID,
		PlayerName: name,
	}
}

func initiateTestData() (room, model.PlayerPayload) {
	var r room = initRoomAndCircuit()
	var p = buildNewPlayer(1, "Dejan")

	return r, p
}

func TestRoom_InsertPlayer(t *testing.T) {
	r, p := initiateTestData()

	//Check if room member's size is added by 1 or not
	var oldSize int = len(r.players)

	if r.InsertPlayer(p) != true {
		t.Fatal("Insert player failed!")
	}

	var newSize int = len(r.players)

	utility.AssertNotEquals(oldSize, newSize, "newSize & oldSize must not be equal!", t)
	utility.AssertTrue(newSize - oldSize == 1, "newSize - oldSize must be 1!", t)
}

func TestRoom_RemovePlayer(t *testing.T) {
	r, p := initiateTestData()
	r.InsertPlayer(p)
	var oldSize int = len(r.players)

	if r.RemovePlayer(1) != true {
		t.Fatal("Fail to remove player!")
	}

	var newSize int = len(r.players)

	utility.AssertNotEquals(oldSize, newSize, "newSize & oldSize must not be equal!", t)
	utility.AssertTrue(newSize - oldSize == -1, "newSize - oldSize must be -1!", t)
}

func TestRoom_QueryPlayer(t *testing.T) {
	r, p := initiateTestData()
	r.InsertPlayer(p)

	queryPlayerResult := r.QueryPlayer(p.PlayerID)
	utility.AssertNotNil(queryPlayerResult, "queryPlayerResult should not be nil!", t)
	if queryPlayerResult.PlayerID() != p.PlayerID {
		t.Fatal("User query isn't correct!")
	}
}

func TestRoom_QueryPlayer_NotFound(t *testing.T) {
	r, p := initiateTestData()
	r.InsertPlayer(p)
	queryPlayerResult := r.QueryPlayer(2)
	utility.AssertTrue(queryPlayerResult == nil, "queryPlayerResult should be nil!", t)
}