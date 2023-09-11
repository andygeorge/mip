package main

import (
	"fmt"

	externalip "github.com/glendc/go-external-ip"
)

func main() {
	ipVs := []uint{4, 6}

	for _, ipV := range ipVs {
		externalipConsensus := externalip.DefaultConsensus(nil, nil)
		externalipConsensus.UseIPProtocol(ipV)
		externalipIp, err := externalipConsensus.ExternalIP()
		if err == nil {
			myIp := externalipIp.String()
			fmt.Println(myIp)
		}
	}
}

func errorHandler(err error) {
	panic(err)
}
