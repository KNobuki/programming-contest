package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 解答欄
func solve() {
	n, l := nextInt2()
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		if dp[i] == 0 {
			continue
		}
		if i+1 <= n {
			dp[i+1] = madd(dp[i+1], dp[i])
		}
		if i+l <= n {
			dp[i+l] = madd(dp[i+l], dp[i])
		}
	}
	out.Println(dp[n])
}

const bufsize = 4 * 1024 * 1024
const MaxInt = int(^uint(0) >> 1)
const mod = 1000000007

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

func next() string {
	in.Scan()
	return in.Text()
}

func nexts(n int) []string {
	res := make([]string, n)
	for i := range res {
		res[i] = next()
	}
	return res
}

func nextInt() int {
	in.Scan()
	ret, e := strconv.Atoi(in.Text())
	if e != nil {
		panic(e)
	}
	return ret
}

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func nextInts(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = nextInt()
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

type Dsu struct {
	n            int
	parentOrSize []int
}

func NewDsu(n int) *Dsu {
	d := &Dsu{
		n:            n,
		parentOrSize: make([]int, n),
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
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
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
	/* ダイクストラ法 */
	// 最短距離が確定したかどうかのリスト
	confirm := make([]bool, N+1)
	// 最短距離リスト
	dist := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dist[i] = MaxInt
	}
	// スタート地点をキューに追加
	dist[start] = 0
	h := &EdgeHeap{
		{To: start, Weight: 0},
	}
	heap.Init(h)
	// ダイクストラ法
	// 確定候補から一番最短距離が近いとこを確定させていくのに
	// 毎回ループで探すのは大変なため
	// ヒープキューで最小ノードを取り出していく
	// (確定候補がなくなったら終了)
	for h.Len() > 0 {
		// ヒープからキュー取得
		edge := heap.Pop(h).(Edge)

		// 次に確定させるべき頂点を求める
		position := edge.To
		// すでに最短距離が確定している場合
		if confirm[position] {
			continue
		}

		// 最短距離確定を更新する
		confirm[position] = true
		// 隣接しているノードをループする
		for _, p := range graph[position] {
			// 次のノード, 1~positionまでの最短距離 + 現在のノードから次ぎのノードにいくコスト
			to, weight := p.To, dist[position]+p.Weight
			// 現在のノードから行った場合に、他のノードから行った場合より距離が短い場合
			if dist[to] > weight {
				// 最短距離リスト更新
				dist[to] = weight
				// 確定候補キューに格納
				heap.Push(h, Edge{Weight: dist[to], To: to})
			}
		}
	}
	return dist
}

func dijkstraWithPath(N int, start int, graph [][]Edge) ([]int, []int) {
	/* ダイクストラ法 */
	// 最短距離が確定したかどうかのリスト
	confirm := make([]bool, N+1)
	// 最短距離リスト
	dist := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dist[i] = MaxInt
	}
	// 経路を保存するためのリスト
	from := make([]int, N+1)
	for i := 0; i <= N; i++ {
		from[i] = -1
	}
	// スタート地点をキューに追加
	dist[start] = 0
	h := &EdgeHeap{
		{To: start, Weight: 0},
	}
	heap.Init(h)
	// ダイクストラ法
	// 確定候補から一番最短距離が近いとこを確定させていくのに
	// 毎回ループで探すのは大変なため
	// ヒープキューで最小ノードを取り出していく
	// (確定候補がなくなったら終了)
	for h.Len() > 0 {
		// ヒープからキュー取得
		edge := heap.Pop(h).(Edge)

		// 次に確定させるべき頂点を求める
		position := edge.To
		// すでに最短距離が確定している場合
		if confirm[position] {
			continue
		}

		// 最短距離確定を更新する
		confirm[position] = true
		// 隣接しているノードをループする
		for _, p := range graph[position] {
			// 次のノード, 1~positionまでの最短距離 + 現在のノードから次ぎのノードにいくコスト
			to, weight := p.To, dist[position]+p.Weight
			// 現在のノードから行った場合に、他のノードから行った場合より距離が短い場合
			if dist[to] > weight {
				// 最短距離リスト更新
				dist[to] = weight
				// 確定候補キューに格納
				heap.Push(h, Edge{Weight: dist[to], To: to})
				// 経路リスト更新
				from[to] = position
			}
		}
	}
	return dist, from
}

type Edge struct {
	To     int // 次に移動できるノード
	Weight int // 移動にかかる重み(コスト)
	idx    int
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

func init() {
	in.Split(bufio.ScanWords)
	in.Buffer(make([]byte, bufsize), bufsize)
	out.writer = bufio.NewWriterSize(os.Stdout, bufsize)
}

func main() {
	defer out.writer.Flush()
	solve()
}
