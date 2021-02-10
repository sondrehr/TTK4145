package main

import(
	"fmt"
	"time"
	"net"
)



func recieve() {
	buffer := make([]byte, 1024)
	addr,_:= net.ResolveUDPAddr("udp4", ":30000")
	conn,_:= net.ListenUDP("udp4", addr)
	for{
		bytes, _, _ := conn.ReadFromUDP(buffer)
		str := string(buffer[:bytes])
		fmt.Println(str)
	}
}

func send(){
	connection, _:= net.Dial("udp4","10.100.23.255:20010")
	for{
		connection.Write([]byte("Hei server:) hvordan går det?"))
		time.Sleep(1*time.Second)
	}
}

func recieveM() {
	buffer := make([]byte, 1024)
	addr,_:= net.ResolveUDPAddr("udp4", ":20010")
	conn,_:= net.ListenUDP("udp4", addr)
	for{
		bytes, _, _ := conn.ReadFromUDP(buffer)
		str := string(buffer[:bytes])
		fmt.Println(str)
	}
}


func main(){
	//go recieve()
	go send()
	go recieveM()

	for{}
}
