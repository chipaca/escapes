package escapes_test

import (
	"image/color"
	"testing"

	"github.com/chipaca/escapes"
)

var red color.Color = color.RGBA{R: 255}

func TestColorized(t *testing.T) {
	type TC struct {
		fg *color.Color
		bg *color.Color
		s  string
		w  int
	}

	for name, tc := range map[string]TC{
		"none": {
			fg: nil,
			bg: nil,
			s:  "foo",
			w:  3,
		},
		"redfg": {
			fg: &red,
			bg: nil,
			s:  "\x1b[38;2;255;0;0mfoo\x1b[0m",
			w:  3,
		},
		"redbg": {
			fg: nil,
			bg: &red,
			s:  "\x1b[48;2;255;0;0mfoo\x1b[0m",
			w:  3,
		},
		"red2": {
			fg: &red,
			bg: &red,
			s:  "\x1b[38;2;255;0;0;48;2;255;0;0mfoo\x1b[0m",
			w:  3,
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := escapes.Colorized("foo", tc.fg, tc.bg)
			if s := got.String(); s != tc.s {
				t.Errorf("got string %q, expected %q", s, tc.s)
			}
			if w := got.Width(); w != tc.w {
				t.Errorf("got width %d, expected %d", w, tc.w)
			}
		})
	}
}
