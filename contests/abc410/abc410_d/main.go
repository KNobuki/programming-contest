package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/trees/redblacktree"
)

// 解答欄
func solve() {
	n, m := ni2()
	ans := -1
	g := make([][]Edge, n)
	for i := 0; i < m; i++ {
		a, b, w := ni3()
		a--
		b--
		g[a] = append(g[a], Edge{
			To:     b,
			Weight: w,
		})
	}
	type node struct {
		i       int
		xorCost int
	}
	vis := make(map[int][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, 1024)
	}
	que := make([]node, 0, n*1024)
	start := node{
		i:       0,
		xorCost: 0,
	}
	vis[0][0] = true
	que = append(que, start)
	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		for _, edge := range g[cur.i] {
			next := node{
				i:       edge.To,
				xorCost: cur.xorCost ^ edge.Weight,
			}
			if vis[next.i][next.xorCost] {
				continue
			}
			vis[next.i][next.xorCost] = true
			que = append(que, next)
			if next.i == n-1 && ans == -1 {
				ans = next.xorCost
				continue
			}
			if next.i == n-1 {
				ans = min(ans, next.xorCost)
			}
		}
	}
	out.Println(ans)
}

const bufsize = 4 * 1024 * 1024
const MaxInt = int(^uint(0) >> 1)
const mod = 998244353

var in = bufio.NewScanner(os.Stdin)
var out Out

type Out struct {
	writer *bufio.Writer
}

func (out *Out) Println(a ...interface{}) {
	fmt.Fprintln(out.writer, a...)
}

func (out *Out) Printf(format string, a ...interface{}) {
	fmt.Fprintf(out.writer, format, a...)
}

func (out *Out) PrintStringsln(a []string) {
	out.Println(strings.Join(a, " "))
}

func (out *Out) PrintIntsLn(a []int) {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = fmt.Sprint(v)
	}
	out.Println(strings.Join(b, " "))
}

func (out *Out) PrintLenAndIntsLn(a []int) {
	b := make([]string, len(a)+1)
	b[0] = fmt.Sprint(len(a))
	for i, v := range a {
		b[i+1] = fmt.Sprint(v)
	}
	out.Println(strings.Join(b, " "))
}

func (out *Out) Putc(c byte) {
	out.Printf("%c", c)
}

func (out *Out) YesNo(cond bool) {
	if cond {
		out.Println("Yes")
	} else {
		out.Println("No")
	}
}

func (out *Out) YESNO(cond bool) {
	if cond {
		out.Println("YES")
	} else {
		out.Println("NO")
	}
}

func (out *Out) Print2DIntArray(a [][]int, format string) {
	for _, aa := range a {
		for _, v := range aa {
			out.Printf(format, v)
		}
		out.Println()
	}
}

func (out *Out) Print2DBoolArray(a [][]bool, format string) {
	for _, aa := range a {
		for _, v := range aa {
			out.Printf(format, v)
		}
		out.Println()
	}
}

func ns() string {
	in.Scan()
	return in.Text()
}

func nss(n int) []string {
	res := make([]string, n)
	for i := range res {
		res[i] = ns()
	}
	return res
}

func ni() int {
	in.Scan()
	ret, e := strconv.Atoi(in.Text())
	if e != nil {
		panic(e)
	}
	return ret
}

func nf() float64 {
	in.Scan()
	ret, e := strconv.ParseFloat(in.Text(), 64)
	if e != nil {
		panic(e)
	}
	return ret
}

func ni2() (int, int) {
	return ni(), ni()
}

func ni3() (int, int, int) {
	return ni(), ni(), ni()
}

func ni4() (int, int, int, int) {
	return ni(), ni(), ni(), ni()
}

func nis(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = ni()
	}
	return res
}

