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

type resmsg struct {
    Header int //it will count how many requests are served
    Body []byte
}

func main() {

    start := time.Now() //program starts

    sigs := make(chan os.Signal, 1) //make channel
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM) //that can get ^C signal
    
    serverPort := "12000"

    listener, _:= net.Listen("tcp", ":" + serverPort) //listening for incoming TCP requests
    fmt.Printf("The server is ready to receive on port %s\n", serverPort)

    go func() { //if sigs gets msg(^C), terminate server and exit
        <-sigs
        fmt.Println("Bye bye~") 
        os.Exit(0) //terminate program(dont execute defer)
    }()

    buffer := make([]byte, 1024)
    packet := resmsg{0, []byte{}}

    for {

        conn, connerr:= listener.Accept() //waits on accept() for incoming requests, create socket on return
        if nil != connerr {
            fmt.Println("Can't connect to client")
            continue
        }
        
        go func() { //if sigs gets msg(^C), terminate server
            <-sigs
            return
        }()

        defer func () { //execute before returning main func
            conn.Close() //close socket
            listener.Close() //close listener
            fmt.Println("Bye bye~") 
        }()

        for {
            count, err := conn.Read(buffer) //read from socket
            option := string(buffer[:1][0]) //get client's option
            if nil != err {
                fmt.Printf("Disconnected from client: %s"+"\n", conn.RemoteAddr().String())
                break
            }
          
            fmt.Printf("Connection request from %s\n", conn.RemoteAddr().String()) //print client's address from connection
            fmt.Printf("Command %s\n\n", option)
    
            switch option {
            case "1":
                packet.Header += 1
                packet.Body = bytes.ToUpper(buffer[1:count])
                conn.Write(packet.Body) //make upper case except option and send it
            case "2":
                remote := conn.RemoteAddr().String()
                packet.Header += 1
                packet.Body = []byte(remote)
                conn.Write(packet.Body) //send client's address
            case "3": //how many clients requests it has served so far
                packet.Header += 1
                packet.Body = []byte(strconv.Itoa(packet.Header))
                conn.Write(packet.Body)
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

                packet.Header += 1
                packet.Body = []byte(hhmmss)
                conn.Write(packet.Body) //send how long the server has been running
            
            default: continue
            }
        }
        
    }
}

