// Package main where the main function of the application are
package main

import (
	"fmt"

	//"github.com/jeremydumais/iplookup/ipaddress"
    "github.com/alexflint/go-arg"
)

type args struct {
    IPAddress string `arg:"positional"`
}

func (args) Version() string {
	return "iplookup 0.1"
}

func main() {
    var args args
    arg.MustParse(&args)

    //addr := ipaddress.CreateIPAddress(1, 2, 3, 4)
    fmt.Printf("Looking info for : %v\n", args.IPAddress)
}
