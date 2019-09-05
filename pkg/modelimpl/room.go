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
const eachDataSeparator = "-"

const host = "broker.hivemq.com:1883"

type room struct {
	roomMutex sync.RWMutex

	roomID uint64
	players map[uint64]*player
	capacity int
	circuit model.Circuit

	eventFeed *roomEventFeed

	mqttAdaptor *mqtt.Adaptor
	robot *gobot.Robot
}

func newRoom(roomID uint64, capacity int, circuitID uint64) *room {
	r := &room{
		players:  make(map[uint64]*player),
		roomID:   roomID,
		capacity: capacity,
		circuit: &circuit{
			circuitID: circuitID,
		},
		eventFeed: newRoomEventFeed(),

	}

	mqttAdaptor := mqtt.NewAdaptor(host, "pinger")
	work := func() {
		//data example: 1#11.3,31.6
		mqttAdaptor.On(fmt.Sprintf("gr-update-racer-position-room-%d", roomID), func(msg mqtt.Message) {
			log.Println("Receive message!")
			responseData := strings.Split(string(msg.Payload()), dataSeparator)
			positionData := strings.Split(responseData[1], positionSeparator)
			position := model.NewPosition(positionData)

			playerID, _ := strconv.ParseUint(responseData[0], 10, 64)
			p := r.players[playerID]
			p.SetPosition(position)
		})

		//TODO: publish players' position to players
		gobot.Every(10 * time.Millisecond, func() {
			payload := r.buildMessagePayload()
			if _, err := mqttAdaptor.PublishWithQOS(fmt.Sprintf("gr-blast-racer-position-room-%d", roomID), 1, []byte(payload)); err != nil {
				log.Println("Fail to publish message! Error:", err)
			}
		})
	}
	r.mqttAdaptor = mqttAdaptor
	robot := gobot.NewRobot(
		fmt.Sprintf("mqttBot-room-%d", roomID),
		[]gobot.Connection{r.mqttAdaptor},
		work,
	)
	r.robot = robot

	return r
}

func (r *room) buildMessagePayload() string {
	//Example: 1#11.3,31.6-2#5.4,3.5
	var payload string
	for playerID, player := range r.players {
		payload += fmt.Sprint(playerID)
		payload += dataSeparator
		payload += fmt.Sprint(player.Position().X)
		payload += positionSeparator
		payload += fmt.Sprint(player.Position().Y)
		payload += eachDataSeparator
	}
	payload = strings.Trim(payload, eachDataSeparator)
	return payload
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
	go r.robot.Start()
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