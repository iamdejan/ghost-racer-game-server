package httphandler

import (
	"github.com/iamdejan/ghost-racer-game-server/internal/utility"
	"github.com/iamdejan/ghost-racer-game-server/pkg/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const errorMessageKey = "ErrorMessage"

type Handler struct {
	RequestMaxDuration time.Duration
	Game model.Game
}

func (handler *Handler) newRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func (handler *Handler) Routes() http.Handler {
	router := handler.newRouter()

	router.HandleFunc("/room/new", handler.createRoom).Methods(http.MethodPost)
	router.HandleFunc("/room/{roomID}/player/join", handler.joinRoom).Methods(http.MethodPost)
	router.HandleFunc("/room/{roomID}/events/{lastID}", handler.getEvents).Methods(http.MethodGet)

	return router
}

func (handler *Handler) createRoom(writer http.ResponseWriter, request *http.Request) {
	byteData, done := utility.ReadBody(request, writer)
	if done {
		return
	}

	var payload createRoomRequest
	if utility.ParseJSON(writer, byteData, &payload) != true {
		return
	}

	room, statusCode, errorMessage := handler.Game.CreateRoom(payload.Capacity, payload.CircuitID)
	if room == nil {
		writer.WriteHeader(statusCode)
		writer.Write([]byte("{\"" + errorMessage + "\": \"" + errorMessage + "\"}"))
		return
	}

	utility.WriteJSON(writer, createRoomResponse{
		RoomID:    room.ID(),
		Capacity:  room.Capacity(),
		CircuitID: room.Circuit().ID(),
	})
}

func (handler *Handler) joinRoom(writer http.ResponseWriter, request *http.Request) {
	byteData, done := utility.ReadBody(request, writer)
	if done {
		return
	}

	var roomPayload joinRoomRequest
	if utility.ParseJSON(writer, byteData, &roomPayload) != true {
		return
	}

	roomID, err := strconv.ParseUint(mux.Vars(request)["roomID"], 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("{\"" + errorMessageKey + "\": \"" + "Can't parse roomID!" + "\"}"))
		return
	}

	room := handler.Game.QueryRoom(roomID)
	if room == nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("{\"" + errorMessageKey + "\": \"" + "Room isn't found!" + "\"}"))
		return
	}

	var playerPayload = model.PlayerPayload(roomPayload)
	if room.InsertPlayer(playerPayload) != true {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("{\"" + errorMessageKey + "\": \"" + "Fail to insert player!" + "\"}"))
		return
	}

	utility.WriteJSON(writer, joinRoomResponse{
		RoomID:    roomID,
		Capacity:  room.Capacity(),
		CircuitID: room.Circuit().ID(),
	})
}

func (handler *Handler) getEvents(writer http.ResponseWriter, request *http.Request) {
	lastID, err := strconv.Atoi(mux.Vars(request)["lastID"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("{\"" + errorMessageKey + "\": \"" + "Fail to get lastID from URL!" + "\"}"))
		return
	}

	roomID, err := strconv.ParseUint(mux.Vars(request)["roomID"], 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("{\"" + errorMessageKey + "\": \"" + "Fail to get roomID from URL!" + "\"}"))
		return
	}

	room := handler.Game.QueryRoom(roomID)
	if room == nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("{\"" + errorMessageKey + "\": \"" + "Room isn't found!" + "\"}"))
		return
	}

	events, newLastID, waitChannel := room.EventFeed().List(lastID)
	if len(events) == 0 {
		select {
			case <- waitChannel:
				events, newLastID, _ = room.EventFeed().List(lastID)
			case <- time.After(handler.RequestMaxDuration):
		}
	}

	utility.WriteJSON(writer, roomEventsResponse{
		Events: events,
		LastID: newLastID,
	})
}