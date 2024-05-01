package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	zCols := make(map[int]bool, n)
	zRows := make(map[int]bool, m)
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if matrix[y][x] == 0 {
				zCols[x] = true
				zRows[y] = true
			}
		}
	}
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if zCols[x] || zRows[y] {
				matrix[y][x] = 0
			}
		}
	}
}