func sumOfInts(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxOfInts(a []int) int {
	res := -MaxInt
	for _, v := range a {
		res = max(res, v)
	}
	return res
}

func minOfInts(a []int) int {
	res := MaxInt
	for _, v := range a {
		res = min(res, v)
	}
	return res
}

func uniqueInts(a []int) []int {
	m := make(map[int]bool)
	res := make([]int, 0, len(m))
	for _, v := range a {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func uniqueStrings(a []string) []string {
	m := make(map[string]bool)
	res := make([]string, 0, len(m))
	for _, v := range a {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = res * a
		}
		a = a * a
		n >>= 1
	}
	return res
}

func gcd(x int, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func lcm(x int, y int) int {
	return x * y / gcd(x, y)
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func primeFact(x int) map[int]int {
	ret := make(map[int]int, x)
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			ret[i] += 1
			x /= i
		}
	}
	if x != 1 {
		ret[x] += 1
	}
	return ret
}

func divisor(x int) []int {
	ret := make([]int, 0, x)
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			ret = append(ret, i)
			if x/i != i {
				ret = append(ret, x/i)
			}
		}
	}
	return ret
}

type Dsu struct {
	n            int
	parentOrSize []int
	edgeNum      []int
}

func NewDsu(n int) *Dsu {
	d := &Dsu{
		n:            n,
		parentOrSize: make([]int, n),
		edgeNum:      make([]int, n),
	}
	for i := range d.parentOrSize {
		d.parentOrSize[i] = -1
	}
	return d
}

func (d *Dsu) Merge(a, b int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		d.edgeNum[x]++
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	d.edgeNum[x] += d.edgeNum[y] + 1
	return x
}

func (d *Dsu) Same(a, b int) bool {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	return d.Leader(a) == d.Leader(b)
}

func (d *Dsu) Leader(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d *Dsu) Size(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	return -d.parentOrSize[d.Leader(a)]
}

func (d *Dsu) EdgeNum(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	return d.edgeNum[d.Leader(a)]
}

func (d *Dsu) Groups() [][]int {
	leaderBuf, groupSize := make([]int, d.n), make([]int, d.n)
	for i := 0; i < d.n; i++ {
		leaderBuf[i] = d.Leader(i)
		groupSize[leaderBuf[i]]++
	}
	result := make([][]int, d.n)
	for i := 0; i < d.n; i++ {
		result[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < d.n; i++ {
		result[leaderBuf[i]] = append(result[leaderBuf[i]], i)
	}
	eraseEmpty := func(a [][]int) [][]int {
		result := make([][]int, 0, len(a))
		for i := range a {
			if len(a[i]) != 0 {
				result = append(result, a[i])
			}
		}
		return result
	}
	return eraseEmpty(result)
}

func dijkstra(N int, start int, graph [][]Edge) []int {
	dist := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dist[i] = MaxInt
	}
	dist[start] = 0
	h := &EdgeHeap{
		{To: start, Weight: 0},
	}
	heap.Init(h)
	for h.Len() > 0 {
		edge := heap.Pop(h).(Edge)
		position := edge.To
		if dist[position] != edge.Weight {
			continue
		}
		for _, p := range graph[position] {
			to, weight := p.To, dist[position]+p.Weight
			if dist[to] <= weight {
				continue
			}
			dist[to] = weight
			heap.Push(h, Edge{Weight: dist[to], To: to})
		}
	}
	return dist
}

func dijkstraWithPath(N int, start int, graph [][]Edge) ([]int, []int) {
	dist := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dist[i] = MaxInt
	}
	from := make([]int, N+1)
	for i := 0; i <= N; i++ {
		from[i] = -1
	}
	dist[start] = 0
	h := &EdgeHeap{
		{To: start, Weight: 0},
	}
	heap.Init(h)
	for h.Len() > 0 {
		edge := heap.Pop(h).(Edge)
		position := edge.To
		if dist[position] != edge.Weight {
			continue
		}
		for _, p := range graph[position] {
			to, weight := p.To, dist[position]+p.Weight
			if dist[to] <= weight {
				continue
			}
			dist[to] = weight
			heap.Push(h, Edge{Weight: dist[to], To: to})
			from[to] = position
		}
	}
	return dist, from
}

type Edge struct {
	To     int
	Weight int
}
type EdgeHeap []Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push データ格納
func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

// Pop データ取り出し
func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var fac, ifac, inv []int

func combInit(n int) {
	fac = make([]int, n+1)
	ifac = make([]int, n+1)
	inv = make([]int, n+1)

	fac[0], fac[1] = 1, 1
	ifac[0], ifac[1] = 1, 1
	inv[1] = 1

	for i := 1; i < n; i++ {
		fac[i+1] = fac[i] * (i + 1) % mod
		inv[i+1] = mod - inv[mod%(i+1)]*(mod/(i+1))%mod
		ifac[i+1] = ifac[i] * inv[i+1] % mod
	}
}

