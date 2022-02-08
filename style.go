package marker

import "fmt"

const esc = "\x1b"    // escape
const csi = esc + "[" // control sequence introducer
// const dcs = "\x90" // device control string
// const osc = "\x9d" // operating system command

// Color Styles

// Reset resets all styles and colors
const Reset = 0

// Color number values. These correspond to the foreground colors, but can be made into
// background colors just by adding 10 which is done automatically by the Background
// method.
const (
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	_
	Default // Default resets colors
)

// Bright color number values. These correspond to the foreground colors, but can be
// made into background colors just by adding 10 which is done automatically by the
// Background method.
const (
	BrtBlack = iota + Black + 60
	BrtRed
	BrtGreen
	BrtYellow
	BrtBlue
	BrtMagenta
	BrtCyan
	BrtWhite
)

type style struct {
	color      int8
	background int8
}

// Style returns a style type that can be used to generate ansi escape codes. The style
// type implements fmt.Stringer, which will returns the ansi code that was constructed.
func Style() *style {
	return &style{}
}

// Color take a color value and sets it as the color of the text. The color value is
// expected to correspond to the actual escape code value for that color and, since
// no one wants to have to remember that, the programmer is expected to use one of the
// provided const values.
//
// If a value falls outside of the expected range, nothing will be set and the passed
// value will be ignored.
func (s *style) Color(c int8) *style {
	if c >= Black && c <= White {
		s.color = c
	} else if c >= BrtBlack && c <= BrtWhite {
		s.color = c
	} else if c == Default {
		s.color = c
	}

	return s
}

// Background take a color value and sets it as the color of the text. The color value is
// expected to correspond to the actual escape code value for that color as a foreground
// color. Background will convert that value to a background color automatically. Since
// no one wants to have to remember those values, the programmer is expected to use one
// of the provided const values.
//
// If a value falls outside of the expected range, nothing will be set and the passed
// value will be ignored.
func (s *style) Background(c int8) *style {
	if c >= Black && c <= White {
		s.background = c + 10
	} else if c >= BrtBlack && c <= BrtWhite {
		s.background = c + 10
	} else if s.color == Default {
		s.background = c + 10
	}

	return s
}

// String satisfies the fmt.Stringer method and allows the style type to be passed as a
// string to fmt function. The resulting string will be the escape code value that was
// generated by previous method calls.
func (c *style) String() string {
	codes := ""

	if c.background != 0 {
		codes = fmt.Sprint(c.background)
	}

	if c.color != 0 {
		if len(codes) > 0 {
			codes += ";"
		}

		codes += fmt.Sprint(c.color)
	}

	return fmt.Sprintf("%s%sm", csi, codes)
}