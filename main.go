package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "io"
)

func main() {
    if len(os.Args) != 5 {
        fmt.Println("Usage: ./port-forwarder <source-addr> <source-port> <destination-addr> <destination-port>")
        os.Exit(1)
    }

    sourceAddr := os.Args[1]
    sourcePort := os.Args[2]
    destAddr := os.Args[3]
    destPort := os.Args[4]

    listener, err := net.Listen("tcp", sourceAddr+":"+sourcePort)
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    log.Printf("Listening on %s:%s and forwarding to %s:%s...\n", sourceAddr, sourcePort, destAddr, destPort)

    for {
        sourceConn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go forwardToDestination(sourceConn, destAddr+":"+destPort)
    }
}

func forwardToDestination(sourceConn net.Conn, destAddr string) {
    destConn, err := net.Dial("tcp", destAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer destConn.Close()

    go func() {
        defer sourceConn.Close()

        _, err := io.Copy(destConn, sourceConn)
        if err != nil {
            log.Fatal(err)
        }
    }()

    go func() {
        defer destConn.Close()

        _, err := io.Copy(sourceConn, destConn)
        if err != nil {
            log.Fatal(err)
        }
    }()
}

