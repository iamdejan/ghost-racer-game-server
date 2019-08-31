package model

type roomEventWithPayload struct {
	eventType roomEventType

	Payload interface{}
}

func (re roomEventWithPayload) Type() roomEventType {
	return re.eventType
}

type roomEventWithoutPayload struct {
	eventType roomEventType
}

func (re roomEventWithoutPayload) Type() roomEventType {
	return re.eventType
}

func RoomEventInsertPlayer(payload *PlayerJoinRoomEventPayload) RoomEvent {
	return roomEventWithPayload{
		eventType: playerJoinEventType,
		Payload:   payload,
	}
}

func RoomEventRemovePlayer(payload *PlayerLeaveRoomEventPayload) RoomEvent {
	return roomEventWithPayload{
		eventType: playerLeaveEventType,
		Payload:   payload,
	}
}

func RoomEventRaceStarts() RoomEvent {
	return roomEventWithoutPayload{
		eventType: raceStartsEventType,
	}
}