package main

import (

	"log"
	"net"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	// Listen on TCP port 2000 on all interfaces.
fmt.Println("Launching server...")

l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for electiontimer:=0;electiontimer<10;electiontimer++{
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
go func(c net.Conn) {
			
 message,_:= bufio.NewReader(conn).ReadString('\n')

fmt.Print("Message Received:", string(message))
if strings.TrimRight(message, "\n")=="e" { 
    // send new string back to client
   fmt.Print("in loop",message)
conn.Write([]byte(message + "\n"))
}

			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
