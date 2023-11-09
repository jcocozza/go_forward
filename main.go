package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "io"
)

func main() {
    // CHEAP way to get all the args in
    if len(os.Args) != 5 {
        fmt.Println("Usage: ./forwarder <source-interface-ip-addr> <source-port> <destination-interface-ip-addr> <destination-port>")
        os.Exit(1)
    }

    sourceInterface := os.Args[1]
    sourcePort := os.Args[2]
    destInterface := os.Args[3]
    destPort := os.Args[4]

    sourceAddr := fmt.Sprintf("%s:%s", sourceInterface, sourcePort)
    destAddr := fmt.Sprintf("%s:%s", destInterface, destPort)

    sourceListener, err := net.Listen("tcp", sourceAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer sourceListener.Close()

    log.Printf("Listening on %s...\n", sourceAddr)

    for {
        sourceConn, err := sourceListener.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go forwardConnection(sourceConn, destAddr)
    }

}

func forwardConnection(sourceConn net.Conn, destAddr string) {
    destConn, err := net.Dial("tcp", destAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer destConn.Close()
    
    log.Printf("Forwarding from %s to %s...\n", sourceConn.RemoteAddr(), destConn.RemoteAddr())

    go func() {
        _, err := io.Copy(destConn, sourceConn)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Forwarded from %s to %s\n", sourceConn.RemoteAddr(), destConn.RemoteAddr())
    }()

    _, err = io.Copy(sourceConn, destConn)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Success! Forwarded from %s to %s\n", destConn.RemoteAddr(), sourceConn.RemoteAddr())
}