func comb(a, b int) int { return (fac[a] * ifac[b] % mod) * ifac[a-b] % mod }
func perm(a, b int) int { return fac[a] * ifac[a-b] % mod }

func madd(a, b int) int {
	return ((a % mod) + (b % mod)) % mod
}
func msub(a, b int) int {
	return (a - b + mod) % mod
}
func mmul(a, b int) int {
	return ((a % mod) * (b % mod)) % mod
}
func minv(a int) int {
	return mpow(a, mod-2)
}
func mdiv(a, b int) int {
	return ((a % mod) * minv(b)) % mod
}

func mpow(a, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n >>= 1
	}
	return res
}

func reverseInts(a []int) []int {
	res := make([]int, len(a))
	copy(res, a)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func reverseBytes(a []byte) []byte {
	res := make([]byte, len(a))
	copy(res, a)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func reverseString(a string) string {
	return string(reverseBytes([]byte(a)))
}

func eratosthenes(n int) []bool {
	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if !prime[i] {
			continue
		}
		for j := i * 2; j <= n; j += i {
			prime[j] = false
		}
	}
	return prime
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewIntArray(n, init int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = init
	}
	return ret
}

func New2DIntArray(n, m, init int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, m)
		for j := 0; j < m; j++ {
			ret[i][j] = init
		}
	}
	return ret
}

// nextPermutation
// example: for ns, a := true, genPerm(k); ns; ns = nextPermutation(a)
func nextPermutation(aa []int) bool {
	n := len(aa)
	l := n - 2
	for l >= 0 && aa[l] > aa[l+1] {
		l--
	}
	if l < 0 {
		return false
	}
	r := n - 1
	for l < r && aa[l] > aa[r] {
		r--
	}
	aa[l], aa[r] = aa[r], aa[l]
	l++
	r = n - 1
	for l < r {
		aa[l], aa[r] = aa[r], aa[l]
		l++
		r--
	}
	return true
}

func genPerm(size int) []int {
	if size < 0 {
		panic("genPerm: size is negative")
	}
	ret := make([]int, size)
	for i := 0; i < size; i++ {
		ret[i] = i
	}
	return ret
}

// SegTree
// example:
//
//	a := make([]int, n-1)
//	op := func(a, b int) int {
//		return a * b
//	}
//	e := func() int {
//		return 1
//	}
//	seg := NewSegTreeFromSlice(op, e, a)
type SegTree[T any] struct {
	n, size, log int
	d            []T
	op           func(a, b T) T
	e            func() T
}

func (s *SegTree[T]) update(k int) {
	s.d[k] = s.op(s.d[2*k], s.d[2*k+1])
}

func NewSegTree[T any](op func(a, b T) T, e func() T, n int) *SegTree[T] {
	v := make([]T, n)
	for i := 0; i < n; i++ {
		v[i] = e()
	}
	return NewSegTreeFromSlice(op, e, v)
}

func CeilPow2(n int) int {
	x := 0
	for (uint(1) << x) < uint(n) {
		x++
	}
	return x
}

func NewSegTreeFromSlice[T any](op func(a, b T) T, e func() T, v []T) *SegTree[T] {
	n := len(v)
	log := CeilPow2(n)
	size := 1 << log
	d := make([]T, 2*size)
	for i := 0; i < 2*size; i++ {
		d[i] = e()
	}
	for i := 0; i < n; i++ {
		d[size+i] = v[i]
	}
	for i := size - 1; i >= 1; i-- {
		d[i] = op(d[2*i], d[2*i+1])
	}

	return &SegTree[T]{
		n:    n,
		size: size,
		log:  log,
		d:    d,
		op:   op,
		e:    e,
	}
}

func (s *SegTree[T]) Set(p int, x T) {
	if p < 0 || s.n <= p {
		panic("")
	}
	p += s.size
	s.d[p] = x
	for i := 1; i <= s.log; i++ {
		s.update(p >> i)
	}
}

func (s *SegTree[T]) Get(p int) T {
	if p < 0 || s.n <= p {
		panic("")
	}
	return s.d[p+s.size]
}

