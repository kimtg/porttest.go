package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Host: ")
	scanner.Scan()
	host := scanner.Text()
	fmt.Print("Port: ")
	scanner.Scan()
	port := scanner.Text()

	for {
		fmt.Printf("%v %v %v ", time.Now(), host, port)
		conn, err := net.Dial("tcp", host+":"+port)
		if err != nil {
			fmt.Println("closed")
		} else {
			conn.Close()
			fmt.Println("open")
		}
		time.Sleep(10 * time.Second)
	}
}
