package timeconv

// SecToMs second to millisecond
func SecToMs(s Second) Millisecond {
	ms := s * 1000
	return Millisecond(ms)
}

func MsToSec(ms Millisecond) Second {
	return Second(ms / 1000)
}

func SecToMicroSec(s Second) MicroSec {
	return MicroSec(s * 1000000)
}

func MicroSecToSec(m MicroSec) Second {
	return Second(m / 1000000)
}
