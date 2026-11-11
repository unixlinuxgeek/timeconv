package timeconv

// SecToMs Second to Millisecond
func SecToMs(s Second) Millisecond {
	return Millisecond(s * 1e+3)
}

// SecToMicroSec Second to MicroSecond
func SecToMicroSec(s Second) Microsecond {
	return Microsecond(s * 1e+6)
}

// SecToNanoSec Second to NanoSecond
func SecToNanoSec(s Second) Nanosecond {
	return Nanosecond(s * 1e+9)
}
