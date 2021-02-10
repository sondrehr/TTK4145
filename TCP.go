package main

import(
	"fmt"
	//"time"
	"net"
	"os"
	//"io/ioutil"
	"bufio"
)

//Connect to: 10.100.23.131:20010  

func client(){
	conn, _:= net.Dial("tcp","10.100.23.147:33546")

	for{
		// what to send?
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to server
		fmt.Fprintf(conn, text + "\000")
		// wait for reply
		message, _:= bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}




func main(){
	go client()

	for{

	}
}
