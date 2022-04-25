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
		S: "\x1b[" + strings.Join(bits, ";") + "m" + s + "\x1b[0m",
		W: runewidth.StringWidth(s),
	}
}

// Reverse returns a string that turns on reverse video mode for the
// text given.
func Reverse(s string) width.StringWidther {
	return width.StringAndWidth{
		S: "\x1b[7m" + s + "\x1b[0m",
		W: runewidth.StringWidth(s),
	}
}

// Bold returns a string that turns on bold mode for the text given.
func Bold(s string) width.StringWidther {
	return width.StringAndWidth{
		S: "\x1b[1m" + s + "\x1b[0m",
		W: runewidth.StringWidth(s),
	}
}

// Italic returns a string that turns on italic mode for the text given.
//
// NOTE many terminals don't implement this (or have support for it
//      turned off by default).
func Italic(s string) width.StringWidther {
	return width.StringAndWidth{
		S: "\x1b[3m" + s + "\x1b[0m",
		W: runewidth.StringWidth(s),
	}
}
