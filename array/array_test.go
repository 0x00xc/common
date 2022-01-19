package array

import "testing"

func TestIn(t *testing.T) {
	//t.Log(In([]int64{1, 2, 3}, 3))
	//t.Log(In([]float64{1, 2, 3}, 3.0))
	//t.Log(In([]string{"1","2", "3"}, "3"))
	t.Log(In("123", byte('3')))
	t.Log(In(map[string]interface{}{"1": 1}, 1))
}
