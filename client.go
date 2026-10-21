package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Connected to server")
	reply, err := bufio.NewReader(ln).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Print("Server reply: ", reply)
}
