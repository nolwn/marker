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

type colorCode struct {
	bright     bool
	background bool
	color      int8
}

func Color() *colorCode {
	return &colorCode{}
}

func (c *colorCode) Black() *colorCode {
	c.color = black

	return c
}

func (c *colorCode) Red() *colorCode {
	c.color = red

	return c
}

func (c *colorCode) Green() *colorCode {
	c.color = green

	return c
}

func (c *colorCode) Yellow() *colorCode {
	c.color = yellow

	return c
}

func (c *colorCode) Blue() *colorCode {
	c.color = blue

	return c
}

func (c *colorCode) Magenta() *colorCode {
	c.color = magenta

	return c
}

func (c *colorCode) Cyan() *colorCode {
	c.color = cyan

	return c
}

func (c *colorCode) White() *colorCode {
	c.color = white

	return c
}

func (c *colorCode) Default() *colorCode {
	c.color = def

	return c
}

func (c *colorCode) Reset() *colorCode {
	c.color = reset

	return c
}

func (c *colorCode) Background() *colorCode {
	c.background = true

	return c
}

func (c *colorCode) Bright() *colorCode {
	c.bright = true

	return c
}

func (c *colorCode) Value() string {
	color := c.color

	if c.bright {
		color += 60
	}

	if c.background {
		color += 10
	}

	return fmt.Sprintf("%s%dm", csi, color)
}
