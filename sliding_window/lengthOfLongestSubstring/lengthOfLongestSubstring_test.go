package lengthOfLongestSubstring

import "testing"

func TestLengthOfSub(t *testing.T) {
	sub1 := "abcabcbb"
	sub2 := "bbbbb"
	sub3 := "pwwkew"
	sub4 := "aab"
	sub5 := "dvdf"

	if my(sub1) != 3 {
		t.Fatal(sub1)
	}
	if my(sub2) != 1 {
		t.Fatal(sub2)
	}
	if my(sub3) != 3 {
		t.Fatal(sub3)
	}
	if my(sub4) != 2 {
		t.Fatal(sub3)
	}
	if my(sub5) != 3 {
		t.Fatal(sub3)
	}
	if best(sub1) != 3 {
		t.Fatal(best(sub1))
	}
	if best(sub2) != 1 {
		t.Fatal(sub2)
	}
	if best(sub3) != 3 {
		t.Fatal(best(sub3))
	}
	if best(sub4) != 2 {
		t.Fatal(best(sub4))
	}
	if best(sub5) != 3 {
		t.Fatal(best(sub5))
	}
}
