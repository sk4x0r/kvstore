package kvstore

import (
	"encoding/json"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"strconv"
)

var dict = make(map[string]string)

type Command struct {
	Operation string
	Key       string
	Value     string
}

func CreateServer(port int) {
	responder, err := zmq.NewSocket(zmq.REP)
	if err != nil {
		panic("couldn't start listening on port" + strconv.Itoa(port) + ". " + err.Error())
	}
	defer responder.Close()
	responder.Bind("tcp://*:" + strconv.Itoa(port))
	for {
		request, _ := responder.Recv(0)
		var command Command
		err := json.Unmarshal([]byte(request), &command)
		if err != nil {
			fmt.Println("Could not understand the request")
			responder.Send(RESPONSE_INVALID_REQUEST, 0)
			continue
		}
		reply := handleRequest(command)
		responder.Send(reply, 0)
	}
}

func handleRequest(command Command) string {
	operation := command.Operation
	switch operation {
	case OPERATION_GET:
		key := command.Key
		value, ok := dict[key]
		if ok {
			return value
		} else {
			return RESPONSE_KEY_NOT_FOUND
		}

	case OPERATION_SET:
		key := command.Key
		value := command.Value
		dict[key] = value
		return "OK"

	case OPERATION_DELETE:
		key := command.Key
		_, ok := dict[key]
		if ok {
			delete(dict, key)
			return "OK"
		} else {
			return RESPONSE_KEY_NOT_FOUND
		}

	default:
		return RESPONSE_INVALID_REQUEST
	}
}
