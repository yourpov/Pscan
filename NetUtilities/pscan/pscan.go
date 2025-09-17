package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	Pscan()
	os.Exit(0)
}

func Pscan() {
	if len(os.Args) != 3 {
		fmt.Println("\x1b[93mUsage: ./pscan <host> <maxPort>\x1b[0m")
		return
	}

	host := os.Args[1]
	ports := os.Args[2]

	fmt.Printf("Scanning host %s...\n", host)
	scanPorts(host, ports)
}

func scanPorts(host, ports string) {
	portList := parsePortList(ports)
	for port := 1; port <= portList; port++ {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, time.Second)

		if err != nil {
			fmt.Printf("\x1b[91mPort\x1b[97m: \x1b[93m%d\x1b[91m is closed\x1b[0m\n", port)
		} else {
			fmt.Printf("\x1b[92mPort\x1b[97m: \x1b[93m%d\x1b[92m is open\x1b[0m\n", port)
			conn.Close()
		}
	}
}

func parsePortList(ports string) int {
	port, err := strconv.Atoi(ports)
	if err != nil {
		fmt.Printf("\x1b[93mInvalid port\x1b[97m: \x1b[93m%s\n\x1b[91m\x1b[0m", ports)
		os.Exit(1)
	}
	return port
}
