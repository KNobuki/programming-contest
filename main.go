package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const bufsize = 4 * 1024 * 1024

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

func maxOfInts(a []int) int {
	max := 0
	for _, v := range a {
		if v >= max {
			max = v
		}
	}
	return max
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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func pow(x int, y int) int {
	ret := 1
	for i := 0; i < y; i++ {
		ret *= x
	}
	return ret
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

func init() {
	in.Split(bufio.ScanWords)
	in.Buffer(make([]byte, bufsize), bufsize)
	out.writer = bufio.NewWriterSize(os.Stdout, bufsize)
}

func main() {
	defer out.writer.Flush()
	solve()
}

func solve() {
	
}
