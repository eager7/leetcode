package longestPalindrome

import (
	"testing"
)

func TestExample(t *testing.T) {
	type Te struct {
		In  string
		Out string
	}

	var list = []Te{
		{In: "aba", Out: "aba"},
		{In: "abc", Out: "a"},
		{In: "abba", Out: "abba"},
		{In: "babad", Out: "bab"},
		{In: "cbbd", Out: "bb"},
	}

	for _, te := range list {
		if my(te.In) != te.Out {
			t.Fatal(te.In + "|" + te.Out)
		}
	}
	for _, te := range list {
		if best(te.In) != te.Out {
			t.Fatal(te.In + "|" + te.Out)
		}
	}

}
