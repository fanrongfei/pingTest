package pkg

import (
	"bytes"
	"fmt"
	"net"
	"os"
)
func Server(serverAddress string) {
	addr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	for{
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	fmt.Println("receive message "+string(data[:n]))
	fmt.Println(n, remoteAddr)
	var bb bytes.Buffer
	bb.WriteString("word")
	conn.WriteToUDP(bb.Bytes(), remoteAddr)
}
