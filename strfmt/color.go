package strfmt

const (
	StyleRed   Style = "\033[31m"
	StyleGreen Style = "\033[32m"
)

func color(c Style, s string) string {
	return string(c) + s + string(esc)
}

func Green(s string) string {
	return color(StyleGreen, s)
}

func Red(s string) string {
	return color(StyleRed, s)
}
