package main

import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
	"strconv"
)

var DEBUG int=1 //Setting this to value 0 will disable console outputs
func print(s string){
	if DEBUG==1{
		fmt.Println(s)
	}
}

const PORT = 12343

func main() {
	command:=""
	for i:=1;i<len(os.Args);i++{
		command+=os.Args[i]+" "
	}
	command=command[:len(command)-1]+"\n"
	
	print("Command:"+command)
	
	server := "127.0.0.1:"+strconv.Itoa(PORT)

	print("Dialing connection")
	conn, err := net.Dial("tcp", server)
	checkError(err)
	print("Connected to server")

	_, err = conn.Write([]byte(command))
	checkError(err)
	print("Sent data")

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	print("Received response:"+string(result))
	print("")
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
