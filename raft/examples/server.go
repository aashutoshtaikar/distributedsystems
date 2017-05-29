package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing


func main() {

  fmt.Println("Launching server...")

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":8081")

  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  for {
    // will listen for message to process ending in newline (\n)
    message,_:= bufio.NewReader(conn).ReadString('\n')
    // output message received
    fmt.Print("Message Received:", string(message))
    // sample process for string received
     newmessage:= strings.ToUpper(message)
    //election detect 
      if strings.TrimRight(message, "\n")=="e" { 
    // send new string back to client
   fmt.Print("in loop",newmessage)
conn.Write([]byte(newmessage + "\n"))
}
  }
}
