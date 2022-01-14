package fluit

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	defaultBP int = 60
)

var (
	breakpoint = defaultBP
)

type usage struct {
	argu, descr string
}

type Usages struct {
	usageItem  []usage
	maxArgLen  int
	longArgLen int
}

// Sprintwrap() wraps s with breakpoint as specified using
// SetBreakpoint(). If no breakpoint is set, it will use the
// default value of 60.
// It also add margin with the size of marginLen. If marginLen is set to
// 0, it will only be wrapped. The marginLen will be set to 0 if
// it's negative or if it's larger than the breakpoint length.
// It return a string without printing.
func Sprintwrap(marginLen int, s ...interface{}) string {
	sstr := fmt.Sprint(s...)
	if marginLen < 0 || marginLen >= breakpoint {
		marginLen = 0
	}
	margin := strings.Repeat(" ", marginLen)
	if breakpoint > len(sstr)+marginLen {
		return margin + sstr + "\n"
	}
	var ptn string
	lineBstr := strconv.Itoa(breakpoint - marginLen)
	// if there is less [:space:] than there is possible line. It will be
	// treated as a hash like.
	if !strings.Contains(sstr, " ") ||
		strings.Count(sstr, " ") < len(sstr)/(breakpoint-marginLen) {
		ptn = `(.{1,` + lineBstr + `})`
	} else {
		ptn = `(.{1,` + lineBstr + `})(?:\s|$)`
	}
	return regexp.MustCompile(ptn).
		ReplaceAllString(sstr, margin+"$1\n")
}

// SprintUsg() build a usage and return it as string. If the length of
// arg is larger than maxArgLen, arg will have its own line. For building usages
// en mass. You should use type Usages instead.
func SprintUsg(arg, desc string, maxArgLen int) string {
	const argPadLen int = 2 // argument's padding length for both side
	var (
		fmtDesc     string // Text wrapped Arg's Description
		fmtArg      string // Formatted Arg. Aka, padded string
		argCol      int    // Arg collumn length. Its maxLen + both side padding length
		argLeftPad  string // self explanatory
		argRightPad string // Maximum argument's length - its actual length + padding length.
	)
	if maxArgLen < 0 || maxArgLen >= breakpoint {
		maxArgLen = 0
	}
	argCol = maxArgLen + (argPadLen * 2)
	fmtDesc = Sprintwrap(argCol, desc)
	argLength := len(arg)
	if argLength > maxArgLen {
		fmtArg = Sprintwrap(argPadLen, arg)
		return fmtArg + fmtDesc
	}
	argLeftPad = strings.Repeat(" ", argPadLen)
	argRightPad = strings.Repeat(" ", maxArgLen-argLength+argPadLen)
	fmtArg = argLeftPad + arg + argRightPad
	return strings.Replace(
		fmtDesc,
		strings.Repeat(" ", argCol),
		fmtArg, 1)
}

// SetBreakpoint sets at what length the texts.should
// break a newline. Console with the column length
// of 70 may only use 70 breakpoint or less. If breakpoint is not
// set, it will use the default value of 60.
//
// NOTE: You can get the width of the console using
// /x/term.GetScreenwidth(). That allows you to create a
// breakpoint which responsive to user's screen width. Sounds great.
// However not all console are supported, eg emacs mini shell XD.
// It may return 0 or -1 value. In that case. The value will be
// Ignored and default to 60.
func SetBreakpoint(bp int) {
	if bp > 0 {
		breakpoint = bp
	}
}

// AddUsg() method adds usages which can later be printed using PrintUsg()
// If SetArglen() is not specified. It will compare the len() of arg
// and replace the current longest with it if larger.
func (u *Usages) AddUsg(arg, desc string) {
	if u.maxArgLen == 0 && len(arg) > u.longArgLen {
		u.longArgLen = len(arg)
	}
	u.usageItem = append(u.usageItem, usage{arg, desc})
}

// SetArgLen() adds a constant width to argument's collumn that is not.
// relative to the longest argument length.
//
// This is useful --if-you-have-a-longer-than-usual-flag. And don't want
// it to affect the argument collumn size.
// By default, if the argument size is longer than the specified
// maximum argument length, it will have its own line.
func (u *Usages) SetArgLen(l int) {
	if l > 0 {
		u.maxArgLen = l
	}
}

// PrintUsg() prints all the added usages to terminal.
// Upon calling this method the object is reseted to nil, and can be used // again.
// breakpoint value from SetBreakpoint() is not affected.
func (u *Usages) PrintUsg() {
	var b int
	if u.maxArgLen != 0 {
		b = u.maxArgLen
	} else {
		b = u.longArgLen
	}
	for _, usage := range u.usageItem {
		fmt.Print(SprintUsg(usage.argu, usage.descr, b))
	}
	*u = Usages{}
}
