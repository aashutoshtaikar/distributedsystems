package main

import (
    "fmt"
    "net"
    "bufio"
    "strings"
    "os"
    "time"
    "runtime"
    "sync"
)


var cTimer bool
var votesen string
var voterec int
var leader string
var heartBeat bool
var heartbeatTracker bool
var members int
var state string
var mutex = &sync.Mutex{}


func main () {
	var ports [5]string
	var address string
	
	cTimer=false
	heartbeatTracker=false
	runtime.GOMAXPROCS(100)
	address="127.0.0.1"
	ports[0]="8000"
	ports[1]="8001"
	ports[2]="8002"
	ports[3]="8003"
	ports[4]="8004"

	//c := make(chan int)
	hostserver(ports,address)
	var input string
	fmt.Scanln(&input)
	//Election(addres)
}

/*
	Setup the server
*/
func hostserver(ports [5]string, address string,) {

        go ListenPortConf(ports[0])
			for i :=1;i<5;i++ {
				go dialsoc(address,ports[i],ports[0])
			}
}

/*
	Dial connection
*/
func dialsoc(address string, ports string, port string,) {


		  loop:
		  servAddr := address+":"+ports
			  tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
			  if err != nil {
				  fmt.Println("ResolveTCPAddr failed:", err.Error())
					  os.Exit(1)
			  }
		  time.Sleep(1000 * time.Millisecond)

			  dial, err := net.DialTCP("tcp", nil, tcpAddr)

			  if err != nil {
				  //fmt.Println("Dial failed:", err.Error())
					  //os.Exit(1)
					  goto loop
			  }
		  go PeerComm(dial,ports)
			  fmt.Println("Connected to ",address,ports, " by ",port)
			  //fmt.Println("getting out of peer")

			for {

			}
}

/*
	Check and Communicate with Peers when Needed
*/
func PeerComm(conn net.Conn,ports string) {
		loop:
		if cTimer == true {

				if votesen == "" || votesen == "8000" {

					fVoteReq(conn,ports)
				} else {

					fmt.Println("Cannot request for vote as my state is",state)
				}
		} else {
			goto loop
		}
		if heartBeat == true {

			for {
				SendHeartBeat(conn, ports)
			}
		} else {
			fmt.Println("Receving heartbeat from ",votesen)
		}
		for {

		}
}

/* 
	Request For Votes from Peers
*/
func fVoteReq(conn net.Conn,ports string) {

	message :="RequestVote 8000"
	conn.Write([]byte(message + "\n"))
	fmt.Println("Vote requested from ",ports)
	message,_ =bufio.NewReader(conn).ReadString('\n')
	mutex.Lock()
	if strings.Contains(message,"Vote") {
		voterec++
		fmt.Println("Current votes received =",voterec)
		fmt.Println("vote received from",ports)
	}
	mutex.Unlock()
	time.Sleep(time.Second * 10)

	if voterec >= 3 && votesen == "8000"{
		mutex.Lock()
		fmt.Println("Majority votes received")
		fmt.Println("comunicating to peers...")
		state="leader"
		if(heartBeat == false) {

			heartBeat = true
		}
		mutex.Unlock()
		SendLeaderMessage(conn,ports)
	} else  {
		fmt.Println("Majority votes not received , voterec = ",voterec)
		state = "follower"
	}
}

/*
	Candidate Timeout
*/

func fCandTimer() {

	time.Sleep(1 * time.Second)

		if cTimer !=true {
			mutex.Lock()
			cTimer=true
			fmt.Println("Election timer timed out")
			if votesen == "" {
				// Become a candidate
				votesen ="8000"
				fmt.Println("State:Candidate")
			}
			mutex.Unlock()
		}

}


