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


func showmenu() {
    fmt.Println("<Menu>")
    fmt.Println("1) convert text to UPPER-case")
    fmt.Println("2) get my IP address and port number")
    fmt.Println("3) get server request count")
    fmt.Println("4) get server running time")
    fmt.Println("5) exit")
}

func main() {

    sigs := make(chan os.Signal, 1) //make channel
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM) //that can get ^C signal

    serverName := "127.0.0.1"
    serverPort := "12000"

    conn, err:= net.Dial("tcp", serverName+":"+serverPort) //create and connect TCP socket
    if err != nil {
        fmt.Println("Can't connect to server")
        return
    }

    localAddr := conn.LocalAddr().(*net.TCPAddr)

    fmt.Printf("The client is running on port %d\n", localAddr.Port) //check own port

    var option int


    defer func() { //execute before returning main
        conn.Close() //close socket
        fmt.Println()
        fmt.Println("Bye bye~")
    }()

    go func() { //IF sigs channel gets ^C
		<-sigs
		conn.Close() //close socket
        fmt.Println("Bye bye~") 
        os.Exit(0) //don't execute defer
	}()

    readBuffer := make([]byte, 1024)

    for {

        showmenu()
        fmt.Printf("Input option: ")

        fmt.Scanf("%d\n",&option) //get option from keyboard

        switch option {
        case 1: //convert text to UPPER case
            fmt.Printf("Input sentence: ")
            input, _ := bufio.NewReader(os.Stdin).ReadString('\n') //readOneline

            sent_t := time.Now() //sent time measured before sending the command
            _, err := conn.Write([]byte("1"+input)) //send a packet with payload(option+input) to socket

            if err != nil {
                fmt.Print("Disconnected")
                return
            }

            count, readErr := conn.Read(readBuffer) //read reply
            recv_t := time.Since(sent_t) //how long it took to receive a packet after sending a packet
            if readErr != nil {
                fmt.Print("Disconnected")
                return
            }

            rtt := float64(recv_t.Microseconds()) / 1e3 //micro to milli
            fmt.Println()
            fmt.Printf("Reply from server: %s", string(readBuffer[:count])) //print as much as received
            fmt.Println("RTT =", rtt, "ms") //print rtt
            fmt.Println()
            
        case 2: //ask the server what the IP and port of the client is
            sent_t := time.Now()
            conn.Write([]byte("2")) //send a packet with payload(option) to socket

            count, err := conn.Read(readBuffer)
            recv_t := time.Since(sent_t)
            if err != nil {
                fmt.Print("Disconnected")
                return
            }

            l_add := strings.Split(string(readBuffer[:count]), ":") //split by ":"
            rtt := float64(recv_t.Microseconds()) / 1e3
            fmt.Println()
            fmt.Printf("Reply from server: client IP = %s, port = %s"+"\n", l_add[0], l_add[1])
            fmt.Println("RTT =", rtt, "ms") //print rtt
            fmt.Println()
        
        case 3: //ask the server how many client commands it has served so far
            sent_t := time.Now()
            conn.Write([]byte("3"))

            count, err := conn.Read(readBuffer)
            recv_t := time.Since(sent_t)
            if err != nil {
                fmt.Print("Disconnected")
                return
            }

            rtt := float64(recv_t.Microseconds()) / 1e3
            fmt.Println()
            fmt.Printf("Reply from server: requests served = %s"+"\n", string(readBuffer[:count]))
            fmt.Println("RTT =", rtt, "ms") //print rtt
            fmt.Println()
        
        case 4: //ask the server program how long it has been running for since it started(seconds)
            sent_t := time.Now()
            conn.Write([]byte("4"))

            count, err := conn.Read(readBuffer)
            recv_t := time.Since(sent_t)
            if err != nil {
                fmt.Print("Disconnected")
                return
            }

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
