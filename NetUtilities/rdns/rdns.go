package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ReverseDNS()
	os.Exit(0)
}

func ReverseDNS() {
	if len(os.Args) < 2 {
		fmt.Println("\x1b[93mUsage\x1b[97m rdns <ip>")
		return
	}

	ip := os.Args[1]
	hostnames, err := net.LookupAddr(ip)
	if err != nil {
		fmt.Println("\x1b[93mError\x1b[97m \x1b[91m", err)
		return
	}

	fmt.Println("\x1b[93mHostnames\x1b[97m:")
	for _, hostname := range hostnames {
		fmt.Println("\x1b[93m  -\x1b[97m", hostname)
	}
}
