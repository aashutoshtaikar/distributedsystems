package main

import "net"
import "fmt"
import "bufio"
import "strings"


func main() {

// connect to this socket

for  {
conn, _ := net.Dial("tcp", "127.0.0.1:2000") 

conn.Write([]byte( "e"+ "\n"))
conn.Close()
message, _ := bufio.NewReader(conn).ReadString('\n') 
if strings.TrimRight(message,"\n")=="votereq" {
 // send vote to socket
    fmt.Fprintf(conn, "voteyes" + "\n")
 // listen for reply
message, _ := bufio.NewReader(conn).ReadString('\n') 
fmt.Print("Message from server: "+message)
}
}

}
