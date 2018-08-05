package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var c byte
	var i uint8
	for ; i < 8; i++ {
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}

func PopCount3(x uint64) int {
	var c int
	var i uint8
	for ; i < 64; i++ {
		if x == 0 {
			break
		}
		x &= x - 1
		c++
	}
	return c
}
