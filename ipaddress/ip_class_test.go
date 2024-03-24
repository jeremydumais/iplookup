// Package ipaddress used to store the different project models related
// to ip addresses
package ipaddress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPClassStringWithAReturnA(t *testing.T) {
    assert.Equal(t, "A", A.String(), "The IP class should be equal to A")
}

func TestIPClassStringWithBReturnB(t *testing.T) {
    assert.Equal(t, "B", B.String(), "The IP class should be equal to B")
}

func TestIPClassStringWithCReturnC(t *testing.T) {
    assert.Equal(t, "C", C.String(), "The IP class should be equal to C")
}

func TestIPClassStringWithDReturnD(t *testing.T) {
    assert.Equal(t, "D", D.String(), "The IP class should be equal to D")
}

func TestIPClassStringWithEReturnE(t *testing.T) {
    assert.Equal(t, "E", E.String(), "The IP class should be equal to E")
}

func TestIPClassStringWithOtherReturnOther(t *testing.T) {
    assert.Equal(t, "Other", Other.String(), "The IP class should be equal to Other")
}
