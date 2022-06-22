package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) (retCount int) {
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			retCount++
		}
		x >>= 1
	}
	return retCount
}
