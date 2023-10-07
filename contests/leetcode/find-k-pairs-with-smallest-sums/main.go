package main

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	type pqst struct {
		i1  int
		i2  int
		sum int
	}
	p := newpq([]compFunc{func(p, q interface{}) int {
		if p.(pqst).sum < q.(pqst).sum {
			return -1
		}
		if p.(pqst).sum > q.(pqst).sum {
			return 1
		}
		return 0
	}})
	heap.Init(p)
	heap.Push(p, pqst{
		i1:  0,
		i2:  0,
		sum: nums1[0] + nums2[0],
	})
	var ret [][]int
	vis := make(map[pqst]bool)
	for i := 0; i < k && !p.IsEmpty(); i++ {
		now := heap.Pop(p).(pqst)
		ret = append(ret, []int{nums1[now.i1], nums2[now.i2]})
		if now.i1+1 < len(nums1) {
			np := pqst{
				i1:  now.i1 + 1,
				i2:  now.i2,
				sum: nums1[now.i1+1] + nums2[now.i2],
			}
			if !vis[np] {
				heap.Push(p, np)
				vis[np] = true
			}
		}
		if now.i2+1 < len(nums2) {
			np := pqst{
				i1:  now.i1,
				i2:  now.i2 + 1,
				sum: nums1[now.i1] + nums2[now.i2+1],
			}
			if !vis[np] {
				heap.Push(p, np)
				vis[np] = true
			}
		}
	}
	return ret
}

type pq struct {
	arr   []interface{}
	comps []compFunc
}

type compFunc func(p, q interface{}) int

func newpq(comps []compFunc) *pq {
	return &pq{
		comps: comps,
	}
}

func (pq pq) Len() int {
	return len(pq.arr)
}

func (pq pq) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq pq) Less(i, j int) bool {
	for _, comp := range pq.comps {
		result := comp(pq.arr[i], pq.arr[j])
		switch result {
		case -1:
			return true
		case 1:
			return false
		case 0:
			continue
		}
	}
	return true
}

func (pq *pq) Push(x interface{}) {
	pq.arr = append(pq.arr, x)
}

func (pq *pq) Pop() interface{} {
	n := pq.Len()
	item := pq.arr[n-1]
	pq.arr = pq.arr[:n-1]
	return item
}

func (pq *pq) IsEmpty() bool {
	return pq.Len() == 0
}
