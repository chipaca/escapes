// Package escapes has a handful of useful functions for when you're
// wanting to send formatting escape sequences to the terminal while
// tracking the width of your output for alignment purposes.
//
// The escape sequences always restore things back to 'normal' when
// done, which could be considered suboptimal when adding styles but
// keeps things sane.
//
// This does not use terminfo, and just relies on ECMA-48 and related
// ANSI standards; see https://en.wikipedia.org/wiki/ANSI_escape_code
// for too much and yet too little detail.
package escapes

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/chipaca/width"
	"github.com/mattn/go-runewidth"
)

// Hyperlink returns an escape sequence that links to the url u with the text t.
func Hyperlink(u, t string) width.StringWidther {
	return width.StringAndWidth{
		S: fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", u, t),
		W: runewidth.StringWidth(t),
	}
}

// Colorized returns a string with escape sequences that shows the
// text t with foreground and backgroun colours as given.
//
// It assumes 24-bit colour support.
func Colorized(s string, fg, bg *color.Color) width.StringWidther {
	if fg == nil && bg == nil {
		return width.String(s)
	}
	var bits []string
	if fg != nil {
		fgR, fgG, fgB, _ := (*fg).RGBA()
		bits = append(bits,
			"38",
			"2",
			strconv.FormatUint(uint64(fgR>>8), 10),
			strconv.FormatUint(uint64(fgG>>8), 10),
			strconv.FormatUint(uint64(fgB>>8), 10),
		)
	}
	if bg != nil {
		bgR, bgG, bgB, _ := (*bg).RGBA()
		bits = append(bits,
			"48",
			"2",
			strconv.FormatUint(uint64(bgR>>8), 10),
			strconv.FormatUint(uint64(bgG>>8), 10),
			strconv.FormatUint(uint64(bgB>>8), 10),
		)
	}
	return width.StringAndWidth{
		S: CSI + strings.Join(bits, ";") + "m" + s + SGR0,
		W: runewidth.StringWidth(s),
	}
}

// CSI or Control Sequence Introducer is the start of most things we deal with here
const CSI = "\x1b["

// Several constants that might be handy to have (and are used by us)
const (
	BeginBold      = CSI + "1m"
	BeginDim       = CSI + "2m"
	BeginItalic    = CSI + "3m"
	BeginUnderline = CSI + "4m"
	BeginReverse   = CSI + "7m"
	SGR0           = CSI + "0m"
	ClrEOL         = CSI + "K"
)

// Reverse returns a string that turns on reverse video mode for the
// text given.
func Reverse(s string) width.StringWidther {
	return width.StringAndWidth{
		S: BeginReverse + s + SGR0,
		W: runewidth.StringWidth(s),
	}
}

// Bold returns a string that turns on bold mode for the text given.
func Bold(s string) width.StringWidther {
	return width.StringAndWidth{
		S: BeginBold + s + SGR0,
		W: runewidth.StringWidth(s),
	}
}

// Dim returns a string that turns on dim mode for the text given.
func Dim(s string) width.StringWidther {
	return width.StringAndWidth{
		S: BeginDim + s + SGR0,
		W: runewidth.StringWidth(s),
	}
}

// Italic returns a string that turns on italic mode for the text given.
//
// NOTE many terminals don't implement this (or have support for it
//      turned off by default).
func Italic(s string) width.StringWidther {
	return width.StringAndWidth{
		S: BeginItalic + s + SGR0,
		W: runewidth.StringWidth(s),
	}
}

// Underline returns a string that turns on underline mode for the text given.
func Underline(s string) width.StringWidther {
	return width.StringAndWidth{
		S: BeginUnderline + s + SGR0,
		W: runewidth.StringWidth(s),
	}
}

// ReverseLine returns a string that uses reverse video and clears to the end of the line.
//
// This includes a newline to avoid confusion (otherwise the cursor
// maybe be in an unexpected place).
//
// Note that the width of the string is terminal-dependent and thus not knowable from here.
func ReverseLine(s string) string {
	return BeginReverse + s + ClrEOL + SGR0 + "\n"
}
