// https://codeforces.com/problemset/problem/1270/G

// CodeForces 1270G - Subset with Zero Sum

// You are given n integers a1,a2,…,an, such that for each 1≤i≤n holds i−n≤ai≤i−1.
// Find some nonempty subset of these integers, whose sum is equal to 0. It can be shown that such a subset exists under given constraints. If there are several possible subsets with zero-sum, you can find any of them.

// Input
// Each test contains multiple test cases. The first line contains the number of test cases t (1≤t≤10^6). The description of the test cases follows.
// The first line of each test case contains a single integer n (1≤n≤10^6).
// The second line of each test case contains n integers a1,a2,…,an (i−n≤ai≤i−1).
// It is guaranteed that the sum of n over all test cases does not exceed 10^6.

// Output
// For each test case, output two lines.
// In the first line, output s (1≤s≤n) — the number of elements in your subset.
// In the second line, output s integers i1,i2,…,is (1≤ik≤n). All integers have to be pairwise different, and ai1+ai2+⋯+ais has to be equal to 0. If there are several possible subsets with zero-sum, you can find any of them.

// Example
// input
// 2
// 5
// 0 1 2 3 4
// 4
// -3 1 1 1
// output
// 1
// 1
// 4
// 1 4 3 2

// Note
// In the first example, we get sum is a1=0.
// In the second example, we get sum is a1+a4+a3+a2=0.

// Tutorial
// Note that the condition i−n≤ai≤i−1 is equivalent to 1≤i−ai≤n.
// Let's build an oriented graph G on n nodes with the following principle: for each i from 1 to n, draw an oriented edge from vertex i to vertex i−ai. In this graph, there is an outgoing edge from every vertex, so it has an oriented cycle. Let vertices of this cycle be — i1,i2,…,ik.
// Then:
// i1−ai1=i2
// i2−ai2=i3
// ⋮
// in−ain=i1
// After adding all these equalities, we get ai1+ai2+⋯+aik=0
// We can find some oriented cycle in O(n) (just go by an available edge until you get to previously visited vertex).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, v int
	var t, n, i uint
	var a []int
	var used []bool
	var ind []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]int, n)
			used = make([]bool, n)
			ind = make([]uint, 0, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
				used[i] = false
			}

			v = 0

			for !used[v] {
				used[v] = true
				ind = append(ind, uint(v))
				v -= a[v]
			}

			for i = 0; ind[i] != uint(v); i++ {
			}

			fmt.Fprintln(writer, len(ind)-int(i))

			for ; i < uint(len(ind)); i++ {
				fmt.Fprint(writer, ind[i]+1)

				if i+1 < uint(len(ind)) {
					fmt.Fprint(writer, " ")
				}
			}

			fmt.Fprintln(writer)
		}
	}

	writer.Flush()
}
