package codes_test

import (
	"testing"

	"github.com/nolwn/marker/codes"
)

func TestColor(t *testing.T) {
	value := codes.Color().Magenta().Value()

	if value != "\x1b[35m" {
		t.Fatalf("expected magenta escape code, but received %s", value[1:])
	}
}

func TestBrightColor(t *testing.T) {
	value := codes.Color().Blue().Bright().Value()

	if value != "\x1b[94m" {
		t.Fatalf("expected bright blue escape code, but received %s", value[1:])
	}
}

func TestBackgroundColor(t *testing.T) {
	value := codes.Color().White().Background().Value()

	if value != "\x1b[47m" {
		t.Fatalf("expected white background escape code, but received %s", value[1:])
	}
}

func TestBrightBackgroundColor(t *testing.T) {
	value := codes.Color().Cyan().Bright().Background().Value()

	if value != "\x1b[106m" {
		t.Fatalf("expected bright white background escape code, but received %s", value[1:])
	}
}
