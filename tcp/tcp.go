package tcp

import (
    "log"
    "net"
    "io"
)

func Tcp(sourceInterface string, sourcePort string, destInterface string, destPort string) {


    listener, err := net.Listen("tcp", sourceInterface+":"+sourcePort)
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    log.Printf("Listening on %s:%s and forwarding to %s:%s...\n", sourceInterface, sourcePort, destInterface, destPort)

    for {
        sourceConn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go forwardToDestination(sourceConn, destInterface+":"+destPort)
    }
}

func forwardToDestination(sourceConn net.Conn, destInterface string) {
    destConn, err := net.Dial("tcp", destInterface)
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

