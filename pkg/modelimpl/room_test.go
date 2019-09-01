package modelimpl

import (
	"github.com/iamdejan/ghost-racer-game-server/internal/utility"
	"github.com/iamdejan/ghost-racer-game-server/pkg/model"
	"testing"
)

func buildNewRoom(c circuit) room {
	return *newRoom(1, 2, 1)
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

func TestRoom_EventsWhenInsertPlayer(t *testing.T) {
	r, p := initiateTestData()
	r.InsertPlayer(p)

	eventFeed := r.EventFeed()
	roomEvents, newLastID, _ := eventFeed.List(-1)
	utility.AssertNotNil(roomEvents, "roomEvents should not be nil!", t)
	utility.AssertTrue(len(roomEvents) == 1, "roomEvents length should be 1!", t)
	utility.AssertNotEquals(newLastID, -1, "newLastID must be > -1!", t)
}

func TestRoom_EventsWhenRemovePlayer(t *testing.T) {
	r, p := initiateTestData()
	r.InsertPlayer(p)
	r.RemovePlayer(p.PlayerID)

	eventFeed := r.EventFeed()
	roomEvents, newLastID, _ := eventFeed.List(-1)
	utility.AssertNotNil(roomEvents, "roomEvents should not be nil!", t)
	utility.AssertTrue(len(roomEvents) == 2, "roomEvents length should be 2!", t)
	utility.AssertEquals(newLastID, 1, "newLastID must be 1!", t)
}

func TestRoom_StartRaceWhenFull(t *testing.T) {
	r, p1 := initiateTestData()
	r.InsertPlayer(p1)
	p2 := buildNewPlayer(2, "Daniel")
	r.InsertPlayer(p2)

	assertRoomIsFull(r, t)
}

func assertRoomIsFull(r room, t *testing.T) {
	eventFeed := r.EventFeed()
	roomEvents, newLastID, _ := eventFeed.List(-1)
	utility.AssertNotNil(roomEvents, "roomEvents should not be nil!", t)
	utility.AssertTrue(len(roomEvents) == 3, "roomEvents length should be 3!", t)
	utility.AssertEquals(newLastID, 2, "newLastID must be 2!", t)
	utility.AssertEquals(roomEvents[2].Type(), model.RoomEventRaceStarts().Type(), "Race should start!", t)
}

func TestRoom_JoinRoomWhenFull(t *testing.T) {
	r, p1 := initiateTestData()
	r.InsertPlayer(p1)
	p2 := buildNewPlayer(2, "Daniel")
	r.InsertPlayer(p2)
	p3 := buildNewPlayer(3, "Johan")
	r.InsertPlayer(p3)

	assertRoomIsFull(r, t)
}