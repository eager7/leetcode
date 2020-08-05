package dp1

import "testing"

func TestName(t *testing.T) {
	if numWays(2) != 2 {
		t.Fatal(numWays(2))
	}
	if numWays(0) != 1 {
		t.Fatal(numWays(0))
	}
	if numWays(7) != 21 {
		t.Fatal(numWays(7))
	}
}
