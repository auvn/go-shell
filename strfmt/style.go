package strfmt

import (
	"fmt"
	"io"
)

type Style string

const (
	esc Style = "\033[0m"

	StyleBold Style = "\033[1m"
)

var (
	DefaultStyle = esc
)

func style(st Style, s string) string {
	return string(st) + s + string(DefaultStyle)
}

func Bold(s string) string {
	return style(StyleBold, s)
}

func Normal(s string) string {
	return style(esc, s)
}

func Fprintf(w io.Writer, st Style, format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, style(st, format), args...)
}
