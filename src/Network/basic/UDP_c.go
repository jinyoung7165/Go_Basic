/**
 * name: Choi Jin Young
 * sID: 50223038
**/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)


func showMenu() {
    fmt.Println("<Menu>")
    fmt.Println("1) convert text to UPPER-case")
    fmt.Println("2) get my IP address and port number")
    fmt.Println("3) get server request count")
    fmt.Println("4) get server running time")
    fmt.Println("5) exit")
}

func checkServer(buffer *[]byte, count int) bool {
    if strings.Compare(string((*buffer)[:count]), "error: server terminated") == 0 { //if server is terminated -> close UDP socket and exit client
        fmt.Println("Diconnected")
        return true
    }
    return false
}

func main() {

    sigs := make(chan os.Signal, 1) //make channel
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM) //that can get ^C signal

    serverName := "nsl2.cau.ac.kr"
    serverPort := "23038"

    pconn, _:= net.ListenPacket("udp", ":") //create UDP socket

    localAddr := pconn.LocalAddr().(*net.UDPAddr) //local address
    fmt.Printf("The client is running on port %d\n", localAddr.Port) //check own port
    
    server_addr, _ := net.ResolveUDPAddr("udp", serverName+":"+serverPort) //get server address object

    var option int
    readBuffer := make([]byte, 1024)

    go func() { //IF sigs channel gets ^C
		<-sigs
        pconn.Close() //close socket
        fmt.Println("Bye bye~") 
		os.Exit(0) //terminate program(dont execute defer)
	}()

    defer func() { //execute before returning main
        pconn.Close() //close socket
        fmt.Println("Bye bye~") 
    }()

    for {
       
        showMenu()
        fmt.Printf("Input option: ")
  
        fmt.Scanf("%d\n",&option) //get option from keyboard

        switch option {
        case 1: //convert text to UPPER case
            fmt.Printf("Input sentence: ")
            input, _ := bufio.NewReader(os.Stdin).ReadString('\n') //readOneline

            sent_t := time.Now() //sent time measured before sending the command
            pconn.WriteTo([]byte("1"+input), server_addr) //send a packet with payload(option+input) to socket

            count, _, _ := pconn.ReadFrom(readBuffer) //read reply from socket
            recv_t := time.Since(sent_t) //how long it took to receive a packet after sending a packet
            if tf := checkServer(&readBuffer, count); tf { return } //check whether server is terminated

            rtt := float64(recv_t.Microseconds()) / 1e3 //micro to milli
            fmt.Println()
            fmt.Printf("Reply from server: %s", string(readBuffer[:count])) //print as much as received
            fmt.Println("RTT =", rtt, "ms") //print rtt
            fmt.Println()

        case 2: //ask the server what the IP and port of the client is
            sent_t := time.Now()
            pconn.WriteTo([]byte("2"), server_addr) //send a packet with payload(option) to socket

            count, _, _ := pconn.ReadFrom(readBuffer)
            recv_t := time.Since(sent_t)
            if tf := checkServer(&readBuffer, count); tf { return }

            l_add := strings.Split(string(readBuffer[:count]), ":") //split by ":"
            rtt := float64(recv_t.Microseconds()) / 1e3
            fmt.Println()
            fmt.Printf("Reply from server: client IP = %s, port = %s"+"\n", l_add[0], l_add[1])
            fmt.Println("RTT =", rtt, "ms") //print rtt
            fmt.Println()

        case 3: //ask the server how many client commands it has served so dar
            sent_t := time.Now()
            pconn.WriteTo([]byte("3"), server_addr)

            count, _, _ := pconn.ReadFrom(readBuffer)
            recv_t := time.Since(sent_t)
            if tf := checkServer(&readBuffer, count); tf { return }

            rtt := float64(recv_t.Microseconds()) / 1e3
            fmt.Println()
            fmt.Printf("Reply from server: requests served = %s"+"\n", string(readBuffer[:count]))
            fmt.Println("RTT =", rtt, "ms") //print rtt
            fmt.Println()

        case 4: //ask the server program how long it has been running for since it started(seconds)
            sent_t := time.Now()
            pconn.WriteTo([]byte("4"), server_addr)

            count, _, _ := pconn.ReadFrom(readBuffer)
            recv_t := time.Since(sent_t)
            if tf := checkServer(&readBuffer, count); tf { return }

            rtt := float64(recv_t.Microseconds()) / 1e3
            fmt.Println()
            fmt.Printf("Reply from server: run time = %s"+"\n", string(readBuffer[:count])) //HH:MM:SS
            fmt.Println("RTT =", rtt, "ms") //print rtt
            fmt.Println()

        case 5:
            return

        default:
            fmt.Println()
        }
    }

}

