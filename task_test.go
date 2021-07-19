package ozon_task

import (
	"testing"
)

func TestMaxOnesAfterRemoveItem(t *testing.T) {
	cases := []struct {
		in   []byte
		want uint
	}{
		{[]byte{0, 0}, 0},
		{[]byte{0, 1}, 1},
		{[]byte{1, 0}, 1},
		{[]byte{1, 1}, 1},
		{[]byte{1, 1, 0, 1, 1}, 4},
		{[]byte{1, 1, 0, 1, 1, 0, 1, 1, 1}, 5},
		{[]byte{1, 1, 0, 1, 1, 0, 1, 1, 1, 0}, 5},

		{[]byte{1}, 0},
		{[]byte{0, 1, 1}, 2},
		{[]byte{1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1}, 6},
		{[]byte{1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1}, 5},
	}

	for _, c := range cases {
		got := maxOnesAfterRemoveItem(c.in)
		if got != c.want {
			t.Errorf("maxOnesAfterRemoveItem(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
