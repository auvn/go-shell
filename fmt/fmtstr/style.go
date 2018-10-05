package fmtstr

import (
	"fmt"
	"io"
)

type Style string

const (
	StyleBold   Style = "\033[1m"
	StyleNormal Style = "\033[0m"
)

var (
	DefaultStyle = StyleNormal
)

func style(st Style, s string) string {
	return string(st) + s + string(DefaultStyle)
}

func Bold(s string) string {
	return style(StyleBold, s)
}

func Normal(s string) string {
	return style(StyleNormal, s)
}

func Fprintf(w io.Writer, style Style, format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, string(style)+format, args...)
}
