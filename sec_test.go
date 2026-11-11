package timeconv

import (
	"testing"

	"github.com/unixlinuxgeek/randbtw"
)

// SecToMs Second to Millisecond
func TestSecToMs(t *testing.T) {
	s := randbtw.Generate(1, 1000)
	correct := Millisecond(s * 1e+3)
	needCheck := SecToMs(Second(s))

	// -v flag is enabled
	if testing.Verbose() {
		t.Logf("correct  : %f ", correct)
		t.Logf("need_check: %f ", needCheck)
	}

	if needCheck != correct {
		t.Errorf("need_check: \"%f\" Not Equal correct: \"%f\"", needCheck, correct)
	}
}

// SecToMicroSec Second to Microsecond
func TestSecToMicroSec(t *testing.T) {
	s := randbtw.Generate(1, 1000)
	correct := Microsecond(s * 1e+6)
	needCheck := SecToMicroSec(Second(s))

	// check is -v flag is enabled
	if testing.Verbose() {
		t.Logf("correct: %f\n", correct)
		t.Logf("need_check: %f\n", needCheck)
	}

	if needCheck != correct {
		t.Errorf("need_check: \"%f\" Not Equal correct: \"%f\"", needCheck, correct)
	}
}

// SecToNanoSec
func TestSecToNanoSec(t *testing.T) {
	s := randbtw.Generate(1, 1000)
	correct := Nanosecond(s * 1e+9)
	needCheck := SecToNanoSec(Second(s))

	// check is -v flag is enabled
	if testing.Verbose() {
		t.Logf("correct: %f\n", correct)
		t.Logf("need_check: %f\n", needCheck)
	}

	if needCheck != correct {
		t.Errorf("need_check: \"%f\" Not Equal correct: \"%f\"", needCheck, correct)
	}
}
