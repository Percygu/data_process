package bitmap

type BitMap struct {
	MaxValue  uint32
	ByteSlice []byte // 用byte数组来实现位数组
}

func NewBitMap(maxValue uint32) *BitMap {
	return &BitMap{
		MaxValue:  maxValue,
		ByteSlice: make([]byte, (maxValue+8)/8), // +7: 向上取整，不足一个byte，凑齐一个
	}
}

// Set 将给定数的二进制位置1
func (m *BitMap) Set(num uint32) {
	if num > m.MaxValue {
		return
	}
	// 找到num所在的二进制位
	byteIndex := num / 8
	b := m.ByteSlice[byteIndex] // 确定当前给定的数在bite数组的第几个

	bitIndex := num % 8 // 确定当前给定的数在这个字节的第几位

	m.ByteSlice[byteIndex] = b | (1 << bitIndex)
}

func (m *BitMap) IsExists(num uint32) bool {
	if num > m.MaxValue {
		return false
	}
	byteIndex := num / 8
	bitIndex := num % 8
	b := m.ByteSlice[byteIndex]
	return (b & (1 << bitIndex)) != 0 // !=0 说明存在，返回true，否则返回false
}
