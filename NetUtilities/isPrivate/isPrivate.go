package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\x1b[93mUsage: ./isPrivate <host>\x1b[0m")
		return
	}

	ipAddr := os.Args[1]
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		fmt.Println("\x1b[91mInvalid IP address\x1b[97m.\x1b[0m")
		os.Exit(1)
	}

	network := Results(ip)
	fmt.Printf("\x1b[93mIP Address\x1b[97m: \x1b[97m%s\n", ip.String())
	fmt.Printf("\x1b[93mNetwork\x1b[97m: \x1b[97m%s\n", network.String())
	fmt.Printf("\x1b[93mIs Loopback\x1b[97m: \x1b[97m%t\n", network.IP.IsLoopback())
	fmt.Printf("\x1b[93mIs Private\x1b[97m: \x1b[97m%t\x1b[0m\n", isPrivate(network.IP))
}

func prompt(promptText string) string {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Results(ip net.IP) *net.IPNet {
	ones, bits := ip.DefaultMask().Size()
	network := &net.IPNet{
		IP:   ip.Mask(net.CIDRMask(ones, bits)),
		Mask: ip.DefaultMask(),
	}
	return network
}

func isPrivate(ip net.IP) bool {
	privateIPBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	for _, block := range privateIPBlocks {
		_, privateNet, _ := net.ParseCIDR(block)
		if privateNet.Contains(ip) {
			return true
		}
	}

	return false
}
