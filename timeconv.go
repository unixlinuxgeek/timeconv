package time

import "fmt"

type Sec float64
type Min float64
type Hour float64

func (s Sec) String() string {
	return fmt.Sprintf("%g Сек.", s)
}
func (m Min) String() string {
	return fmt.Sprintf("%g Мин.", m)
}
func (h Hour) String() string {
	return fmt.Sprintf("%g ч.", h)
}

func SecToMin(s Sec) Min {
	return Min(s / 60)
}

func MinToSec(m Min) Sec {
	return Sec(m * 60)
}

func HourToMin(h Hour) Min {
	return Min(h * 60)
}

func HourToSec(h Hour) Sec {
	return Sec(h * (60 * 60))
}
