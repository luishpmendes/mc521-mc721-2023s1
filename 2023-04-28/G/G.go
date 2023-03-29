// https://codeforces.com/problemset/problem/1469/B

// CodeForces 1469B - Red and Blue

// Monocarp had a sequence a consisting of n+m integers a1,a2,…,an+m. He painted the elements into two colors, red and blue; n elements were painted red, all other m elements were painted blue.
// After painting the elements, he has written two sequences r1,r2,…,rn and b1,b2,…,bm. The sequence r consisted of all red elements of a in the order they appeared in a; similarly, the sequence b consisted of all blue elements of a in the order they appeared in a as well.
// Unfortunately, the original sequence was lost, and Monocarp only has the sequences r and b. He wants to restore the original sequence. In case there are multiple ways to restore it, he wants to choose a way to restore that maximizes the value of
// f(a)=max(0,a1,(a1+a2),(a1+a2+a3),…,(a1+a2+a3+⋯+an+m))
// Help Monocarp to calculate the maximum possible value of f(a).

// Input
// The first line contains one integer t (1≤t≤1000) — the number of test cases. Then the test cases follow. Each test case consists of four lines.
// The first line of each test case contains one integer n (1≤n≤100).
// The second line contains n integers r1,r2,…,rn (−100≤ri≤100).
// The third line contains one integer m (1≤m≤100).
// The fourth line contains m integers b1,b2,…,bm (−100≤bi≤100).

// Output
// For each test case, print one integer — the maximum possible value of f(a).

// Example
// input
// 4
// 4
// 6 -5 7 -3
// 3
// 2 3 -4
// 2
// 1 1
// 4
// 10 -3 2 2
// 5
// -1 -2 -3 -4 -5
// 5
// -1 -2 -3 -4 -5
// 1
// 0
// 1
// 0
// output
// 13
// 13
// 0
// 0

// Note
// In the explanations for the sample test cases, red elements are marked with "".
// In the first test case, one of the possible sequences a is ["6",2,"−5",3,"7","−3",−4].
// In the second test case, one of the possible sequences a is [10,"1",−3,"1",2,2].
// In the third test case, one of the possible sequences a is ["−1",−1,−2,−3,"−2",−4,−5,"−3","−4","−5"].
// In the fourth test case, one of the possible sequences a is [0,"0"].

// Tutorial
// Denote pi as the sum of first i elements of r, and qj as the sum of first j elements of b. These values can be calculated in O(n+m) with prefix sums.
// The first solution is to use dynamic programming. Let dpi,j be the maximum value of f(a) if we placed the first i elements of r and the first j elements of b. Transitions can be performed in O(1): we either place an element from r (then we go to dpi+1,j and update it with max(dpi,j,pi+1+qj)), or place an element from b (then we go to dpi,j+1 and update it with max(dpi,j,pi+qj+1)). The answer is stored in dpn,m, and this solution works in O(nm).
// The second solution: observe that the sum of several first elements of a is the sum of several first elements of r and several first elements of b. So each prefix sum of a (and the answer itself) is not greater than maxni=0pi+maxmj=0pj. It's easy to show how to obtain exactly this answer: let k be the value of i such that pi is maximized, and l be the value of j such that qj is maximized. Let's place the first k elements of r, then the first l elements of b (so the current sum is exactly maxni=0pi+maxmj=0pj), and place all of the remaining elements in any possible order. So, the answer is maxni=0pi+maxmj=0pj. This solution works in O(n+m).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, ans int
	var t, n, m, i, j uint
	var r, b []int
	var dp [][]int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			r = make([]int, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &r[i])
			}

			fmt.Fscan(reader, &m)
			b = make([]int, m)

			for i = 0; i < m; i++ {
				fmt.Fscan(reader, &b[i])
			}

			dp = make([][]int, n+1)

			for i = 0; i <= n; i++ {
				dp[i] = make([]int, m+1)

				for j = 0; j <= m; j++ {
					dp[i][j] = -1e9
				}
			}

			dp[0][0] = 0
			ans = 0

			for i = 0; i <= n; i++ {
				for j = 0; j <= m; j++ {
					if i < n {
						if dp[i+1][j] < dp[i][j]+r[i] {
							dp[i+1][j] = dp[i][j] + r[i]
						}
					}

					if j < m {
						if dp[i][j+1] < dp[i][j]+b[j] {
							dp[i][j+1] = dp[i][j] + b[j]
						}
					}

					if ans < dp[i][j] {
						ans = dp[i][j]
					}
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
