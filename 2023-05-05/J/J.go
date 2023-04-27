// https://codeforces.com/problemset/problem/1675/G

// CodeForces 1675G - Sorting Pancakes

// Nastya baked m pancakes and spread them on n dishes. The dishes are in a row and numbered from left to right. She put ai pancakes on the dish with the index i.
// Seeing the dishes, Vlad decided to bring order to the stacks and move some pancakes. In one move, he can shift one pancake from any dish to the closest one, that is, select the dish i (ai>0) and do one of the following:
//  • if i>1, put the pancake on a dish with the previous index, after this move ai=ai−1 and ai−1=ai−1+1;
//  • if i<n, put the pancake on a dish with the following index, after this move ai=ai−1 and ai+1=ai+1+1.
// Vlad wants to make the array anon-increasing, after moving as few pancakes as possible. Help him find the minimum number of moves needed for this.
// The array a=[a1,a2,…,an] is called non-increasing if ai≥ai+1 for all i from 1 to n−1.

// Input
// The first line of the input contains two numbers n and m (1≤n,m≤250) — the number of dishes and the number of pancakes, respectively.
// The second line contains n numbers ai (0≤ai≤m), the sum of all ai is equal to m.

// Output
// Print a single number: the minimum number of moves required to make the array a non-increasing.

// Examples

// input
// 6 19
// 5 3 2 3 3 3
// output
// 2

// input
// 3 6
// 3 2 1
// output
// 0

// input
// 3 6
// 2 1 3
// output
// 1

// input
// 6 19
// 3 4 3 3 5 1
// output
// 3

// input
// 10 1
// 0 0 0 0 0 0 0 0 0 1
// output
// 9

// Note
// In the first example, you first need to move the pancake from the dish 1, after which a=[4,4,2,3,3,3]. After that, you need to move the pancake from the dish 2 to the dish 3 and the array will become non-growing a=[4,3,3,3,3,3].

// Tutorial
// For convenience, we will calculate the prefix sums on the array a, we will also enter the array b containing the indexes of all pancakes and calculate the prefix sums on it.
// Let's use dynamic programming. Let's define dp[i][last][sum] as the required number of operations to correctly lay out the i-th prefix, with the final ai=last, and ∑ij=1aj=sum. Then you can go to dp[i][last][sum] from dp[i−1][last+j][sum−last] (the previous number must be greater, and the sum is fixed). To dp[i−1][last+j][sum−last], it will be necessary to add a certain number of actions necessary to get ai=last, let's call it add (all the terrible prefix sums are needed to count it). Since add depends only on last and sum, we only need to choose the minimum dp[i−1][last+j][sum−last], the choice can be optimized by suffix minima. As a result, the solution works for O(n∗m2), that's how many states need to be processed.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, m, i, j, k uint
	var a []uint
	var dp [][]uint

	for t, _ = fmt.Fscan(reader, &n, &m); t == 2; t, _ = fmt.Fscan(reader, &n, &m) {
		a = make([]uint, n)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])

			if i > 0 {
				a[i] += a[i-1]
			}
		}

		dp = make([][]uint, n+1)

		for i = 0; i <= n; i++ {
			dp[i] = make([]uint, m+1)

			for j = 0; j <= m; j++ {
				dp[i][j] = math.MaxUint32
			}
		}

		dp[0][0] = 0

		for j = m + 1; j > 0; j-- {
			for i = 1; i <= n; i++ {
				for k = j - 1; k <= m; k++ {
					if k >= a[i-1] {
						if dp[i][k] > dp[i-1][k-j+1]+k-a[i-1] {
							dp[i][k] = dp[i-1][k-j+1] + k - a[i-1]
						}
					} else {
						if dp[i][k] > dp[i-1][k-j+1]+a[i-1]-k {
							dp[i][k] = dp[i-1][k-j+1] + a[i-1] - k
						}
					}
				}
			}
		}

		fmt.Fprintln(writer, dp[n][m])
	}

	writer.Flush()
}
