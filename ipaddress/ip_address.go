// Package ipaddress used to store the different project models
package ipaddress

import (
    "errors"
    "fmt"
    "strings"
)

// The IPAddress struct used by the project
type IPAddress struct {
    byte1 uint8
    byte2 uint8
    byte3 uint8
    byte4 uint8
}

// CreateIPAddress creates a new IPAddress struct
func CreateIPAddress(byte1 uint8,
                     byte2 uint8,
                     byte3 uint8,
                     byte4 uint8) IPAddress {
    return IPAddress{byte1: byte1,
                     byte2: byte2,
                     byte3: byte3,
                     byte4: byte4}
}

// GetByte1 returns the first byte of the IPAddress
func (i IPAddress) GetByte1() uint8 {
    return i.byte1
}

// GetByte2 returns the second byte of the IPAddress
func (i IPAddress) GetByte2() uint8 {
    return i.byte2
}

// GetByte3 returns the third byte of the IPAddress
func (i IPAddress) GetByte3() uint8 {
    return i.byte3
}

// GetByte4 returns the fourth byte of the IPAddress
func (i IPAddress) GetByte4() uint8 {
    return i.byte4
}

// Display the IPAddress as a String
func (i IPAddress) String() string {
    return fmt.Sprintf("%v.%v.%v.%v",
    i.byte1, i.byte2, i.byte3, i.byte4)
}

// ParseIPAddress try to parse a string value into an IPAddress struct
func ParseIPAddress(value string) (IPAddress, error) {
    parts := strings.Split(value, ".")
    if len(parts) != 4 {
        return IPAddress{}, errors.New("An IP Address element must have 4 parts separated by dots.")
    }
    return IPAddress{}, errors.New(value)
}
