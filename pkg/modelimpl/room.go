package modelimpl

import (
	"fmt"
	"github.com/iamdejan/ghost-racer-game-server/pkg/model"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

const dataSeparator = "#"
const positionSeparator = ","

type room struct {
	roomMutex sync.RWMutex

	roomID uint64
	players map[uint64]model.Player
	capacity int
	circuit model.Circuit

	eventFeed *roomEventFeed

	robot *gobot.Robot
}


func newPosition(data []string) (p model.Position) {
	x, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		log.Fatal("Fail to parse! Error:", err)
	}
	y, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		log.Fatal("Fail to parse! Error:", err)
	}
	p = model.Position{
		X: x,
		Y: y,
	}
	return p
}


//data example: 1#11.3,31.6
func newRoom(roomID uint64, capacity int, circuitID uint64) *room {
	r := &room{
		players:  make(map[uint64]model.Player),
		roomID:   roomID,
		capacity: capacity,
		circuit: &circuit{
			circuitID: circuitID,
		},
		eventFeed: newRoomEventFeed(),

	}

	mqttAdaptor := mqtt.NewAdaptor("test.mosquitto.org:1883", "pinger")
	work := func() {
		mqttAdaptor.On(fmt.Sprintf("gr-update-racer-position-at-room-%d", roomID), func(msg mqtt.Message) {
			responseData := strings.Split(string(msg.Payload()), dataSeparator)
			positionData := strings.Split(responseData[1], positionSeparator)
			position := newPosition(positionData)

			playerID, _ := strconv.ParseUint(responseData[0], 10, 64)
			p := r.QueryPlayer(playerID)
			p.SetPosition(position)
		})

		//TODO: publish players' position to players
		gobot.Every(10 * time.Millisecond, func() {
			//TODO: decide format
		})
	}

	robot := gobot.NewRobot(fmt.Sprintf("mqttBot-room-%d", roomID), []gobot.Connection{mqttAdaptor}, work)
	r.robot = robot

	return r
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
	r.eventFeed.put(model.RoomEventRaceStarts())
}

func (r *room) InsertPlayer(payload model.PlayerPayload) bool {
	r.roomMutex.Lock()
	defer r.roomMutex.Unlock()

	if r.doesGameStart() {
		return false
	}

	player := buildPlayerWithPayload(payload)
	
	r.players[player.PlayerID()] = player

	r.eventFeed.put(model.RoomEventInsertPlayer(&model.PlayerJoinRoomEventPayload{
		PlayerID:   player.PlayerID(),
		PlayerName: player.Name(),
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

	r.eventFeed.put(model.RoomEventRemovePlayer(&model.PlayerLeaveRoomEventPayload{
		PlayerID: playerID,
	}))

	return true
}

func (r *room) QueryPlayer(playerID uint64) model.Player {
	return r.players[playerID]
}

func (r *room) EventFeed() model.RoomEventFeed {
	return r.eventFeed
}