/*
	Process Received Message
*/
func fcheckmsg(message string) string{

		mutex.Lock()
		if strings.Contains(message,"RequestVote 8001") {
				if(votesen == "") {
					votesen="8001"
					message = "Vote"
					state = "follower"
					fmt.Println("voted for 8001")
				} else {
					message = "NoVote"
					fmt.Println("Cannot vote for 8001,as previously voted for ",votesen)
				}
		} else if strings.Contains(message,"RequestVote 8002") {
				if(votesen == "") {
					votesen="8002"
					message = "Vote"
					state = "follower"
					fmt.Println("voted for 8002")
				} else {
					message = "NoVote"
					fmt.Println("Cannot vote for 8002,as previously voted for ",votesen)
				}
		} else if strings.Contains(message,"RequestVote 8003") {
				if(votesen == "") {
					votesen="8003"
					message = "Vote"
					state = "follower"
					fmt.Println("voted for 8003")
				} else {
					message = "NoVote"
					fmt.Println("Cannot vote for 8003,as previously voted for ",votesen)
				}
		} else if strings.Contains(message,"RequestVote 8004") {
				if(votesen == "") {
					votesen="8004"
					message = "Vote"
					state = "follower"
					fmt.Println("voted for 8004")
				} else {
					message = "NoVote"
					fmt.Println("Cannot vote for 8004,as previously voted for ",votesen)
				}

		} else if strings.Contains(message,"Leader:8001") {

			leader="8001"
			message="Ack Leader"
		} else if strings.Contains(message,"Leader:8002") {

			leader="8002"
			message="Ack Leader"
		} else if strings.Contains(message,"Leader:8003") {

			leader="8003"
			message="Ack Leader"
		} else if strings.Contains(message,"Leader:8004") {

			leader="8004"
			message="Ack Leader"
		} else if strings.Contains(message,"heartbeat") && leader != "8000"{

			message="heartbeatAck" 
		}
		mutex.Unlock()
		return message
}

/*
	Monitor heartbeat
*/

func fCheckHeartBeat(conn net.Conn) {

		//timeoutDuration :=10 * time.Second
		//conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		for {
			time.Sleep(time.Second * 2)
			message, _:= bufio.NewReader(conn).ReadString('\n')

			if strings.Contains(message,"heartbeat") {
				if heartbeatTracker !=true {
						mutex.Lock()
						heartbeatTracker=true
						mutex.Unlock()
				}
				message ="heartbeatAck"
				conn.Write([]byte(message + "\n"))
				fmt.Println("heartbeat response Transmitted")

			} else if message == "" {

				fmt.Println("heartbeat not received")
			}
		}

}

/*
	Receive and Serve Messages for connected peer
*/
func ReceiveAndServe(conn net.Conn) {

	var message string
	for {

		message, _ = bufio.NewReader(conn).ReadString('\n')
		message1 := message
		message1 =fcheckmsg(message1)
		conn.Write([]byte(message + "\n"))
		if  strings.Contains(message,"heartbeat") && leader != "" {
				fmt.Println("Leader Elected =",leader)
				fCheckHeartBeat(conn)
		}

	}
}

/*
	Send hearbeats to followers
*/
func SendHeartBeat (conn net.Conn, ports string) {

		for {
			message :="heartbeat"
			conn.Write([]byte(message + "\n"))
			message, _ = bufio.NewReader(conn).ReadString('\n')
			if strings.Contains(message,"heartbeatAck") {

				fmt.Println("Received heartbeat Ack From ",ports)
			} else {

				fmt.Println("Missed an heartbeat Ack From",ports)
			}
		}
}

/*
	Send Leader Message over the connection
*/

func SendLeaderMessage(conn net.Conn,ports string) {

	message :="Leader:8000"
	conn.Write([]byte(message + "\n"))
	fmt.Println("Leader Update Message Sent to",ports)
	message,_=bufio.NewReader(conn).ReadString('\n')
    if strings.Contains(message,"Ack Leader") {
        fmt.Println("Successfully updated Leader information by",ports)
    } else {
        fmt.Println("Not Received Leader Ack from",ports)
    } 
}



/*
	Listen on Connected and ports and repond back
*/
func ListenAndServe (conn net.Conn, port string) {
		/*
			Increment Member Count And wait till Network is setup
		*/
		mutex.Lock()
		members++
		mutex.Unlock()
		loop:
		if members != 4 {

			goto loop
		} else {
			fmt.Println("All Members connected to socket,Candidate Timer Begins")
			go fCandTimer()
		}
		ReceiveAndServe(conn)
}




/*
	SetupListen Port
*/
func ListenPortConf(port string) {

		 service := ":"+port

			 tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
			 if err!= nil {
				 fmt.Println("Resovle failed ",err.Error())
					 os.Exit(1)
			 }
		 listener, err := net.ListenTCP("tcp", tcpAddr)
			 if err != nil {
				 // Handle error
				 fmt.Println(err.Error())
			 }

		 for {
				 conn, err := listener.Accept()
				 if err != nil {
					 // Handle error
					 fmt.Println(err.Error())
				 }
			 go ListenAndServe(conn,port)
		 }
		for {

		}
}