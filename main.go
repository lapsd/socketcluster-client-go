package main

import (
	"fmt"
	"os"

	"github.com/lapsd/socketcluster-client-go/scclient"
)

func onConnect(client scclient.Client) {
	fmt.Println("Connected to server")
}

func onDisconnect(client scclient.Client, err error) {
	fmt.Printf("onDisconnect: %s\n", err.Error())
	os.Exit(1)
}

func onConnectError(client scclient.Client, err error) {
	fmt.Printf("onConnectError: %s\n", err.Error())
	os.Exit(1)
}

func onSetAuthentication(client scclient.Client, token string) {
	fmt.Println("Auth token received :", token)
}

func onAuthentication(client scclient.Client, isAuthenticated bool) {
	fmt.Println("Client authenticated :", isAuthenticated)
	go start(client)
}

func main() {
	var workc = make(chan int)
	client := scclient.New("ws://localhost:7000/socket.io/?EIO=3&transport=websocket&code=1234")
	client.SetBasicListener(onConnect, onConnectError, onDisconnect)
	client.SetAuthenticationListener(onSetAuthentication, onAuthentication)
	client.On("connected", func(eventName string, data interface{}) {
		onConnected(eventName, data, client)
	})
	client.On("test", onTest)
	client.EnableLogging()
	go client.Connect()
	<-workc
}

func start(client scclient.Client) {

}

func onConnected(eventName string, data interface{}, client scclient.Client) {
	fmt.Println("Got data ", data, " for event ", eventName)
	// client.EmitAck("test", "This is a sample message", func(eventName string, err interface{}, data interface{}) {
	// 	if err == nil {
	// 		fmt.Println("Got ack for emit event with data ", data, " and error ", err)
	// 	}
	// 	fmt.Println("ERROR", err)
	// })
}

func onTest(eventName string, data interface{}) {
	fmt.Println("TESTING", eventName)
}
