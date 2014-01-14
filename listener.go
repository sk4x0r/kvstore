package main
import(
	"fmt"
	"net"
	"strconv"
	"bufio"
	"strings"
)
var DEBUG int=1 //Setting this to value 0 will disable console outputs
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
		print("Received command:"+command)
		words:=strings.Fields(command)
		action:=words[0]
		
		switch action{
			case "get":
				key:=words[1]
				value, ok := dict[key]
				if ok{
					conn.Write([]byte(value))
				}else{
					print("Key does not exist")
					conn.Write([]byte("KEY_DOES_NOT_EXIST"))
				}
				
			case "set":
				key:=words[1]
				value:=""
				for i:=2;i<len(words);i++{
					value+=words[i]+" "
				}
				value=value[:len(value)-1]
				dict[key]=value
				conn.Write([]byte("SUCCESS"))
			case "delete":
				key:=words[1]
				_, ok:=dict[key]
				if ok{
					delete(dict,key)
					conn.Write([]byte("SUCCESS"))
				}else{
					conn.Write([]byte("KEY_DOES_NOT_EXIST"))
				}
		}
}
