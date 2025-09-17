package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	Validate()
	os.Exit(0)
}

func Validate() {
	if len(os.Args) < 2 {
		fmt.Println("\x1b[93mUsage: validate <ip>\x1b[0m")
		return
	}

	ip := os.Args[1]
	isValid := validateIP(ip)

	if isValid {
		fmt.Printf("\x1b[93m%s\x1b[97m is a \x1b[92mvalid IP address.\x1b[0m\n", ip)
	} else {
		fmt.Printf("\x1b[93m%s\x1b[91m is \x1b[91mnot a valid IP address.\x1b[0m\n", ip)
	}
}

func validateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}
