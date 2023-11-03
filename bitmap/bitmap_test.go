package bitmap

import "testing"

func TestBitMap(t *testing.T) {
	bitMap := NewBitMap(100)
	bitMap.Set(10)
	bitMap.Set(50)
	bitMap.Set(80)

	t.Log(bitMap.IsExists(10))
	t.Log(bitMap.IsExists(50))
	t.Log(bitMap.IsExists(80))
	t.Logf("\n")
	t.Log(bitMap.IsExists(20))
	t.Log(bitMap.IsExists(30))
}
