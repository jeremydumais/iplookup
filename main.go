// Package main where the main function of the application are
package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/jeremydumais/iplookup/ipaddress"
)

type args struct {
    IPAddress string `arg:"positional" help:"The IPv4 address to lookup"`
}

func (args) Version() string {
	return "iplookup 0.1"
}

func main() {
    os.Exit(mainReturnWithCode())
}

func mainReturnWithCode() int {
    var args args
    arg.MustParse(&args)

    addr, err := ipaddress.ParseIPAddress(args.IPAddress)
    if err != nil {
        fmt.Printf("Unable to convert %v to an IPv4 Address\n", args.IPAddress)
        return 1;
    }
    fmt.Printf("IP Address : %v\n", addr)
    fmt.Printf("Class: %v\n", addr.GetClass().String())

    var isPrivate string
    if addr.IsPrivate() {
        isPrivate = "true"
    } else {
        isPrivate = "false"
    }
    fmt.Printf("Is Private: %v\n", isPrivate)
    return 0;
}
