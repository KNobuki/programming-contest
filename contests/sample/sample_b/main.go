package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

// 実際の処理
func main() {
	rbt := redblacktree.NewWithIntComparator()
	for i := 0; i < 1e7; i++ {
		// データのPut
		rbt.Put(i, i)
	}
	for i := 0; i < 1e7; i++ {
		// データのGet
		_, _ = rbt.Get(i)
	}
}
