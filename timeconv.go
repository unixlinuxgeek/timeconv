package timeconv

import (
	"fmt"
)

type Millisecond float64
type Second float64

func (ms Millisecond) String() string {
	return fmt.Sprintf("%.3f ms", ms)
}

func (s Second) String() string {
	return fmt.Sprintf("%.3f sec", s)
}
