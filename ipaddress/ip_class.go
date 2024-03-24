// Package ipaddress used to store the different project models related
// to ip addresses
package ipaddress

// IPClass represents the 5 IPv4 classes detailed in public/private including
// including special ranges
type IPClass uint8

const (
    // A represents 1.0.0.0 to 127.0.0.0
    A IPClass = iota
    // B represents 128.0.0.0 to 191.255.0.0
    B
    // C represents 192.0.0.0 to 223.255.255.0
    C
    // D represents 224.0.0.0 to 239.255.255.255
    D
    // E represents 240.0.0.0 to 255.255.255.255
    E
    // Other represents special ranges like 0.0.0.0 to 1.0.0.0
    // and 127.0.0.0 to 127.255.255.255
    Other
)

func (i IPClass) String() string {
    switch i {
        case A:
            return "A"
        case B:
            return "B"
        case C:
            return "C"
        case D:
            return "D"
        case E:
            return "E"
    }
    return "Other"
}

