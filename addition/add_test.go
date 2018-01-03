package addition

import "testing"

func TestSum(t *testing.T){
	var a, b, c int
	a = 1
	b = 2
	//exported function
	c = Sum(a,b)
	if c != 3 {
	 t.Error("Expected 3, got",c)
	}
}
