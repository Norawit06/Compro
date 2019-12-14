package w5_3

import "testing"

func TestaddMult(t *testing.T)
{
	a := addMult(7, 5)
	 
	if a != 14, 35{
		t.Errorf("want 17, got%v", a)
	}
}