package main

import "fmt"

func main() {
	board := []string{"ABCE", "SFCS", "ADEE"}
	bboard := make([][]byte, 0)
	for i := 0; i < len(board); i++ {
		bboard = append(bboard, make([]byte, 0, len(board[i])))
		for j := 0; j < len(board[i]); j++ {
			bboard[i] = append(bboard[i], board[i][j])
		}
	}
	word := "ABCB"
	fmt.Println(exist(bboard, word))
}

func exist(board [][]byte, word string) bool {
	h := len(board)
	w := len(board[0])
	vis := make([][]bool, len(board))
	for i := 0; i < h; i++ {
		vis[i] = make([]bool, w)
	}
	dy := []int{-1, 0, 1, 0}
	dx := []int{0, 1, 0, -1}
	var dfs func(y, x, c int) bool
	dfs = func(y, x, c int) bool {
		if c == len(word)-1 {
			return true
		}
		vis[y][x] = true
		for i := 0; i < 4; i++ {
			ny := y + dy[i]
			nx := x + dx[i]
			if ny >= h || ny < 0 || nx >= w || nx < 0 || vis[ny][nx] {
				continue
			}
			if word[c+1] == board[ny][nx] && dfs(ny, nx, c+1) {
				return true
			}
		}
		vis[y][x] = false
		return false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if word[0] == board[i][j] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
