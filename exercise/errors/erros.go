//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minutes, seconds int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	timeComponents := strings.Split(input, ":")

	if len(timeComponents) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time components", input}
	}

	hour, err := strconv.Atoi(timeComponents[0])

	if err != nil {
		return Time{}, &TimeParseError{"Error parsing the hour component", timeComponents[0]}
	}

	minutes, err := strconv.Atoi(timeComponents[1])

	if err != nil {
		return Time{}, &TimeParseError{"Error parsing the minutes component", timeComponents[1]}
	}

	seconds, err := strconv.Atoi(timeComponents[2])

	if err != nil {
		return Time{}, &TimeParseError{"Error parsing the seconds component", timeComponents[2]}
	}

	if hour < 0 || hour > 23 {
		return Time{}, &TimeParseError{
			"hour value our of range, hour < 0 or hour > 23",
			fmt.Sprintf("%v", hour),
		}
	}

	if minutes < 0 || minutes > 59 {
		return Time{}, &TimeParseError{
			"minutes value out of range, minutes < 0 or minutes > 59",
			fmt.Sprintf("%v", minutes),
		}
	}

	if seconds < 0 || seconds > 59 {
		return Time{}, &TimeParseError{
			"seconds value out of range, seconds < 0 || seconds > 59",
			fmt.Sprintf("%v", seconds),
		}
	}

	return Time{hour, minutes, seconds}, nil
}
