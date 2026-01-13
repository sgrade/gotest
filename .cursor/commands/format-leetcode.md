# format-leetcode

Review the code and apply [Effective Go](https://go.dev/doc/effective_go) and the [Google Go Style Guide](https://google.github.io/styleguide/go/) with the following exceptions:
- Do NOT change function names or signatures specified by LeetCode (they may not follow Go conventions).
- Package documentation and function documentation are OPTIONAL (the problem comment at the top serves this purpose).

## Standard Problem Annotation

Ensure the file starts with exactly this format:
```go
// [number]. [Problem Title]
// https://leetcode.com/problems/[problem-slug]/

package problemslug

// [Optional comment about approach or time/space complexity]
func solutionFunction() {
```

Example:
```go
// 3453. Separate Squares I
// https://leetcode.com/problems/separate-squares-i/

package separatesquaresi

// Binary search to find the horizontal line that divides total area in half.
// Time: O(n log maxY), Space: O(1)
func separateSquares(squares [][]int) float64 {
```
