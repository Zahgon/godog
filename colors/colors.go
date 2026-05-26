package colors

const ansiEscape = "\x1b"

// a color code type
type color int

// some ansi colors
const (
	black color = iota + 30
	red
	green
	yellow
	blue    // unused
	magenta // unused
	cyan
	white
)

func colorize(s interface{}, c color) string { _ = "STUB: not implemented"; return "" }

// ColorFunc is a helper type to create colorized strings.
type ColorFunc func(interface{}) string

// Bold will accept a ColorFunc and return a new ColorFunc
// that will make the string bold.
func Bold(fn ColorFunc) ColorFunc { _ = "STUB: not implemented"; return *new(ColorFunc) }

// Green will accept an interface and return a colorized green string.
func Green(s interface{}) string { _ = "STUB: not implemented"; return "" }

// Red will accept an interface and return a colorized red string.
func Red(s interface{}) string { _ = "STUB: not implemented"; return "" }

// Cyan will accept an interface and return a colorized cyan string.
func Cyan(s interface{}) string { _ = "STUB: not implemented"; return "" }

// Black will accept an interface and return a colorized black string.
func Black(s interface{}) string { _ = "STUB: not implemented"; return "" }

// Yellow will accept an interface and return a colorized yellow string.
func Yellow(s interface{}) string { _ = "STUB: not implemented"; return "" }

// White will accept an interface and return a colorized white string.
func White(s interface{}) string { _ = "STUB: not implemented"; return "" }