func (s *SegTree[T]) Prod(l, r int) T {
	if l < 0 || r < l || s.n < r {
		panic("")
	}
	sml, smr := s.e(), s.e()
	l += s.size
	r += s.size

	for l < r {
		if l&1 > 0 {
			sml = s.op(sml, s.d[l])
			l++
		}
		if r&1 > 0 {
			r--
			smr = s.op(s.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return s.op(sml, smr)
}

func (s *SegTree[T]) AllProd() T {
	return s.d[1]
}

func (s *SegTree[T]) MaxRight(l int, f func(x T) bool) int {
	if l < 0 || s.n < l {
		panic("")
	}
	if !f(s.e()) {
		panic("")
	}
	if l == s.n {
		return s.n
	}
	l += s.size
	sm := s.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !f(s.op(sm, s.d[l])) {
			for l < s.size {
				l = 2 * l
				if f(s.op(sm, s.d[l])) {
					sm = s.op(sm, s.d[l])
					l++
				}
			}
			return l - s.size
		}
		sm = s.op(sm, s.d[l])
		l++
		if (l & -l) == l {
			break
		}
	}
	return s.n
}

func (s *SegTree[T]) MinLeft(r int, f func(x T) bool) int {
	if r < 0 || s.n < r {
		panic("")
	}
	if !f(s.e()) {
		panic("")
	}
	if r == 0 {
		return 0
	}
	r += s.size
	sm := s.e()
	for {
		r--
		for r > 1 && r%2 == 1 {
			r >>= 1
		}
		if !f(s.op(s.d[r], sm)) {
			for r < s.size {
				r = 2*r + 1
				if f(s.op(s.d[r], sm)) {
					sm = s.op(s.d[r], sm)
					r--
				}
			}
			return r + 1 - s.size
		}
		sm = s.op(s.d[r], sm)
		if (r & -r) == r {
			break
		}
	}
	return 0
}

// SegTreeLazy
// example:
// ex := X{0}
//
//	fx := func(a, b X) X {
//		return X{max(a.val, b.val)}
//	}
//
// em := M{0}
//
//	fm := func(a, b M) M {
//		return M{a.val + b.val}
//	}
//	fa := func(a X, b M) X {
//		return X{a.val + b.val}
//	}
//
// seg := NewSegTreeLazy(n, ex, em, fx, fm, fa)
type SegTreeLazy struct {
	Size int
	Ex   X  // identity element in monoid X
	Em   M  // identity element in monoid M
	Fx   FX // binary operation in monoid X
	Fm   FM // binary operation in monoid M
	Fa   FA // binary operation between X and M
	Dat  []X
	Lazy []M
}

type X struct {
	val int
}
type M struct {
	val int
}
type FX func(a, b X) X
type FM func(a, b M) M
type FA func(a X, b M) X

func NewSegTreeLazy(n int, ex X, em M, fx FX, fm FM, fa FA) *SegTreeLazy {
	sg := &SegTreeLazy{
		Size: 1,
		Fx:   fx,
		Fm:   fm,
		Fa:   fa,
		Ex:   ex,
		Em:   em,
	}

	for sg.Size < n {
		sg.Size *= 2
	}
	sg.Dat = make([]X, 2*sg.Size)
	for i := range sg.Dat {
		sg.Dat[i] = sg.Ex
	}
	sg.Lazy = make([]M, 2*sg.Size)
	for i := range sg.Lazy {
		sg.Lazy[i] = sg.Em
	}
	return sg
}

// Set and Build used for bulk construction in O(n)
func (sg *SegTreeLazy) Set(k int, x X) {
	sg.Dat[k+sg.Size-1] = x
}
func (sg *SegTreeLazy) Build() {
	for k := sg.Size - 2; k >= 0; k-- {
		sg.Dat[k] = sg.Fx(sg.Dat[2*k+1], sg.Dat[2*k+2])
	}
}

// Eval propagates the lazy data
func (sg *SegTreeLazy) Eval(k int) {
	if sg.Lazy[k] == sg.Em {
		return
	}

	// propagate k-th lazy to its children
	if k < sg.Size-1 {
		sg.Lazy[2*k+1] = sg.Fm(sg.Lazy[2*k+1], sg.Lazy[k])
		sg.Lazy[2*k+2] = sg.Fm(sg.Lazy[2*k+2], sg.Lazy[k])
	}
	// update itself
	sg.Dat[k] = sg.Fa(sg.Dat[k], sg.Lazy[k])
	sg.Lazy[k] = sg.Em
}

func (sg *SegTreeLazy) update(a, b int, x M, k, l, r int) {
	sg.Eval(k)
	if a <= l && r <= b {
		sg.Lazy[k] = sg.Fm(sg.Lazy[k], x)
		sg.Eval(k)
	} else if a < r && l < b {
		sg.update(a, b, x, 2*k+1, l, (l+r)/2)
		sg.update(a, b, x, 2*k+2, (l+r)/2, r)
		sg.Dat[k] = sg.Fx(sg.Dat[2*k+1], sg.Dat[2*k+2])
	}
}

// Update updates [a, b) to x
func (sg *SegTreeLazy) Update(a, b int, x M) {
	sg.update(a, b, x, 0, 0, sg.Size)
}

func (sg *SegTreeLazy) query(a, b, k, l, r int) X {
	sg.Eval(k)
	if r <= a || b <= l {
		return sg.Ex
	}
	if a <= l && r <= b {
		return sg.Dat[k]
	}
	vl := sg.query(a, b, 2*k+1, l, (l+r)/2)
	vr := sg.query(a, b, 2*k+2, (l+r)/2, r)
	return sg.Fx(vl, vr)
}

// Query returns the query result in [a, b)
func (sg *SegTreeLazy) Query(a, b int) X {
	return sg.query(a, b, 0, 0, sg.Size)
}

func update(tree *redblacktree.Tree, key interface{}, x int) {
	old, found := tree.Get(key)
	if found {
		x += old.(int)
	}
	if x <= 0 {
		tree.Remove(key)
		return
	}
	tree.Put(key, x)
}
func increment(tree *redblacktree.Tree, key interface{}) {
	update(tree, key, 1)
}
func decrement(tree *redblacktree.Tree, key interface{}) {
	update(tree, key, -1)
}

/*
	type pqst struct {
		x int
		y int
	}

	pq := newpq([]compFunc{func(p, q interface{}) int {
		if p.(pqst).x != q.(pqst).x {
			// get from bigger
			// if p.(pqst).x > q.(pqst).x {
			if p.(pqst).x < q.(pqst).x {
				return -1
			} else {
				return 1
			}
		}
		if p.(pqst).y != q.(pqst).y {
			// get from bigger
			// if p.(pqst).y > q.(pqst).y {
			if p.(pqst).y < q.(pqst).y {
				return -1
			} else {
				return 1
			}
		}
		return 0
	}})
	heap.Init(pq)
	heap.Push(pq, pqst{x: 1, y: 1})
	for !pq.IsEmpty() {
		v := heap.Pop(pq).(pqst)
	}
*/

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

// pq.GetRoot().(edge)
func (pq *pq) GetRoot() interface{} {
	return pq.arr[0]
}

type runLength struct {
	c byte
	l int
}

func runLengthEncoding(s string) []runLength {
	res := make([]runLength, 0, len(s))
	for l := 0; l < len(s); {
		r := l
		for r < len(s)-1 && s[r] == s[r+1] {
			r++
		}
		res = append(res, runLength{
			c: s[l],
			l: r - l + 1,
		})
		l = r + 1
	}
	return res
}

func runLengthDecoding(rl []runLength) string {
	res := make([]byte, 0)
	for _, r := range rl {
		for i := 0; i < r.l; i++ {
			res = append(res, r.c)
		}
	}
	return string(res)
}

func zAlgorithm(s string) []int {
	if len(s) == 0 {
		return []int{}
	}
	l, r := 0, 0
	n := len(s)
	z := make([]int, n)
	z[0] = n
	for i := 1; i < n; i++ {
		if z[i-l] < r-i {
			z[i] = z[i-l]
		} else {
			r = max(r, i)
			for r < n && s[r] == s[r-i] {
				r++
			}
			z[i] = r - i
			l = i
		}
	}
	return z
}

func reverseSortIntSlice(a []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	return a
}

func init() {
	in.Split(bufio.ScanWords)
	in.Buffer(make([]byte, bufsize), bufsize)
	out.writer = bufio.NewWriterSize(os.Stdout, bufsize)
}

func main() {
	defer out.writer.Flush()
	solve()
}
