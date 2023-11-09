package main

import (
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
    if len(os.Args) != 5 {
        fmt.Println("Usage: ./forwarder <source-interface> <source-port> <destination-interface> <destination-port>")
        os.Exit(1)
    }

    sourceInterface := os.Args[1]
    sourcePort := os.Args[2]
    destInterface := os.Args[3]
    destPort := os.Args[4]

    sourceAddr := fmt.Sprintf("%s:%s", sourceInterface, sourcePort)
    destAddr := fmt.Sprintf("%s:%s", destInterface, destPort)

    sourceConn, err := net.ListenPacket("tcp", sourceAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer sourceConn.Close()

    log.Printf("Listening on %s...\n", sourceAddr)

    for {
        buf := make([]byte, 1500)
        n, addr, err := sourceConn.ReadFrom(buf)
        if err != nil {
            log.Fatal(err)
        }

        go forwardPacket(buf[:n], addr, destAddr)
    }
}

func forwardPacket(data []byte, sourceAddr net.Addr, destAddr string) {
    destConn, err := net.Dial("udp", destAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer destConn.Close()

    _, err = destConn.Write(data)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Forwarded packet from %s to %s\n", sourceAddr, destAddr)
}

