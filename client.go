package main

import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
)

var DEBUG int=1
func print(s string){
	if DEBUG==1{
		fmt.Println(s)
	}
}

func main() {
	command:=""
	for i:=1;i<len(os.Args);i++{
		command+=os.Args[i]+" "
	}
	command=command[:len(command)-1]+"\n"
	
	print(command)
	
	server := "127.0.0.1:12343"

	print("Dialing connection")
	conn, err := net.Dial("tcp", server)
	checkError(err)
	print("Connected")

	_, err = conn.Write([]byte(command))
	checkError(err)
	print("Sent data")

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	print("Result is")
	print(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
