package model

type RoomEventFeed interface {
	List(lastID int) ([]RoomEvent, int, <-chan struct{})
}

type roomEventType string
type RoomEvent interface {
	Type() roomEventType
}

const (
	playerJoinEventType roomEventType = "PLAYER_JOIN"
	playerLeaveEventType roomEventType = "PLAYER_LEAVE"
	raceStartsEventType roomEventType = "RACE_STARTS"
	raceEndsEventType roomEventType = "RACE_ENDS"
)

type PlayerJoinRoomEventPayload struct {
	PlayerID uint64
	Name string
}

type PlayerLeaveRoomEventPayload struct {
	PlayerID uint64
}

type RaceEndsEventPayload struct {
	WinnerID uint64
	WinnerName string
	RaceTime int64 //recorded at finish
}