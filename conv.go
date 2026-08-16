package timeconv

// SecToMs second to millisecond
func SecToMs(s Second) Millisecond {
	ms := s * 1000
	return Millisecond(ms)
}

func MsToSec(ms Millisecond) Second {
	return Second(ms / 1000)
}
