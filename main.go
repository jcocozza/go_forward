package main

import (
	"fmt"
	"log"
    "os"

	"github.com/jcocozza/go_foward/tcp"
	"github.com/jcocozza/go_foward/udp"
)

func main() {
    if len(os.Args) != 6 {
        fmt.Println("Usage: ./port-forwarder <type(tcp-udp)> <source-ip-addr> <source-port> <destination-addr> <destination-port>")
        os.Exit(1)
    }
    dataType := os.Args[1]
    sourceInterface := os.Args[2]
    sourcePort := os.Args[3]
    destInterface := os.Args[4]
    destPort := os.Args[5]

    if dataType == "udp" {
        udp.Udp(sourceInterface, sourcePort, destInterface, destPort)
    } else if dataType == "tcp" {
        tcp.Tcp(sourceInterface, sourcePort, destInterface, destPort)
    } else {
        log.Fatal("improper data type")
    }
}