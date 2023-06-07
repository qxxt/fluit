package fluit

import (
	"fmt"
)

// Arguments's padding length
// Argument are formated to be surrounded by padding:
//
// [padding][Argument][padding]
var ArgumentPaddingLength int = 2

// Default argument length
const DefaultArgumentRowLength int = 20

type Option struct {
	// Text displayed on argument row.
	// These following list are valid:
	// "-l"
	// "--version"
	// "-s, --size"
	// "-w, --width=COLS"
	Argument string
	// Text displayed on description row
	Description string
}

type Usages struct {
	UsageGroup []Option

	// SetArgumentColumnLength() sets a constant width to argument collumn that is not
	// relative to the largest args length.
	//
	// This is useful --if-you-have-a-longer-than-usual-flag. And don't want
	// it to affect the argument collumn size.
	//
	// By default, if the argument len is larger than the specified
	// l, it will have its own line.
	ArgumentRowLength int
}

// AddUsage() adds usage to *u.
func (u *Usages) AddOption(argument, description string) {
	u.UsageGroup = append(u.UsageGroup, Option{argument, description})
}

// PrintUsages() prints all the added usages to console.
func (u *Usages) PrintUsages() {
	argumentRowLength := DefaultArgumentRowLength

	if u.ArgumentRowLength > 0 {
		argumentRowLength = u.ArgumentRowLength
	}

	for _, usage := range u.UsageGroup {
		fmt.Print(SprintUsage(argumentRowLength, usage.Argument, usage.Description))
	}
}

// SprintUsage() build a usage and return it as string. If the length of
// argument is larger than maxArgumentLength, argument will have its own line.
// Intended for testing only.
//
// For building usages en mass. You should use type and method Usages instead.
func SprintUsage(maxArgumentLength int, argument, description string) string {
	if maxArgumentLength < 10 || maxArgumentLength >= DefaultBreakpoint {
		maxArgumentLength = 0
	}

	argumentRowLength := maxArgumentLength + (ArgumentPaddingLength * 2)
	formatedDescription := SprintWrap(argumentRowLength, description)

	if len(argument) > maxArgumentLength {
		return SprintWrap(ArgumentPaddingLength, argument) + "\n" + formatedDescription + "\n"
	}

	formatedArgument := fmt.Sprintf("%*s%-*s", ArgumentPaddingLength, "", maxArgumentLength+ArgumentPaddingLength, argument)

	return formatedArgument + formatedDescription[argumentRowLength:] + "\n"
}
