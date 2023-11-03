package bitmapplus

import "testing"

func TestBitMapPlus(t *testing.T) {
	m := NewBitMapPlus(7)
	nums := []uint32{6, 3, 2, 1, 7, 2, 6}
	for _, num := range nums {
		m.Set(num)
	}
	var result []uint32
	for _, num := range nums {
		//t.Logf("num=%d, result=%d\n", num, m.IsExists(num))
		if m.IsExists(num) == 1 {
			result = append(result, num)
		}
	}
	t.Log(result)
}
