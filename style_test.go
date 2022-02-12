package marker_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/nolwn/marker"
)

func TestColor(t *testing.T) {
	style := marker.Style().Color(marker.Magenta)
	value := fmt.Sprint(style)

	if value != "\x1b[35m" {
		t.Fatalf("expected magenta escape code, but received %s", value[1:])
	}
}

func TestBrightColor(t *testing.T) {
	style := marker.Style().Color(marker.BrtBlue)
	value := fmt.Sprint(style)

	if value != "\x1b[94m" {
		t.Fatalf("expected bright blue escape code, but received %s", value[1:])
	}
}

func TestBackgroundColor(t *testing.T) {
	style := marker.Style().Background(marker.White)
	value := fmt.Sprint(style)

	if value != "\x1b[47m" {
		t.Fatalf("expected white background escape code, but received %s", value[1:])
	}
}

func TestBrightBackgroundColor(t *testing.T) {
	style := marker.Style().Background(marker.BrtCyan)
	value := fmt.Sprint(style)

	if value != "\x1b[106m" {
		t.Fatalf(
			"expected bright cyan background escape code, but received %s",
			value[1:],
		)
	}
}

func TestBackgroundAndForeground(t *testing.T) {
	style := marker.Style().Background(marker.Magenta).Color(marker.Blue)
	value := fmt.Sprint(style)

	values, err := parseCodeValues(value)

	if err != nil {
		t.Fatalf("expected to be able to parse codes, but received error: %s", err)
	}

	if len(values) != 2 {
		t.Fatalf("expected to have two codes but received %d", len(values))
	}

	for _, v := range []int8{45, 34} {
		if !contains(values, v) {
			t.Fatalf("expected to be able to parse codes, but received error: %s", value[1:])
		}
	}
}

func TestEffect(t *testing.T) {
	style := marker.Style().Effect(marker.Dim)
	value := fmt.Sprint(style)

	if value != "\x1b[2m" {
		t.Fatalf("expected dim escape code, but received %s", value[1:])
	}
}

func TestEffectColorAndBackground(t *testing.T) {
	style := marker.Style().Effect(
		marker.Underline,
	).Background(
		marker.Red,
	).Color(
		marker.BrtGreen,
	)

	value := fmt.Sprint(style)

	values, err := parseCodeValues(value)

	if err != nil {
		t.Fatalf("expected to be able to parse codes, but received error: %s", value[1:])
	}

	if len(values) != 3 {
		t.Fatalf("expected to have three codes but received %d", len(values))
	}

	for _, v := range []int8{41, 4, 92} {
		if !contains(values, v) {
			t.Fatalf("expected %s but got these codes: %v", value[1:], values)
		}
	}
}

func TestBackgroundAndEffect(t *testing.T) {
	style := marker.Style().Background(marker.BrtWhite).Effect(marker.Blinking)
	value := fmt.Sprint(style)

	values, err := parseCodeValues(value)

	if err != nil {
		t.Fatalf("expected to be able to parse codes, but received error: %s", value[1:])
	}

	if len(values) != 2 {
		t.Fatalf("expected to have two codes but received: %d", len(values))
	}

	for _, v := range []int8{107, 5} {
		if !contains(values, v) {
			t.Fatalf("expected %s but got these codes: %v", value[1:], values)
		}
	}
}

func TestWrite(t *testing.T) {
	str := marker.Style().Color(marker.Blue).Write("This is some text")
	expected := fmt.Sprintf("%sThis is some text%s", marker.Style().Color(marker.Blue), marker.Reset())

	if str != expected {
		t.Fatalf("expected %s but received: %s", expected, str)
	}
}

func contains(values []int8, value int8) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}

func parseCodeValues(code string) (values []int8, err error) {
	var value int
	codeString := code[2 : len(code)-1]
	codes := strings.Split(codeString, ";")

	for _, c := range codes {
		value, err = strconv.Atoi(string(c))

		if err != nil {
			return nil, err
		}

		values = append(values, int8(value))
	}

	return values, nil
}
