// https://codeforces.com/problemset/problem/1398/C

// CodeForces 1398C - Good Subarrays

// You are given an array a1,a2,…,an consisting of integers from 0 to 9. A subarray al,al+1,al+2,…,ar−1,ar is good if the sum of elements of this subarray is equal to the length of this subarray (∑_{i=l}^{r}{ai}=r−l+1).
// For example, if a=[1,2,0], then there are 3 good subarrays: a1…1=[1],a2…3=[2,0] and a1…3=[1,2,0].
// Calculate the number of good subarrays of the array a.

// Input
// The first line contains one integer t (1≤t≤1000) — the number of test cases.
// The first line of each test case contains one integer n (1≤n≤10^5) — the length of the array a.
// The second line of each test case contains a string consisting of n decimal digits, where the i-th digit is equal to the value of ai.
// It is guaranteed that the sum of n over all test cases does not exceed 10^5.

// Output
// For each test case print one integer — the number of good subarrays of the array a.

// Example
// input
// 3
// 3
// 120
// 5
// 11011
// 6
// 600005
// output
// 3
// 6
// 1

// Note
// The first test case is considered in the statement.
// In the second test case, there are 6 good subarrays: a1…1, a2…2, a1…2, a4…4, a5…5 and a4…5.
// In the third test case there is only one good subarray: a2…6.

// Tutorial
// We use zero indexing in this solution. We also use half-closed interval (so subarray [l,r] is al,al+1,…,ar−1).
// Let's precalculate the array p, where pi=∑_{j=0}^{i−1}{aj} (so px if sum of first x elements of a).
// Then subarray [l,r] is good if pr−pl=r−l, so pr−r=pl−l.
// Thus, we have to group all prefix by value pi−i for i from 0 to n. And if the have x prefix with same value of pi−i then we have to add x(x−1)/2 to the answer.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, i, ans, aux, k uint
	var a string
	var dp []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &a)
			dp = make([]uint, n)
			dp[0] = uint(a[0] - '0')

			for i = 1; i < n; i++ {
				dp[i] = dp[i-1] + uint(a[i]-'0')
			}

			ans = 0

			for i = 0; i < n; i++ {
				dp[i] -= i

				if dp[i] == 1 {
					ans++
				}
			}

			sort.Slice(dp, func(i, j int) bool {
				return dp[i] < dp[j]
			})
			k = 1
			aux = 0

			for i = 1; i < n; i++ {
				if dp[i] == dp[i-1] {
					k++
				} else {
					aux += (k * (k - 1))
					k = 1
				}
			}

			aux += (k * (k - 1))
			ans += aux >> 1
			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
