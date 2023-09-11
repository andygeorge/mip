package main

import (
	"fmt"

	externalip "github.com/glendc/go-external-ip"
)

func main() {
	ipVs := []uint{4, 6}

	for _, ipV := range ipVs {
		c := externalip.DefaultConsensus(nil, nil)
		c.UseIPProtocol(ipV)
		ex, err := c.ExternalIP()
		if err == nil {
			ip := ex.String()
			fmt.Println(ip)
		}
	}
}
