package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/rgamba/evtwebsocket"
)

func PrintMessage(message string) {
	fmt.Println(message)
}

func IsEqual(s string, b []byte) bool {
	if len(s) != len(b) {
		return false
	}
	for i, x := range b {
		if x != s[i] {
			return false
		}
	}
	return true
}

func CreateMessageFromString(message string) evtwebsocket.Msg {
	return evtwebsocket.Msg{
		Body: []byte(message),
	}
}

func CreateMessageFromByte(message []byte) evtwebsocket.Msg {
	return evtwebsocket.Msg{
		Body: message,
	}
}

func SerializeData(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}

func SerializeDataIntoString(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

func DeserializeData(data []byte) (jsonObject interface{}) {
	json.Unmarshal(data, &jsonObject)
	return
}

func DeserializeDataFromString(data string) (jsonObject interface{}) {
	json.Unmarshal([]byte(data), &jsonObject)
	return
}

func GetEventData(str string) (eventname, data string) {
	re := regexp.MustCompile(`[0-9]+`)
	re2 := regexp.MustCompile(`,`)
	result := re.Split(str, 2)
	result = re2.Split(result[1], 2)
	eventname = strings.Replace(result[0], "[", "", -1)
	eventname = strings.Replace(eventname, "\"", "", -1)
	eventname = strings.TrimSpace(eventname)
	eventname = strings.TrimSpace(eventname)
	data = strings.Replace(result[1], "]", "", -1)
	fmt.Println("EVENT NAME: " + eventname)
	return
}
