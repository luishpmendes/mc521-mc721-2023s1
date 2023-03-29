// https://codeforces.com/problemset/problem/1686/B

// CodeForces 1686B - Odd Subarrays

// For an array [b1,b2,…,bm] define its number of inversions as the number of pairs (i,j) of integers such that 1≤i<j≤m and bi>bj. Let's call array b odd if its number of inversions is odd.
// For example, array [4,2,7] is odd, as its number of inversions is 1, while array [2,1,4,3] isn't, as its number of inversions is 2.
// You are given a permutation [p1,p2,…,pn] of integers from 1 to n (each of them appears exactly once in the permutation). You want to split it into several consecutive subarrays (maybe just one), so that the number of the odd subarrays among them is as large as possible.
// What largest number of these subarrays may be odd?

// Input
// The first line of the input contains a single integer t (1≤t≤10^5)  — the number of test cases. The description of the test cases follows.
// The first line of each test case contains a single integer n (1≤n≤10^5)  — the size of the permutation.
// The second line of each test case contains n integers p1,p2,…,pn (1≤pi≤n, all pi are distinct)  — the elements of the permutation.
// The sum of n over all test cases doesn't exceed 2⋅10^5.

// Output
// For each test case output a single integer  — the largest possible number of odd subarrays that you can get after splitting the permutation into several consecutive subarrays.

// Example
// input
// 5
// 3
// 1 2 3
// 4
// 4 3 2 1
// 2
// 1 2
// 2
// 2 1
// 6
// 4 5 6 1 2 3
// output
// 0
// 2
// 0
// 1
// 1

// Note
// In the first and third test cases, no matter how we split our permutation, there won't be any odd subarrays.
// In the second test case, we can split our permutation into subarrays [4,3],[2,1], both of which are odd since their numbers of inversions are 1.
// In the fourth test case, we can split our permutation into a single subarray [2,1], which is odd.
// In the fifth test case, we can split our permutation into subarrays [4,5],[6,1,2,3]. The first subarray has 0 inversions, and the second has 3, so it is odd.

// Tutorial
// Consider any optimal splitting.
// Clearly, for any subarray [b1,b2,…,bm] which is not odd, we can just split it into [b1],[b2],…,[bm],
// For any odd subarray [b1,b2,…,bm] with m≥3, there exists an 1≤i≤m−1 such that bi>bi+1 (otherwise b is sorted and has no inversions). Then, we can split b into [b1],[b2],…,[bi−1],[bi,bi+1],[bi+2],…,[bm], where we also have one odd subarray.
// So, if we can split p into several subarrays such that there are k odd subarrays, we can split it into several subarrays of length ≤2 so that there are k odd subarrays too. Then, let dpi denote the largest number of odd subarrays we can get from splitting p[1:i]. Then, dpi=max(dpi−1,dpi−2+(pi−1>pi)). This dp can be calculated in O(n).
// It's also easy to show that the following greedy algorithm works: traverse the permutation from left to right, whenever you see two elements pi−1>pi, make a subarray [pi−1,pi], and proceed from pi+1.

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
	var t, n, i uint
	var p, dp []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			p = make([]uint, n)
			dp = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &p[i])
			}

			dp[0] = 0

			if n > 1 {
				if p[0] > p[1] {
					dp[1] = 1
				} else {
					dp[1] = 0
				}
			}

			for i = 2; i < n; i++ {
				if p[i-1] > p[i] {
					dp[i] = dp[i-2] + 1
				} else {
					dp[i] = dp[i-1]
				}
			}

			fmt.Fprintln(writer, dp[n-1])
		}
	}

	writer.Flush()
}
