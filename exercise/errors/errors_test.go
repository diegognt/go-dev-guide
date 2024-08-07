package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time       string
		parsedTime Time
		ok         bool
	}{
		{"19:00:12", Time{19, 0, 12}, true},
		{"1:3:44", Time{1, 3, 44}, true},
		{"bad", Time{}, false},
		{"1:-3:44", Time{}, false},
		{"0:59:59", Time{0, 59, 59}, true},
		{"", Time{}, false},
		{"11:22", Time{}, false},
		{"aa:bb:cc", Time{}, false},
		{"5:23:", Time{}, false},
	}

	for _, data := range table {
		result, err := ParseTime(data.time)

		if data.ok && data.parsedTime != result {
			t.Errorf("%v result sould be %v", data.parsedTime, result)
		}

		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}
