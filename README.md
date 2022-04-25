# escapes

Package escapes has a handful of useful functions for when you're
wanting to send formatting escape sequences to the terminal while
tracking the width of your output for alignment purposes.

The escape sequences always restore things back to 'normal' when
done, which could be considered suboptimal when adding styles but
keeps things sane.

This does not use terminfo, and just relies on ECMA-48 and related
ANSI standards; see https://en.wikipedia.org/wiki/ANSI_escape_code
for too much and yet too little detail.
