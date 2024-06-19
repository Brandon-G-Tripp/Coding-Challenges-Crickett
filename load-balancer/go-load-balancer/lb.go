package main

import (
    "fmt"
    "io"
    "net"
)

type LoadBalancer struct {
    port uint16
    backendURL string
}

func NewLoadBalancer(port uint16, backendURL string) *LoadBalancer {
    return &LoadBalancer{port, backendURL}
}

func (lb *LoadBalancer) Start() error {
    listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", lb.port))
    if err != nil {
        return err
    }
    defer listener.Close()

    fmt.Printf("Load balancer listening on port %d\n", lb.port)

    for {
        conn, err := listener.Accept()
        if err != nil {
            return err
        }
        go lb.handleConnection(conn)
    }
}

func (lb *LoadBalancer) handleConnection(clientConn net.Conn) {
    defer clientConn.Close()

    var buffer [1024]byte
    n, err := clientConn.Read(buffer[:])
    if err != nil {
        fmt.Printf("Failed to read from client: %v\n", err)
        return 
    }

    fmt.Printf("Received request from %s\n", clientConn.RemoteAddr())
    fmt.Println(string(buffer[:n]))

    // Forward request to backend server
    backendConn, err := net.Dial("tcp", lb.backendURL)
    if err != nil {
        fmt.Printf("Failed to connect to backend: %v\n", err)
        return 
    }
    defer backendConn.Close()

    _, err = backendConn.Write(buffer[:n])
    if err != nil {
        fmt.Printf("Failed to write to backend: %v\n", err)
        return 
    }

    var response []byte
    for {
        var buf [1024]byte
        n, err := backendConn.Read(buf[:])
        if err != nil {
            if err != io.EOF {
                fmt.Printf("Failed to read from backend: %v\n", err)
            }
            break
        }
        response = append(response, buf[:n]...)
    }

    fmt.Printf("Received response from backend: %s\n", string(response))

    _, err = clientConn.Write(response)
    if err != nil {
        fmt.Printf("Failed to write response to client: %v\n", err)
    }
}
