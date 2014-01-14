package main

import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	command:=""
	for i:=1;i<len(os.Args);i++{
		command+=os.Args[i]+" "
	}
	command=command[:len(command)-1]+"\n"
	
	fmt.Println(command)
	
	server := "127.0.0.1:12343"

	fmt.Println("Dialing..")
	conn, err := net.Dial("tcp", server)
	checkError(err)
	fmt.Println("Connected")

	_, err = conn.Write([]byte(command))
	checkError(err)
	fmt.Println("Sent data")

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println("Result is")
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
