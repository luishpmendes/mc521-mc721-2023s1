// https://codeforces.com/problemset/problem/1380/A

// CodeForces 1380A - Three Indices

// You are given a permutation p1,p2,…,pn. Recall that sequence of n integers is called a permutation if it contains all integers from 1 to n exactly once.
// Find three indices i, j and k such that:
//  • 1≤i<j<k≤n;
//  • pi<pj and pj>pk.
// Or say that there are no such indices.

// Input
// The first line contains a single integer T (1≤T≤200) — the number of test cases.
// Next 2T lines contain test cases — two lines per test case. The first line of each test case contains the single integer n (3≤n≤1000) — the length of the permutation p.
// The second line contains n integers p1,p2,…,pn (1≤pi≤n; pi≠pj if i≠j) — the permutation p.

// Output
// For each test case:
//  • if there are such indices i, j	and k, print YES (case insensitive) and the indices themselves;
//  • if there are no such indices, print NO (case insensitive).
// If there are multiple valid triples of indices, print any of them.

// Example
// input
// 3
// 4
// 2 1 4 3
// 6
// 4 6 1 2 5 3
// 5
// 5 3 1 2 4
// output
// YES
// 2 3 4
// YES
// 3 5 6
// NO

// Tutorial
// A solution in O(n2): iterate on j, check that there exists an element lower than aj to the left of it, and check that there exists an element lower than aj to the right of it. Can be optimized to O(n) with prefix/suffix minima.
// A solution in O(n): note that if there is some answer, we can find an index j such that aj−1<aj and aj>aj+1 (if there is no such triple, the array descends to some point and ascends after that, so there is no answer). So we only have to check n−2 consecutive triples.

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
	var p []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			p = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &p[i])
			}

			for i = 1; i+1 < n; i++ {
				if p[i-1] < p[i] && p[i] > p[i+1] {
					fmt.Fprintln(writer, "YES")
					fmt.Fprintln(writer, i, i+1, i+2)
					break
				}
			}

			if i+1 >= n {
				fmt.Fprintln(writer, "NO")
			}
		}
	}

	writer.Flush()
}
