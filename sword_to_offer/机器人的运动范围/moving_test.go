package movingCount

import "testing"

func TestName(t *testing.T) {
	if out := movingCount(11, 8, 16); out != 88 {
		t.Fatal(88, out)
	}
	if out := movingCount(2, 3, 1); out != 3 {
		t.Fatal(3, out)
	}
	if out := movingCount(3, 1, 0); out != 1 {
		t.Fatal(1, out)
	}
}
