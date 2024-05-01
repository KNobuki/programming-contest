package main

import (
	"fmt"
)

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(mat)
	fmt.Println(mat)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	rotated := make([][]bool, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]bool, n)
	}
	var dfs func(y, x int)
	dfs = func(y, x int) {
		if rotated[y][x] {
			return
		}
		rotated[y][x] = true
		now := matrix[y][x]
		ny, nx := x, n-(y+1)
		dfs(ny, nx)
		matrix[ny][nx] = now
	}
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			dfs(y, x)
		}
	}
}
