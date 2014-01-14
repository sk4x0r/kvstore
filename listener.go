package main
import(
	"fmt"
	"net"
	"strconv"
	"bufio"
	"strings"
)
var DEBUG int=1
func print(s string){
	if DEBUG==1{
		fmt.Println(s)
	}
}

const PORT = 12343
var	dict=make(map[string] string)

func main(){
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		panic("couldn't start listening on port"+strconv.Itoa(PORT)+". "+err.Error())
	}
	
	print("Started Listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		print("Got client")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	reader:=bufio.NewReader(conn)
		command, err :=reader.ReadString('\n')
		if err!=nil{
			fmt.Print(err.Error())
			return
		}
		print("Received command:")
		print(command)
		words:=strings.Fields(command)
		action:=words[0]
		
		print("action="+action)
		switch action{
			case "get":
				print("case get")
				key:=words[1]
				value, ok := dict[key]
				if ok{
					conn.Write([]byte(value+"\n"))
				}else{
					print("Key does not exist")
					conn.Write([]byte("KEY_DOES_NOT_EXIST"))
				}
				
			case "set":
				print("case set")
				key:=words[1]
				value:=words[2]
				dict[key]=value
				print("key="+key)
				print("value="+value)
				print("printing value from dict")
				print(dict[key])
				conn.Write([]byte("SUCCESS\n"))
			case "delete":
				print("case delete")
				key:=words[1]
				_, ok:=dict[key]
				if ok{
					delete(dict,key)
					conn.Write([]byte("SUCCESS\n"))
				}else{
					conn.Write([]byte("KEY_DOES_NOT_EXIST\n"))
				}
		}
}
