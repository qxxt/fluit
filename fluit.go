package fluit

import (
	"regexp"
	"strconv"
	"strings"
)

func Sprintfm(input string, marginLength int, breakpoint int) (res string) {
	if marginLength < 0 || marginLength >= breakpoint {
		marginLength = 0
	}
	margin := strings.Repeat(" ", marginLength)
	if breakpoint > len(input)+marginLength {
		return margin + input + "\r\n"
	}
	lineBreakLimit := strconv.Itoa(breakpoint - marginLength)
	if !strings.Contains(input, " ") {
		regexPtn := regexp.MustCompile("(.{1," + lineBreakLimit + "})")
		return regexPtn.ReplaceAllString(input, margin+"$1\r\n")
	}
	regexPtn := regexp.MustCompile("(.{1," + lineBreakLimit + "})(?:\\s|$)")
	return regexPtn.ReplaceAllString(input, margin+"$1\r\n")
}

func SprintfUsg(arg string, desc string, maxArgLength int, breakpoint int) (res string) {
	const argPaddingLength int = 2
	if maxArgLength < 0 || maxArgLength >= breakpoint {
		maxArgLength = 0
	}
	arg_collumn_length := maxArgLength + (argPaddingLength * 2)
	formatted_desc := Sprintfm(desc, arg_collumn_length, breakpoint)
	argLength := len(arg)
	if argLength > maxArgLength {
		formatted_arg := Sprintfm(arg, argPaddingLength, breakpoint)
		return formatted_arg + formatted_desc
	}
	arg = strings.Repeat(" ", argPaddingLength) + arg + strings.Repeat(" ", maxArgLength-argLength+argPaddingLength)
	res = strings.Replace(formatted_desc, strings.Repeat(" ", arg_collumn_length), arg, 1)
	return res
}
