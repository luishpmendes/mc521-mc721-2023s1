// https://codeforces.com/problemset/problem/1771/B

// CodeForces 1771B - Hossam and Friends

// Hossam makes a big party, and he will invite his friends to the party.
// He has n friends numbered from 1 to n. They will be arranged in a queue as follows: 1,2,3,…,n.
// Hossam has a list of m pairs of his friends that don't know each other. Any pair not present in this list are friends.
// A subsegment of the queue starting from the friend a and ending at the friend b is [a,a+1,a+2,…,b]. A subsegment of the queue is called good when all pairs of that segment are friends.
// Hossam wants to know how many pairs (a,b) there are (1≤a≤b≤n), such that the subsegment starting from the friend a and ending at the friend b is good.

// Input
// The input consists of multiple test cases. The first line contains a single integer t (1≤t≤2⋅104), the number of test cases. Description of the test cases follows.
// The first line of each test case contains two integer numbers n and m (2≤n≤105, 0≤m≤105) representing the number of friends and the number of pairs, respectively.
// The i-th of the next m lines contains two integers xi and yi (1≤xi,yi≤n, xi≠yi) representing a pair of Hossam's friends that don't know each other.
// Note that pairs can be repeated.
// It is guaranteed that the sum of n over all test cases does not exceed 105, and the sum of m over all test cases does not exceed 105.

// Output
// For each test case print an integer — the number of good subsegments.

// Example
// input
// 2
// 3 2
// 1 3
// 2 3
// 4 2
// 1 2
// 2 3
// output
// 4
// 5

// Note
// In the first example, the answer is 4.
// The good subsegments are:
// [1]
// [2]
// [3]
// [1, 2]
// In the second example, the answer is 5.
// The good subsegments are:
// [1]
// [2]
// [3]
// [4]
// [3, 4]

// Tutorial
// Just ai<bi in non-friends pairs. Let's calculate ri= minimum non-friend for all people. So, we can't start subsegment in ai and finish it righter ri.
// Let's process people from right to left and calculate the rightmost positions there subsegment can end. Initially, R=n−1. Then we go to ai just do R=min(R,ri) and add R−i+1 to answer.

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
	var t, n, m, i, x, y, ans uint
	var mn []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &m)
			mn = make([]uint, n)

			for i = 0; i < n; i++ {
				mn[i] = n
			}

			for i = 0; i < m; i++ {
				fmt.Fscan(reader, &x, &y)

				if x > y {
					x, y = y, x
				}

				if mn[x-1] > y-1 {
					mn[x-1] = y - 1
				}
			}

			for i = n - 1; i > 0; i-- {
				if mn[i-1] > mn[i] {
					mn[i-1] = mn[i]
				}
			}

			ans = n
			for i = 0; i < n; i++ {
				ans += (mn[i] - i - 1)
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
