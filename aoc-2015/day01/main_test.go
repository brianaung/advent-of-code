package main

import (
	"testing"
)

func TestNotQuiteLispPart1(t *testing.T) {
	want := 0
	if result := NotQuiteLisp([]byte("(())"), 1); result != want {
		t.Fatalf(`NotQuiteLisp("(())", 1) = %q, want match for %#q, nil`, result, want)
	}

	want = 3
	if result := NotQuiteLisp([]byte("))((((("), 1); result != want {
		t.Fatalf(`NotQuiteLisp("))(((((", 1) = %q, want match for %#q, nil`, result, want)
	}

	want = -1
	if result := NotQuiteLisp([]byte("))("), 1); result != want {
		t.Fatalf(`NotQuiteLisp("))(", 1) = %q, want match for %#q, nil`, result, want)
	}

	want = -3
	if result := NotQuiteLisp([]byte(")())())"), 1); result != want {
		t.Fatalf(`NotQuiteLisp(")())())", 1) = %q, want match for %#q, nil`, result, want)
	}
}

func TestNotQuiteLispPart2(t *testing.T) {
	want := 1
	if result := NotQuiteLisp([]byte(")"), 2); result != want {
		t.Fatalf(`NotQuiteLisp(")", 2) = %q, want match for %#q, nil`, result, want)
	}

	want = 5
	if result := NotQuiteLisp([]byte("()())"), 2); result != want {
		t.Fatalf(`NotQuiteLisp("()())", 2) = %q, want match for %#q, nil`, result, want)
	}
}
