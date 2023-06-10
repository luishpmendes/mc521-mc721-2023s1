// https://codeforces.com/problemset/problem/1799/A

// CodeForces 1799A - Recent Actions

// On Codeforces the "Recent Actions" field shows the last n posts with recent actions.
// Initially, there are posts 1,2,…,n in the field (this is in order from top to down). Also there are infinitely many posts not in the field, numbered with integers n+1,n+2,….
// When recent action happens in the post p:
//  • If it is in the "Recent Actions" field, it moves from its position to the top position.
//  • Otherwise, it is added to the top position, and the post on the down position is removed from the "Recent Actions" field.
// You know, that the next m recent actions will happen in the posts p1,p2,…,pm (n+1≤pi≤n+m) in the moments of time 1,2,…,m. Note, that recent actions only happen with posts with numbers ≥n+1.
// For each post i (1≤i≤n), find the first time it will be removed from the "Recent Actions" field or say, that it won't be removed.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of test cases. Descriptions of test cases follow.
// The first line of each test case contains two integers n, m (1≤n,m≤5⋅10^4) — the size of the "Recent Actions" field and the number of actions.
// The next line contains m integers p1,p2,…,pm (n+1≤pi≤n+m).
// It is guaranteed, that the sum of n and the sum of m for all test cases does not exceed 5⋅10^4.

// Output
// For each test case print n integers t1,t2,…,tn, where ti=−1 if the post i won't be removed or ti equals to the first moment of time the post i will be removed (1≤ti≤m).

// Example
// input
// 10
// 1 1
// 2
// 3 2
// 5 4
// 4 5
// 5 9 9 5 7
// 5 5
// 6 7 8 9 10
// 3 4
// 4 4 4 4
// 4 4
// 5 5 6 6
// 3 5
// 4 5 5 5 4
// 4 20
// 5 5 24 24 24 5 6 7 8 9 10 12 13 14 15 16 17 18 19 20
// 5 7
// 7 8 7 11 7 12 10
// 6 7
// 8 11 7 8 8 8 12
// output
// 1
// -1 2 1
// -1 5 2 1
// 5 4 3 2 1
// -1 -1 1
// -1 -1 3 1
// -1 2 1
// 8 7 3 1
// 7 6 4 2 1
// -1 -1 7 3 2 1

// Note
// In the first test case, the only post 1 will be removed at the moment 1 and replaced by the post 2.
// In the second test case the "Recent Actions" field will be (given an order from top to down):
//  1. Before moment 1: [1,2,3], after moment 1: [5,1,2]. Post number 3 was removed.
//  2. Before moment 2: [5,1,2], after moment 2: [4,5,1]. Post number 2 was removed.
// Post number 1 won't be removed.
// In the third test case the "Recent Actions" field will be (given an order from top to down):
//  1. Before moment 1: [1,2,3,4], after moment 1: [5,1,2,3]. Post number 4 was removed.
//  2. Before moment 2: [5,1,2,3], after moment 2: [9,5,1,2]. Post number 3 was removed.
//  3. Before moment 3: [9,5,1,2], after moment 3: [9,5,1,2]. Nothing was changed.
//  4. Before moment 4: [9,5,1,2], after moment 4: [5,9,1,2]. The order was changed.
//  5. Before moment 5: [5,9,1,2], after moment 5: [7,5,9,1]. Post number 2 was removed.
// Post number 1 won't be removed.

// Tutorial
// Note, that posts will be removed in the order n,n−1,…,1.
// The post n−k+1 will be removed at the first time, when there are at least k different numbers among p1,p2,…,pi.
// So let's calculate the number of different numbers among p1,p2,…,pi for each i using boolean array of length m iterating i from 1 to m. Using them we can calculate the answer.
// Time complexity: O(n+m).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, n, m, p, i, j int
	var v []bool
	var ans []int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &m)
			v = make([]bool, n+m)
			ans = make([]int, n)
			j = n

			for i = 0; i < n; i++ {
				ans[i] = -1
			}

			for i = 0; i < m; i++ {
				fmt.Fscan(reader, &p)
				p--

				if !v[p] && j > 0 {
					ans[j-1] = i + 1
					j--
				}

				v[p] = true
			}

			for i = 0; i+1 < n; i++ {
				fmt.Fprint(writer, ans[i], " ")
			}

			fmt.Fprintln(writer, ans[n-1])
		}
	}

	writer.Flush()
}
