package udp

import (
    "fmt"
    "log"
    "net"
)

func Udp(sourceInterface string, sourcePort string, destInterface string, destPort string) {

    sourceAddr := fmt.Sprintf("%s:%s", sourceInterface, sourcePort)
    destAddr := fmt.Sprintf("%s:%s", destInterface, destPort)

    sourceConn, err := net.ListenPacket("udp", sourceAddr)
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
