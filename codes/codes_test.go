package codes_test

import (
	"testing"

	"github.com/nolwn/marker/codes"
)

func TestColor(t *testing.T) {
	c := codes.Code{}

	value := c.Color("magenta").Value()

	if value != "\x1b[35m" {
		t.Fatalf("expected magenta escape code but received %s", value[1:])
	}
}
