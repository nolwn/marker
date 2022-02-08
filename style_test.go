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
		t.Fatalf("expected bright cyan background escape code, but received %s", value[1:])
	}
}

func TestBackgroundAndForeground(t *testing.T) {
	style := marker.Style().Background(marker.Magenta).Color(marker.Blue)
	value := fmt.Sprint(style)

	values, err := parseCodeValues(value)

	if err != nil {
		t.Fatalf("expected to be able to parse codes, but received error: %s", err)
	}

	for _, v := range []int8{45, 34} {
		contains(values, v)
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
	}

	values = append(values, int8(value))

	return values, nil
}
