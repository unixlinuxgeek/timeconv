package time

import (
	"github.com/unixlinuxgeek/floatgen"
	"testing"
)

func Test_SecToMin(t *testing.T) {
	r := floatgen.GenRan(1, 9)
	s1 := SecToMin(Sec(r))
	s2 := Min(r / 60)

	if s1 == s2 {
		t.Logf("Test Passed: %f == %f", s1, s2)
	} else {
		t.Errorf("Error: %f Not Equal %f", s1, s2)
	}
}

func Test_MinToSec(t *testing.T) {
	r := floatgen.GenRan(1, 9)
	s1 := MinToSec(Min(r))
	s2 := Sec(r * 60)

	if s1 == s2 {
		t.Logf("Test Passed: %f == %f", s1, s2)
	} else {
		t.Error("Test Failed!!!")
	}
}

// Need Check!!!
func Test_HourToMin(t *testing.T) {
	r := floatgen.GenRan(1, 9)
	s1 := HourToMin(Hour(r))
	s2 := Min(r * 60)

	if s1 == s2 {
		t.Log("Test Passed")
	} else {
		t.Error("Test Failed!!!")
	}
}

// 1 hour = 3600 seconds
func Test_HourToSec(t *testing.T) {
	r := floatgen.GenRan(1, 9)
	s1 := HourToSec(Hour(r))
	s2 := Sec(r * (60 * 60))
	t.Logf("r: %f", r)
	if s1 == s2 {
		t.Log("Test Passed")
	} else {
		t.Error("Test Failed!!!")
	}
}
