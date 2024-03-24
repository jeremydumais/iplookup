// Package ipaddress used to store the different project models related
// to ip addresses
package ipaddress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getIPAddressSample() IPAddress {
    return CreateIPAddress(1, 2, 3, 4)
}

func TestCreateIPAddressWith1234ReturnSuccess(t *testing.T) {
    ip := getIPAddressSample()
    if ip.GetByte1() != 1 {
        t.Fatalf("GetByte1 returned %v instead of %v", ip.GetByte1(), 1)
    }
    if ip.GetByte2() != 2 {
        t.Fatalf("GetByte2 returned %v instead of %v", ip.GetByte2(), 2)
    }
    if ip.GetByte3() != 3 {
        t.Fatalf("GetByte3 returned %v instead of %v", ip.GetByte3(), 3)
    }
    if ip.GetByte4() != 4 {
        t.Fatalf("GetByte4 returned %v instead of %v", ip.GetByte4(), 4)
    }
}

func TestGetByte1With1234Return1(t *testing.T) {
    ip := getIPAddressSample()
    actual := ip.GetByte1()
    if actual != 1 {
        t.Fatalf("GetByte1 returned %v instead of %v", actual, 1)
    }
}

func TestGetByte2With1234Return1(t *testing.T) {
    ip := getIPAddressSample()
    actual := ip.GetByte2()
    if actual != 2 {
        t.Fatalf("GetByte2 returned %v instead of %v", actual, 2)
    }
}

func TestGetByte3With1234Return1(t *testing.T) {
    ip := getIPAddressSample()
    actual := ip.GetByte3()
    if actual != 3 {
        t.Fatalf("GetByte3 returned %v instead of %v", actual, 3)
    }
}

func TestGetByte4With1234Return1(t *testing.T) {
    ip := getIPAddressSample()
    actual := ip.GetByte4()
    if actual != 4 {
        t.Fatalf("GetByte4 returned %v instead of %v", actual, 4)
    }
}

func TestStringWith1234Return1234AsString(t *testing.T) {
    ip := getIPAddressSample()
    const expected = "1.2.3.4"
    actual := ip.String()
    if actual != expected {
        t.Fatalf("String returned %v instead of %v", actual, expected)
    }
}

func TestParseIPAddressWithEmptyReturnError(t *testing.T) {
    _, err := ParseIPAddress("")
    expectedErrMsg := "an IP Address element must have 4 parts separated by dots"
    if err != nil && err.Error() != expectedErrMsg {
        t.Fatalf("Error returned is '%v' instead of '%v'", err.Error(), expectedErrMsg)
    }
}

func TestParseIPAddressWithWhiteSpacesReturnError(t *testing.T) {
    _, err := ParseIPAddress("   ")
    expectedErrMsg := "an IP Address element must have 4 parts separated by dots"
    if err != nil && err.Error() != expectedErrMsg {
        t.Fatalf("Error returned is '%v' instead of '%v'", err.Error(), expectedErrMsg)
    }
}

func TestParseIPAddressWith1ReturnError(t *testing.T) {
    _, err := ParseIPAddress("1")
    expectedErrMsg := "an IP Address element must have 4 parts separated by dots"
    if err != nil && err.Error() != expectedErrMsg {
        t.Fatalf("Error returned is '%v' instead of '%v'", err.Error(), expectedErrMsg)
    }
}

func TestParseIPAddressWitha234ReturnError(t *testing.T) {
    _, err := ParseIPAddress("a.2.3.4")
    expectedErrMsg := "each part of an IP Address must be a number between 0 and 255"
    if err != nil && err.Error() != expectedErrMsg {
        t.Fatalf("Error returned is '%v' instead of '%v'", err.Error(), expectedErrMsg)
    }
}

func TestParseIPAddressWith1a23ReturnError(t *testing.T) {
    _, err := ParseIPAddress("1.a.3.4")
    expectedErrMsg := "each part of an IP Address must be a number between 0 and 255"
    if err != nil && err.Error() != expectedErrMsg {
        t.Fatalf("Error returned is '%v' instead of '%v'", err.Error(), expectedErrMsg)
    }
}

func TestParseIPAddressWith12a4ReturnError(t *testing.T) {
    _, err := ParseIPAddress("1.2.a.4")
    expectedErrMsg := "each part of an IP Address must be a number between 0 and 255"
    if err != nil && err.Error() != expectedErrMsg {
        t.Fatalf("Error returned is '%v' instead of '%v'", err.Error(), expectedErrMsg)
    }
}
func TestParseIPAddressWith123aReturnError(t *testing.T) {
    _, err := ParseIPAddress("1.2.3.a")
    expectedErrMsg := "each part of an IP Address must be a number between 0 and 255"
    if err != nil && err.Error() != expectedErrMsg {
        t.Fatalf("Error returned is '%v' instead of '%v'", err.Error(), expectedErrMsg)
    }
}

