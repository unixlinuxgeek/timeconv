package timeconv

import (
	"fmt"
)

type Second float64
type Millisecond float64
type Microsecond float64
type Nanosecond float64

func (s Second) String() string {
	return fmt.Sprintf("%.3f s\n", s)
}
func (mls Millisecond) String() string {
	return fmt.Sprintf("%.3f ms\n", mls)
}
func (us Microsecond) String() string {
	return fmt.Sprintf("%.3f us\n", us)
}
func (ns Nanosecond) String() string {
	return fmt.Sprintf("%.3f ns\n", ns)
}
