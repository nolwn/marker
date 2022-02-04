package codes

import (
	"fmt"
)

const esc = "\x1b"    // escape
const csi = esc + "[" // control sequence introducer
// const dcs = "\x90" // device control string
// const osc = "\x9d" // operating system command

// Colors
const reset = 0
const def = 39
const (
	black = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

type Code struct {
	options []int8
	err     string
}

func (c Code) Color(color string) Code {
	switch color {
	case "black":
		c.options = append(c.options, black)
	case "red":
		c.options = append(c.options, red)
	case "green":
		c.options = append(c.options, green)
	case "yellow":
		c.options = append(c.options, yellow)
	case "blue":
		c.options = append(c.options, blue)
	case "magenta":
		c.options = append(c.options, magenta)
	case "cyan":
		c.options = append(c.options, cyan)
	case "white":
		c.options = append(c.options, white)
	case "reset":
		c.options = append(c.options, reset)
	case "default":
		c.options = append(c.options, def)
	default:
		c.err = "invalid value provided"
	}

	return c
}

func (c Code) Value() string {
	optString := ""
	for i, opt := range c.options {
		if i > 0 {
			optString += ";"
		}

		optString += fmt.Sprint(opt)
	}

	return fmt.Sprintf("%s%sm", csi, optString)
}