func TestGetClassWith1000ReturnA(t *testing.T) {
    ip := CreateIPAddress(1, 0, 0, 0)
    assert.Equal(t, A, ip.GetClass(), "The IP classes should be equals")
}

func TestGetClassWith127000ReturnA(t *testing.T) {
    ip := CreateIPAddress(127, 0, 0, 0)
    assert.Equal(t, A, ip.GetClass(), "The IP classes should be equals")
}

func TestGetClassWith128000ReturnB(t *testing.T) {
    ip := CreateIPAddress(128, 0, 0, 0)
    assert.Equal(t, B, ip.GetClass(), "The IP classes should be equals")
}

func TestGetClassWith191000ReturnB(t *testing.T) {
    ip := CreateIPAddress(191, 0, 0, 0)
    assert.Equal(t, B, ip.GetClass(), "The IP classes should be equals")
}

func TestGetClassWith192000ReturnC(t *testing.T) {
    ip := CreateIPAddress(192, 0, 0, 0)
    assert.Equal(t, C, ip.GetClass(), "The IP classes should be equals")
}

func TestGetClassWith223123ReturnC(t *testing.T) {
    ip := CreateIPAddress(223, 1, 2, 3)
    assert.Equal(t, C, ip.GetClass(), "The IP classes should be equals")
}

func TestGetClassWith224123ReturnD(t *testing.T) {
    ip := CreateIPAddress(224, 1, 2, 3)
    assert.Equal(t, D, ip.GetClass(), "The IP classes should be equals")
}

func TestGetClassWith239123ReturnD(t *testing.T) {
    ip := CreateIPAddress(239, 1, 2, 3)
    assert.Equal(t, D, ip.GetClass(), "The IP classes should be equals")
}

func TestGetClassWith240123ReturnE(t *testing.T) {
    ip := CreateIPAddress(240, 1, 2, 3)
    assert.Equal(t, E, ip.GetClass(), "The IP classes should be equals")
}

func TestGetClassWith255255255255ReturnE(t *testing.T) {
    ip := CreateIPAddress(255, 255, 255, 255)
    assert.Equal(t, E, ip.GetClass(), "The IP classes should be equals")
}

func TestIsPrivateWith1234ReturnFalse(t *testing.T) {
    ip := CreateIPAddress(1, 2, 3, 4)
    assert.Equal(t, false, ip.IsPrivate(), "The IsPrivate should return false")
}

func TestIsPrivateWith10001ReturnTrue(t *testing.T) {
    ip := CreateIPAddress(10, 0, 0, 1)
    assert.Equal(t, true, ip.IsPrivate(), "The IsPrivate should return true")
}

func TestIsPrivateWith1721601ReturnTrue(t *testing.T) {
    ip := CreateIPAddress(172, 16, 0, 1)
    assert.Equal(t, true, ip.IsPrivate(), "The IsPrivate should return true")
}

func TestIsPrivateWith1723101ReturnTrue(t *testing.T) {
    ip := CreateIPAddress(172, 31, 0, 1)
    assert.Equal(t, true, ip.IsPrivate(), "The IsPrivate should return true")
}

func TestIsPrivateWith1723200ReturnFalse(t *testing.T) {
    ip := CreateIPAddress(172, 32, 0, 0)
    assert.Equal(t, false, ip.IsPrivate(), "The IsPrivate should return false")
}


func TestIsPrivateWith19216700ReturnFalse(t *testing.T) {
    ip := CreateIPAddress(192, 167, 0, 0)
    assert.Equal(t, false, ip.IsPrivate(), "The IsPrivate should return false")
}

func TestIsPrivateWith19216801ReturnTrue(t *testing.T) {
    ip := CreateIPAddress(192, 168, 0, 1)
    assert.Equal(t, true, ip.IsPrivate(), "The IsPrivate should return true")
}

func TestIsPrivateWith192168255255ReturnTrue(t *testing.T) {
    ip := CreateIPAddress(192, 168, 255, 255)
    assert.Equal(t, true, ip.IsPrivate(), "The IsPrivate should return true")
}

func TestIsPrivateWith19216900ReturnFalse(t *testing.T) {
    ip := CreateIPAddress(192, 169, 0, 0)
    assert.Equal(t, false, ip.IsPrivate(), "The IsPrivate should return false")
}

