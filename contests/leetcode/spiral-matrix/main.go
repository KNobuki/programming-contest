package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	type vec struct {
		y int
		x int
	}
	dirs := []vec{
		{
			y: 0,
			x: 1,
		},
		{
			y: 1,
			x: 0,
		},
		{
			y: 0,
			x: -1,
		},
		{
			y: -1,
			x: 0,
		},
	}
	ans := make([]int, 0, m*n)
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	y, x, dir := 0, 0, 0
	for len(ans) < m*n {
		ans = append(ans, matrix[y][x])
		vis[y][x] = true
		ny, nx := y+dirs[dir].y, x+dirs[dir].x
		if ny < 0 || ny >= m || nx < 0 || nx >= n || vis[ny][nx] {
			dir = (dir + 1) % 4
			ny, nx = y+dirs[dir].y, x+dirs[dir].x
		}
		y, x = ny, nx
	}
	return ans
}
