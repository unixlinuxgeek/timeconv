package timeconv

import (
	"math/rand/v2"
	"testing"
)

// TestSecToMs second to millisecond
func TestSecToMs(t *testing.T) {
	r := rand.Float64()
	x := Millisecond(r * 1000)

	if SecToMs(Second(r)) != x {
		t.Fatalf("SecToMs(%f) not equal %f\n", r, x)
	}
}

// TestSecToMs millisecond to second
func TestMsToSec(t *testing.T) {
	r := rand.Float64()
	x := Second(r / 1000)

	if MsToSec(Millisecond(r)) != x {
		t.Fatalf("MsToSec(%f) not equal %f\n", r, x)
	}
}
