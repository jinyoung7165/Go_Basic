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

func HHMMSS(t int) string{ //seconds to HH:MM:SS format
    if t < 10 { //append "0" in front of t
        return "0" + strconv.Itoa(t)
    } else {
        return strconv.Itoa(t)
    }
}

type RESMSG struct {
    Header int //it will count how many requests are served
    Body []byte
}

func main() {

    start := time.Now() //program starts

    sigs := make(chan os.Signal, 1) //make channel
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM) //that can get ^C signal
    
    serverPort := "23038"

    pconn, _:= net.ListenPacket("udp", ":"+serverPort) //create UDP socket
    fmt.Printf("The server is ready to receive on port %s\n", serverPort)

    buffer := make([]byte, 1024)
    packet := RESMSG{0, []byte{}}
    
    for {
        var option string
        count, r_addr, err:= pconn.ReadFrom(buffer) //Read from UDP socket into buffer, getting client's address
        if err == nil {
            option = string(buffer[:1][0]) //get client's option
        } else { continue }

        go func() { //if sigs gets msg(^C), send err message and exit
            <-sigs
            if r_addr != nil {
                err := "error: server terminated"
                pconn.WriteTo([]byte(err), r_addr)
            }
            pconn.Close() //close socket
            fmt.Println("Bye bye~") 
            os.Exit(0) //terminate program
        }()

        if r_addr != nil {
            fmt.Printf("Connection request from %s\n", r_addr.String())
            fmt.Printf("Command %s\n\n", option)
        }


        switch option {
        case "1":
            packet.Header += 1
            packet.Body = bytes.ToUpper(buffer[1:count])
            pconn.WriteTo(packet.Body, r_addr) //make upper case except option and send it
        case "2":
            remote := r_addr.String()
            packet.Header += 1
            packet.Body = []byte(remote)
            pconn.WriteTo(packet.Body, r_addr) //send client's address
        case "3": //how many clients requests it has served so far
            packet.Header += 1
            packet.Body = []byte(strconv.Itoa(packet.Header))
            pconn.WriteTo(packet.Body, r_addr)
        case "4":
            end := time.Since(start) //how long the server has been running
            //make it h,m,s form
            seconds := int(end.Seconds())
            hours := math.Floor(float64(seconds) / 60 / 60) //get hours
            seconds = seconds % (60 * 60)
            minutes := math.Floor(float64(seconds) / 60) //get minutes
            seconds = seconds % 60 //get seconds
            //make it hh, mm, ss form
            s_hours := HHMMSS(int(hours))
            s_minutes := HHMMSS(int(minutes))
            s_seconds := HHMMSS(seconds)
            hhmmss := s_hours + ":" + s_minutes + ":" + s_seconds

            packet.Header += 1
            packet.Body = []byte(hhmmss)
            pconn.WriteTo(packet.Body, r_addr) //send how long the server has been running
            
        default: continue
        }
    }

}

