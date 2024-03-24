// Package ipaddress used to store the different project models related
// to ip addresses
package ipaddress

import (
    "errors"
    "fmt"
    "strings"
    "strconv"
)


// The IPAddress struct used by the project
type IPAddress struct {
    byte1, byte2, byte3, byte4 uint8
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

// GetClass returns the class of the IP Address (A, B, C, D or E)
func (i IPAddress) GetClass() IPClass {
    if i.byte1 >= 1 && i.byte1 <= 127 {
        return A
    }
    if i.byte1 >= 128 && i.byte1 <= 191 {
        return B
    }
    if i.byte1 >= 192 && i.byte1 <= 223 {
        return C
    }
    if i.byte1 >= 224 && i.byte1 <= 239 {
        return D
    }
    if i.byte1 >= 240 && i.byte1 <= 255 {
        return E
    }
    return Other
}

// IsPrivate returns true if the IP address is part of a private range and false
// otherwise
func (i IPAddress) IsPrivate() bool {
    if i.byte1 == 10 {
        return true
    }
    if i.byte1 == 172 && (i.byte2 >= 16 && i.byte2 <= 31) {
        return true
    }
    if i.byte1 == 192 && i.byte2 == 168 {
        return true
    }
    return false
}

func convertStringToUInt8(value string) (uint8, error) {
    num, err := strconv.ParseUint(value, 10, 8)
    if err != nil {
        return 0, err
    }
    return uint8(num), nil
}

// ParseIPAddress try to parse a string value into an IPAddress struct
func ParseIPAddress(value string) (IPAddress, error) {
    parts := strings.Split(value, ".")
    if len(parts) != 4 {
        return IPAddress{}, errors.New("an IP Address element must have 4 parts separated by dots")
    }

    var ipParts [4]uint8
    for i, part := range parts {
        num, err := convertStringToUInt8(part)
        if err != nil {
            return IPAddress{}, errors.New("each part of an IP Address must be a number between 0 and 255")
        }
        ipParts[i] = num
    }

    return IPAddress{
        byte1: ipParts[0],
        byte2: ipParts[1],
        byte3: ipParts[2],
        byte4: ipParts[3],
    }, nil
}
