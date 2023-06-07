package fluit

import (
	"fmt"
)

// This is the default breakpoint.
//
// When UserBreakpoint is <= 0. This value will be used instead.
const DefaultBreakpoint int = 60

// UserBreakpoint sets at what length the inputs should break a
// newline. Console with the column length of 70 may only use
// 70 breakpoint or less. If breakpoint is not set, it will
// use the default value from DefaultBreakpoint (60).
//
// You can get the length of the console using from:
// https://pkg.go.dev/golang.org/x/term#GetSize.
//
// This allows you to create a breakpoint that is
// responsive to user's console length. However not all console
// are supported by it.
//
// For example, emacs mini shell is not a terminal and don't have length,
// it will return loopback value of 0 or -1 value.
var UserBreakpoint = -1

// SprintfWrap() formats string using fmt.Sprintf and then wraps them with
// breakpoint from UserBreakpoint. If UserBreakpoint is not set, it will use
// the default value of 60 from constant Breakpoint.
//
// It also add margin with the size of marginLength. If marginLength is set to
// 0, it will only be wrapped. The marginLength will be set to 0 if
// it's negative or if it's larger than the breakpoint's length.
//
// It returns string without print them to stdout.
func SprintWrap(marginLength int, s string) string {
	bp := DefaultBreakpoint
	if UserBreakpoint > 0 {
		bp = UserBreakpoint
	}

	if marginLength > bp {
		marginLength = 0
	}

	marginString := fmt.Sprintf("%*s", marginLength, "")

	bp -= marginLength

	const isLetter = 1

	var (
		statusBit   int
		columnIndex int
		spaceIndex  = -1
	)

	s = marginString + s
	for i := marginLength; i < len(s); i++ {
		if columnIndex == bp {
			columnIndex = 0

			// SpaceIndex is -1 when there is no space found in the entire column.
			// Mostly happen when the input is a long hash text.
			if spaceIndex == -1 {
				s = s[:i] + "\n" + marginString + s[i:]
				i += marginLength
			} else {
				s = s[:spaceIndex] + "\n" + marginString + s[spaceIndex+1:]
				i = spaceIndex + marginLength
			}

			statusBit = 0
			spaceIndex = -1
			columnIndex = 0
			continue
		}

		switch s[i] {
		case ' ', '\t':

			// Only set spaceIndex when there is letter before it.
			// to allow leading space.
			if statusBit&isLetter != 0 {
				spaceIndex = i
			}

		case '\n':
			s = s[:i] + "\n" + marginString + s[i+1:]
			i += marginLength
			statusBit = 0
			spaceIndex = -1
			columnIndex = 0
			continue

		default:
			if statusBit&isLetter == 0 {
				statusBit |= isLetter
			}
		}

		columnIndex++
	}

	return s
}

// Format string with SprintWrap and print them.
func PrintfWrap(marginLength int, format string, a ...interface{}) (int, error) {
	return fmt.Print(SprintWrap(marginLength, fmt.Sprintf(format, a...)))
}

// Format string with SprintWrap and print them with newline.
func PrintlnWrap(marginLength int, s string) (int, error) {
	return fmt.Println(SprintWrap(marginLength, s))
}
