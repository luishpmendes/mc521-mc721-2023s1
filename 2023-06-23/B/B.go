// https://codeforces.com/problemset/problem/1816/B

// CodeForces 1816B - Grid Reconstruction
// Consider a 2×n grid, where n is an even integer. You may place the integers 1,2,…,2n on the grid, using each integer exactly once.
// A path is a sequence of cells achieved by starting at (1,1), then repeatedly walking either downwards or to the right, and stopping when (2,n) is reached. The path should not extend beyond the grid.
// The cost of a path is the alternating sum of the numbers written on the cells in a path. That is, let the numbers written on the cells be a1,a2,…,ak (in the order that it is visited), the cost of the path is a1−a2+a3−a4+…=∑ki=1ai⋅(−1)i+1.
// Construct a way to place the integers 1,2,…,2n on the grid, such that the minimum cost over all paths from (1,1) to (2,n) is maximized. If there are multiple such grids that result in the maximum value, output any of them.

// Input
// The first line contains a single integer t (1≤t≤1000) — the number of test cases. The description of test cases follows.
// The first and the only line of each test case contains a single integer n (2≤n≤10^5, n is even) — the number of the columns in the grid.
// It is guaranteed that the sum of n over all test cases does not exceed 10^5.

// Output
// For each test case, output 2 lines, each containing n integers — the desired grid. If there are multiple solutions, output any of them.

// Example
// input
// 3
// 2
// 4
// 6
// output
// 3 2
// 1 4
// 8 2 6 4
// 1 5 3 7
// 11 5 9 1 7 3
// 6 10 2 8 4 12

// Note
// In the first test case, there are only two paths from cell (1,1) to cell (2,2). Their costs are 3−1+4=6 and 3−2+4=5. Then the minimum cost is 5, which is the maximum possible value.
// In the second test case, there are four paths from cell (1,1) to cell (2,4). Their costs are 8−1+5−3+7=16, 8−2+5−3+7=15, 8−2+6−3+7=16, and 8−2+6−4+7=15. Then the minimum value is 15, which is the maximum possible value.

// Hint
// Consider the parity (odd/even) of i+j. What do you observe?

// Editorial
// Observe that ai,j will be added if (i+j) is even, and will be subtracted otherwise. This forms a checkered pattern.
// Obviously, it is optimal for all values that will be added to be strictly larger than all values that will be subtracted. Also, the difference between the value of adjacent grids should be almost equal (by some definition of almost).
// We construct array a as follows:
//  • a1,1=2n−1 and a2,n=2n
//  • For 2≤i≤n and i is even, a1,i=i and a2,i−1=i−1
//  • For 2≤i≤n and i is odd, a1,i=n+i−1 and a2,i−1=n+(i−1)−1
// For example, when n=10, the output will be
// 19 2 12 4 14 6 16 8 18 10
// 1 11 3 13 5 15 7 17 9 20

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, i, j uint
	var ans [][]uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			ans = make([][]uint, 2)
			ans[0] = make([]uint, n)
			ans[1] = make([]uint, n)
			ans[0][0] = 2*n - 1
			ans[1][n-1] = 2 * n

			for i = 1; i < n; i++ {
				if i%2 != 0 {
					ans[0][i] = i + 1
					ans[1][i-1] = i
				} else {
					ans[0][i] = n + i
					ans[1][i-1] = n + i - 1
				}
			}

			for i = 0; i < 2; i++ {
				for j = 0; j < n-1; j++ {
					fmt.Fprint(writer, ans[i][j], " ")
				}
				fmt.Fprintln(writer, ans[i][n-1])
			}
		}
	}

	writer.Flush()
}
