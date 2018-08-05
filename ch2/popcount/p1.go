package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount1(x uint64) int {
	var c byte
	var i uint8
	for ; i < 8; i++ {
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}
