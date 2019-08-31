package httphandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iamdejan/ghost-racer-game-server/internal/utility"
	"github.com/iamdejan/ghost-racer-game-server/pkg/modelimpl"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var handler Handler

func TestHandler_CreateHandlerAndRouter(t *testing.T) {
	handler = Handler{
		RequestMaxDuration: 0,
		Game: modelimpl.NewGame(),
	}

	router := handler.newRouter()

	utility.AssertEquals(fmt.Sprintf("%T", router), "*mux.Router", "", t)
}

func TestHandler_CreateRoom_ValidData(t *testing.T) {
	TestHandler_CreateHandlerAndRouter(t)

	data := createRoomRequest{
		Capacity: 4,
		CircuitID: 1,
	}

	statusCode, responseData := createRoomTesting(data, t)
	var roomData createRoomResponse
	if err := json.Unmarshal(responseData, &roomData); err != nil {
		t.Fatal("Fail to parse JSON!")
	}

	utility.AssertEquals(statusCode, http.StatusOK, "status code is not OK!", t)
	utility.AssertTrue(roomData.RoomID > 0, "roomID must be > 0!", t)
	utility.AssertTrue(roomData.Capacity >= 2, "room capacity must be at least 2!", t)
}

func TestHandler_CreateRoomFail(t *testing.T) {
	TestHandler_CreateHandlerAndRouter(t)

	data := createRoomRequest{
		Capacity: 1,
		CircuitID: 1,
	}

	statusCode, responseData := createRoomTesting(data, t)
	var message map[string]string
	if err := json.Unmarshal(responseData, &message); err != nil {
		t.Fatal("Fail to parse JSON! Error:", err)
	}
	utility.AssertNotEquals(statusCode, http.StatusOK, "Status Code shouldn't be OK", t)
	utility.AssertEquals(statusCode, http.StatusForbidden, "Status Code should be 403", t)
}

func TestHandler_JoinRoom(t *testing.T) {
	TestHandler_CreateRoom_ValidData(t)

	data := joinRoomRequest{
		PlayerID:   1,
		PlayerName: "Dejan",
	}

	statusCode, responseData := joinRoomTesting(1, data, t)
	utility.AssertEquals(statusCode, http.StatusOK, "Status code != 200!", t)

	var roomData joinRoomResponse
	if err := json.Unmarshal(responseData, &roomData); err != nil {
		t.Fatal("Fail to parse JSON! Error:", err)
	}

	utility.AssertNotNil(roomData, "roomData is nil!", t)
}

func createRoomTesting(data interface{}, t *testing.T) (statusCode int, responseData []byte) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatal("Fail to marshal JSON! Error:", err)
	}
	request, err := http.NewRequest(http.MethodPost, "/room/new", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal("Fail to create request! Error:", err)
	}
	response := httptest.NewRecorder()
	utility.AssertNotNil(handler.Game, "handler.Game is null!", t)
	handler.createRoom(response, request)

	responseData = readBodyData(response, t)
	return response.Code, responseData
}

func joinRoomTesting(roomID uint64, data interface{}, t *testing.T) (statusCode int, responseData []byte) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatal("Fail to marshal JSON! Error:", err)
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/room/%d/player/join", roomID), bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal("Fail to create request! Error:", err)
	}
	request = mux.SetURLVars(request, map[string]string{"roomID": fmt.Sprintf("%d", roomID)})

	response := httptest.NewRecorder()
	utility.AssertNotNil(handler.Game, "handler.Game is null!", t)
	handler.joinRoom(response, request)

	responseData = readBodyData(response, t)
	return response.Code, responseData
}

func readBodyData(response *httptest.ResponseRecorder, t *testing.T) []byte {
	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Fail to read response data! Error:", err)
	}
	return bodyData
}