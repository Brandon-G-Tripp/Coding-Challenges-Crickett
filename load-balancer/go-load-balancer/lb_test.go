package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func getAvailablePort() uint16 {
    listener, err := net.Listen("tcp", "127.0.0.1:0")
    if err != nil {
        panic(err)
    }
    defer listener.Close()
    return uint16(listener.Addr().(*net.TCPAddr).Port)
}

func startBackendServer(port uint16) *exec.Cmd {
    currentDir, _ := os.Getwd()
    parentDir := currentDir + "/.."
    pythonScriptPath := parentDir + "/test_backend_server.py"

    fmt.Printf("Attempting to start backend server with script: %s\n", pythonScriptPath)

    cmd := exec.Command("python3", pythonScriptPath, fmt.Sprintf("%d", port))
    err := cmd.Start()
    if err != nil {
        panic(fmt.Sprintf("Failed to start backend server: %v", err))
    }

    time.Sleep(1 * time.Second)

    return cmd
}

func TestLoadBalancerAcceptsConnection(t *testing.T) {
    backendPort := getAvailablePort()
    backendServer := startBackendServer(backendPort)
    defer backendServer.Process.Kill()

    lbPort := getAvailablePort()
    lb := NewLoadBalancer(lbPort, fmt.Sprintf("127.0.0.1:%d", backendPort))

    go func() {
        err := lb.Start()
        if err != nil {
            t.Errorf("Failed to start load balancer: %v", err)
        }
    }()

    time.Sleep(2 * time.Second)

    conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", lbPort))
    if err != nil {
        t.Fatalf("Failed to connect to load balancer: %v", err)
    }
    defer conn.Close()

    _, err = conn.Write([]byte("GET / HTTP/1.0\r\nHost: localhost\r\n\r\n"))
    if err != nil {
        t.Fatalf("Failed to send request to load balancer: %v", err)
    }

    var response []byte
    for {
        var buf [1024]byte
        n, err := conn.Read(buf[:])
        if err != nil {
            if err != io.EOF {
                t.Fatalf("Failed to read response from load balancer: %v", err)
            }
            break
        }
        response = append(response, buf[:n]...)
    }

    // Verify the response
    responseStr := string(response)

    // check status code
    if !strings.Contains(responseStr, "HTTP/1.0 200 OK") {
        t.Errorf("Expected status code 200, got: %s", responseStr)
    }

    // check content type
    if !strings.Contains(responseStr, "Content-type: text/plain") {
        t.Errorf("Expected Content-Type: text/plain, got: %s", responseStr)
    }

    // check body content

    expectedBody := fmt.Sprintf("Hello from backend server on port %d", backendPort)
    if !strings.Contains(responseStr, expectedBody) {
        t.Errorf("Unexpected response from load balancer. Got: %s, Expected: %s", expectedBody, responseStr)
    }
}


