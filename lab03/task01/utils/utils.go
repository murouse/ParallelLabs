package utils

func Kiss(n int, seed [4]uint64) []uint64 {
	x := make([]uint64, n)
	var t uint64

	for i := 0; i < n; i++ {
		seed[0] = 69069*seed[0] + 123456
		seed[1] = seed[1] ^ (seed[1] << 13)
		seed[1] = seed[1] ^ (seed[1] >> 17)
		seed[1] = seed[1] ^ (seed[1] << 5)
		t = 698769069*seed[2] + seed[3]
		seed[3] = t >> 32
		seed[1] = t
		x[i] = seed[0] + seed[1] + seed[2]
	}

	return x
}
