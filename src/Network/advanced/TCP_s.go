/**
 * name: Choi Jin Young
 * sID: 50223038
**/

package main

import (
	"bytes"
	"fmt"
	"math"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)


func hHmMsS(t int) string{ //seconds to HH:MM:SS format
    if t < 10 { //append "0" in front of t
        return "0" + strconv.Itoa(t)
    } else {
        return strconv.Itoa(t)
    }
}

func routineHandle(conn net.Conn, start time.Time, rCount *int, clientId int, clientSum *int) {
    buffer := make([]byte, 1024)
    packet := make([]byte, 1024)
    
    for {
        count, err := conn.Read(buffer) //read from socket
        option := string(buffer[:1][0]) //get client's option
        if nil != err {
            fmt.Printf("Disconnected from client: %s"+"\n", conn.RemoteAddr().String())
            *clientSum -= 1
            fmt.Printf("Client %d disconnected. Number of connected clients = %d"+"\n", clientId, *clientSum)
            break
        }
      
        fmt.Printf("Connection request from %s\n", conn.RemoteAddr().String()) //print client's address from connection
        fmt.Printf("Command %s\n\n", option)

        switch option {
        case "1":
            *rCount += 1
            packet = bytes.ToUpper(buffer[1:count])
            conn.Write(packet) //make upper case except option and send it
        case "2":
            remote := conn.RemoteAddr().String()
            *rCount += 1
            packet = []byte(remote)
            conn.Write(packet) //send client's address
        case "3": //how many clients requests it has served so far
            *rCount += 1
            packet = []byte(strconv.Itoa(*rCount))
            conn.Write(packet)
        case "4":
            end := time.Since(start) //how long the server has been running
            //make it h,m,s form
            seconds := int(end.Seconds())
            hours := math.Floor(float64(seconds) / 60 / 60) //get hours
            seconds = seconds % (60 * 60)
            minutes := math.Floor(float64(seconds) / 60) //get minutes
            seconds = seconds % 60 //get seconds
            //make it hh, mm, ss form
            s_hours := hHmMsS(int(hours))
            s_minutes := hHmMsS(int(minutes))
            s_seconds := hHmMsS(seconds)
            hhmmss := s_hours + ":" + s_minutes + ":" + s_seconds

            *rCount += 1
            packet = []byte(hhmmss)
            conn.Write(packet) //send how long the server has been running
        
        default: continue
        }
    }
}

func main() {

    start := time.Now() //program starts
    rCount := 0 //count requests
    clientId, clientSum := 0, 0 //which connected at that time, total num of clients connecting

    sigs := make(chan os.Signal, 1) //make channel
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM) //that can get ^C signal
    
    serverPort := "23038"

    listener, _:= net.Listen("tcp", ":" + serverPort) //listening for incoming TCP requests
    fmt.Printf("The server is ready to receive on port %s\n", serverPort)

    go func() { //if sigs gets msg(^C), terminate server and exit
        <-sigs
        fmt.Println("Bye bye~") 
        os.Exit(0) //terminate program(dont execute defer)
    }()


    go func() { //every 1 min, server prints out clientSum
        for {
            <-time.After(1 * time.Minute)
            fmt.Println("Number of connected clients =", clientSum)
            fmt.Println()
        }
    }()

    for {
        
        conn, connerr:= listener.Accept() //waits on accept() for incoming requests, create socket on return
        if nil != connerr {
            fmt.Println("Can't connect to client")
            continue
        }
        
        clientId += 1 //which connected now
        clientSum += 1 //sum of connected clients
        fmt.Printf("Client %d connected. Number of connected clients = %d\n", clientId, clientSum)

        go func() { //if sigs gets msg(^C), terminate server
            <-sigs
            return
        }()

        defer func () { //execute before returning main func
            conn.Close() //close socket
            listener.Close() //close listener
            fmt.Println("Bye bye~") 
        }()

        go routineHandle(conn, start, &rCount, clientId, &clientSum)
    }
}

