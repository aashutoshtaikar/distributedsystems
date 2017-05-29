package main

import "net"
import "fmt"
import "bufio"
import "strings"
import "os"                     // only needed below for sample processing
//import "time"

func main() {

fmt.Println("Launching server...")

                                   // listen on all interfaces
ln, _ := net.Listen("localhost","tcp", ":8081")
fmt.Print("listening \n")
                                  // accept connection on port
  
fmt.Print("accepting \n")
                                 //state of a server
var state string

/* run loop forever (or until ctrl-c)
 ticker := time.NewTicker(time.Millisecond * 2000)
tickertimer for checking vote    */


fmt.Print("starting loop")
for electiontimer:=0;electiontimer<10;electiontimer++ {
message,_:= bufio.NewReader(conn).ReadString('\n')
conn.Close()
if (strings.TrimRight(message,"\n")!="votereq") {
fmt.Println("State:"+state)
conn, _ := net.Dial("tcp", "127.0.0.1:8081")
fmt.Fprintf(conn, "voteyes" + "\n")
os.Exit(1)
} 
go handleRequest(conn)
}	

fmt.Print("loop stopped") 
conn, _ := ln.Accept()
for  {
state="candidate"
                               //fmt.Println(state,"Tick at", t)
                                 // send to socket

 conn.Write([]byte("votereq" + "\n"))

                            // will listen for message to process ending in newline (\n)
    message,_:= bufio.NewReader(conn).ReadString('\n')
                           // output message received
    fmt.Print("Message Received:", string(message))
                           // sample process for string received
    

    //election detect 
      if strings.TrimRight(message, "\n")=="voteyes" { 
    // send new string back to client
   fmt.Print("State:"+state)
  
} else {
fmt.Print("I am not the leader" + "\n")
}
go handleRequest(conn)
}

}
/*handle incoming request
func handleRequest(conn net.Conn) {
conn, _ := ln.Accept()
conn.Close()
}*/
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
fmt.Println("err",reqLen)  
}
  // Send a response back to person contacting us.
  conn.Write([]byte("Message received."))
  // Close the connection when you're done with it.
  conn.Close()
}

   // time.Sleep(time.Millisecond * 1500)
  //  ticker.Stop()
 //   fmt.Println("Ticker stopped")



