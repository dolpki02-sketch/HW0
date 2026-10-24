package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if err := startClient(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startClient() error {
	ln, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return fmt.Errorf("Error connecting to server: %v", err)
	}
	defer ln.Close()
	fmt.Println("Connected to server")
	reply, err := bufio.NewReader(ln).ReadString('\n')
	if err != nil {
		return fmt.Errorf("Error reading from server: %v", err)
	}
	fmt.Print("Server reply: ", reply)
	return nil
}
