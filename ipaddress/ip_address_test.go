package ipaddress

import "testing"

func getIPAddressSample() IPAddress {
    return CreateIPAddress(1, 2, 3, 4)
}

func TestCreateIPAddressWith1234ReturnSuccess(t *testing.T) {
    getIPAddressSample()
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

func TestParseIPAddressWith1ReturnError(t *testing.T) {
    _, err := ParseIPAddress("1")
    expectedErrMsg := "An IP Address element must have 4 parts separated by dots."
    if err != nil && err.Error() != expectedErrMsg {
        t.Fail()
    }
}
