package main

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		if (1<<i)&num > 0 {
			ans += 1 << (31 - i)
		}
	}
	return ans
}
