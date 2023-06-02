// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/

package fibonaccinumber

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	prevPrev, prev, ans := 0, 1, 0
	for i := 0; i < n-1; i++ {
		ans = prevPrev + prev
		prevPrev = prev
		prev = ans
	}
	return ans
}
