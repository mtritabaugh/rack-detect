package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	netCidr1 := os.Getenv("NET_CIDR1")
	netCidr2 := os.Getenv("NET_CIDR2")
	netCidr3 := os.Getenv("NET_CIDR3")

	bgp_peers1 := os.Getenv("BGP_PEERS1")
	bgp_peers2 := os.Getenv("BGP_PEERS2")
	bgp_peers3 := os.Getenv("BGP_PEERS3")

	bgp_as1 := os.Getenv("BGP_AS1")
	bgp_as2 := os.Getenv("BGP_AS2")
	bgp_as3 := os.Getenv("BGP_AS3")

	ip := os.Getenv("IP")

	_, netCidr1P, _ := net.ParseCIDR(netCidr1)
	_, netCidr2P, _ := net.ParseCIDR(netCidr2)
	_, netCidr3P, _ := net.ParseCIDR(netCidr3)
	ipP := net.ParseIP(ip)

	switch {
	case netCidr1P.Contains(ipP):
		fmt.Printf("#!/bin/sh\n")
		fmt.Printf("export bgp_as=%s\n", bgp_as1)
		fmt.Printf("export bgp_peers=%s\n", bgp_peers1)
		fmt.Printf("/kube-vip manager\n")
	case netCidr2P.Contains(ipP):
		fmt.Printf("#!/bin/sh\n")
		fmt.Printf("export bgp_as=%s\n", bgp_as2)
		fmt.Printf("export bgp_peers=%s\n", bgp_peers2)
		fmt.Printf("/kube-vip manager\n")
	case netCidr3P.Contains(ipP):
		fmt.Printf("#!/bin/sh\n")
		fmt.Printf("export bgp_as=%s\n", bgp_as3)
		fmt.Printf("export bgp_peers=%s\n", bgp_peers3)
		fmt.Printf("/kube-vip manager\n")
	default:
		fmt.Println("not on rack")
	}
}
