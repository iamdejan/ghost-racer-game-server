package utility

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func isPointer(dataType string) bool {
	return dataType[0] == '*'
}

func ParseJSON(writer http.ResponseWriter, byteData []byte, value interface{}) bool {
	if isPointer(reflect.TypeOf(value).String()) != true {
		log.Println("interface is not pointer!")
		writer.WriteHeader(http.StatusInternalServerError)
		return false
	}

	if err := json.Unmarshal(byteData, value); err != nil {
		log.Println("Fail to unmarshal JSON! Error:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return false
	}

	return true
}

func WriteJSON(writer http.ResponseWriter, value interface{}) bool {
	byteData, _ := json.Marshal(value)
	writer.Header().Set("Content-Type", "application/json")

	if _, err := writer.Write(byteData); err != nil {
		log.Println("Fail to write to response! Error:", err)
		writer.WriteHeader(500)
		return false
	}

	return true
}

func ReadBody(request *http.Request, writer http.ResponseWriter) ([]byte, bool) {
	byteData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("Fail to read request.Body! Error:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return nil, true
	}
	return byteData, false
}
