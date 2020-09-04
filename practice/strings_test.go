package practice

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	builder := strings.Builder{}
	fmt.Println(builder.String())
}
