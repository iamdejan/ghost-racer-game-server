package model

type roomEventWithPayload struct {
	EventType roomEventType

	Payload interface{}
}

func (re roomEventWithPayload) Type() roomEventType {
	return re.EventType
}

type roomEventWithoutPayload struct {
	EventType roomEventType
}

func (re roomEventWithoutPayload) Type() roomEventType {
	return re.EventType
}

func RoomEventInsertPlayer(payload *PlayerJoinRoomEventPayload) RoomEvent {
	return roomEventWithPayload{
		EventType: playerJoinEventType,
		Payload:   payload,
	}
}

func RoomEventRemovePlayer(payload *PlayerLeaveRoomEventPayload) RoomEvent {
	return roomEventWithPayload{
		EventType: playerLeaveEventType,
		Payload:   payload,
	}
}

func RoomEventRaceStarts() RoomEvent {
	return roomEventWithoutPayload{
		EventType: raceStartsEventType,
	}
}