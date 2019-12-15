package w5_5

import "testing"

func TestFind(t *testing.T) {
	for _, e := range tests {
		res := Find(e.a, e.x)
		if res != e.exp {
			t.Errorf("Find(%v, %d) = %d, expected %d",
				e.a, e.x, res, e.exp)
		}
	}
}
