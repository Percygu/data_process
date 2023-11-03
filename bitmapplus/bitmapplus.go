package bitmapplus

type BitMapPlus struct {
	MaxValue  uint32
	ByteSlice []byte
}

func NewBitMapPlus(maxValue uint32) *BitMapPlus {
	return &BitMapPlus{
		MaxValue:  maxValue,
		ByteSlice: make([]byte, ((maxValue+1)*2+7)/8),
	}
}

func (m *BitMapPlus) Set(num uint32) {
	if num > m.MaxValue {
		return
	}

	byteIndex := num / 4
	b := m.ByteSlice[byteIndex]

	// 后一个bit位
	bitIndex2 := (num % 4) * 2
	// 前一个bit位
	bitIndex1 := bitIndex2 + 1

	//fmt.Printf("num=%d,bitIndex2=%d, bitIndex1=%d\n", num, bitIndex2, bitIndex1)

	// 表示num的两个bit位都为0，没出现过
	if (b&(1<<bitIndex1)) == 0 && (b&(1<<bitIndex2) == 0) {
		//fmt.Printf("before:num=%d,b=%b\n", num, b)
		// 把后一个bit位置为1
		m.ByteSlice[byteIndex] = b | (1 << bitIndex2)
		//fmt.Printf("after:num=%d,b=%b\n", num, m.ByteSlice[byteIndex])
		return
	}

	//fmt.Printf("existed:num=%d,bitIndex1=%d,bitIndex2=%d,1<<bitIndex1=%b,1<<bitIndex1=%b,b=%b\n", num, bitIndex1, bitIndex2, 1<<bitIndex1, 1<<bitIndex2, b)
	//fmt.Printf("b&(1<<bitIndex1))=%d,b&(1<<bitIndex2))=%d,\n", b&(1<<bitIndex1), b&(1<<bitIndex2))
	// 前一个bit位位0，后一个bit位为1，出现过1次
	if (b&(1<<bitIndex1)) == 0 && (b&(1<<bitIndex2) != 0) {
		// 把前一个bit位置为1
		//fmt.Printf("before:num=%d,b=%b\n", num, b)
		m.ByteSlice[byteIndex] = b | (1 << bitIndex1)
		//fmt.Printf("after:num=%d,b=%b\n", num, m.ByteSlice[byteIndex])
	}
}

func (m *BitMapPlus) IsExists(num uint32) int {
	if num > m.MaxValue {
		return 0
	}
	byteIndex := num / 4
	b := m.ByteSlice[byteIndex]
	// 后一个bit位
	bitIndex2 := (num % 4) * 2
	// 前一个bit位
	bitIndex1 := bitIndex2 + 1
	if (b&(1<<bitIndex1)) == 0 && (b&(1<<bitIndex2) == 0) {
		return 0
	}
	if (b&(1<<bitIndex1)) == 0 && (b&(1<<bitIndex2) != 0) {
		return 1
	}
	if (b & (1 << bitIndex1)) == 1 {
		return 2
	}
	return 0
}
