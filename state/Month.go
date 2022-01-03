package state

type MonthState int

const (
	SOLAR = iota
	LUNAR
	LEAP
	NONLEAP
)

type Month struct {
	Name   string
	Day    int
	Status MonthState
}

func newMonth(name string, day int) *Month {
	return &Month{
		Name: name,
		Day:  day,
	}
}

func (m *Month) Solar() bool {
	if m.Day == 31 {
		return true
	}
	return false
}

func (m Month) Lunar() bool {
	if m.Day == 30 {
		return true
	}
	return false
}

func (m Month) Leap() bool {
	if m.Day == 28 {
		return true
	}
	return false
}

func (m Month) NonLeap() bool {
	if m.Day == 29 {
		return true
	}
	return false
}
