package kvstore

import (
	//"flag"
	"encoding/json"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"io/ioutil"
	"strconv"
	"testing"
	"time"
)

type Row struct {
	Key   string
	Value string
}

type Data struct {
	Rows []Row
}

func createSetRequest(key string, value string) string {
	message := "{"
	message += "\"Operation\":\"" + OPERATION_SET + "\"" + ","
	message += "\"Key\":\"" + key + "\","
	message += "\"Value\":\"" + value + "\""
	message += "}"
	return message
}

func createGetRequest(key string) string {
	message := "{"
	message += "\"Operation\":\"" + OPERATION_GET + "\"" + ","
	message += "\"Key\":\"" + key + "\""
	message += "}"
	return message
}

func createDeleteRequest(key string) string {
	message := "{"
	message += "\"Operation\":\"" + OPERATION_DELETE + "\"" + ","
	message += "\"Key\":\"" + key + "\""
	message += "}"
	return message
}

func Send(command string, serverAddress string, serverPort int) string {
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect("tcp://" + serverAddress + ":" + strconv.Itoa(serverPort))
	requester.Send(command, 0)
	reply, err := requester.Recv(0)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	//fmt.Println(reply)
	return reply
}

func sendRequests(serverAddress string, serverPort int, success chan bool, failure chan bool) {
	dataFile := "data.json"
	content, err := ioutil.ReadFile(dataFile)
	if err != nil {
		fmt.Println("Error parsing the config file")
		panic(err)
	}
	var data Data
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("Error parsing the config file")
		panic(err)
	}
	rows := data.Rows
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		request := createSetRequest(row.Key, row.Value)
		response := Send(request, serverAddress, serverPort)
		if response != "OK" {
			fmt.Println("Failure")
			failure <- true
			return
		}
	}
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		request := createGetRequest(row.Key)
		response := Send(request, serverAddress, serverPort)
		if response != row.Value {
			fmt.Println("Expected:", row.Value, "Received:", response)
			failure <- true
			return
		}
	}
	success <- true
}

func TestSetAndGet(t *testing.T) {
	serverAddress := "127.0.0.1"
	serverPort := 12343
	
	go CreateServer(serverPort)
	
	success := make(chan bool, 1)
	failure := make(chan bool, 1)
	go sendRequests(serverAddress, serverPort, success, failure)
	
	select {
	case <-success:
		fmt.Println("Test passed successfully")
		break
	case <-failure:
		t.Errorf("Test failed")
		break
	case <-time.After(50 * time.Second):
		t.Errorf("Could not send messages")
		break
	}
}
