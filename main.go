package main

import (
	"fmt"

	externalip "github.com/glendc/go-external-ip"
)

func main() {
	externalipConsensus := externalip.DefaultConsensus(nil, nil)
	externalipIp, err := externalipConsensus.ExternalIP()
	if err != nil {
		errorHandler(err)
	}
	myIp := externalipIp.String()

	fmt.Println(myIp)
}

func errorHandler(err error) {
	panic(err)
}
