// https://codeforces.com/problemset/problem/1764/C

// CodeForces 1764C - Doremy's City Construction

// Doremy's new city is under construction! The city can be regarded as a simple undirected graph with n vertices. The i-th vertex has altitude ai. Now Doremy is deciding which pairs of vertices should be connected with edges.
// Due to economic reasons, there should be no self-loops or multiple edges in the graph.
// Due to safety reasons, there should not be pairwise distinct vertices u, v, and w such that au≤av≤aw and the edges (u,v) and (v,w) exist.
// Under these constraints, Doremy would like to know the maximum possible number of edges in the graph. Can you help her?
// Note that the constructed graph is allowed to be disconnected.

// Input
// The input consists of multiple test cases. The first line contains a single integer t (1≤t≤10^4) — the number of test cases. The description of the test cases follows.
// The first line of each test case contains a single integer n (2≤n≤2⋅10^5) — the number of vertices.
// The second line of each test case contains n integers a1,a2,…,an (1≤ai≤10^6) — the altitudes of each vertex.
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅10^5.

// Output
// For each test case, output the maximum possible number of edges in the graph.

// Example
// input
// 4
// 4
// 2 2 3 1
// 6
// 5 2 3 1 5 2
// 12
// 7 2 4 9 1 4 6 3 7 4 2 3
// 4
// 1000000 1000000 1000000 1000000
// output
// 3
// 9
// 35
// 2

// Note
// In the first test case, there can only be at most 3 edges in the graph. A possible construction is to connect (1,3), (2,3), (3,4). In the picture below the red number above node i is ai.
// The following list shows all such u, v, w that the edges (u,v) and (v,w) exist.
//  • u=1, v=3, w=2;
//  • u=1, v=3, w=4;
//  • u=2, v=3, w=1;
//  • u=2, v=3, w=4;
//  • u=4, v=3, w=1;
//  • u=4, v=3, w=2.
// Another possible construction is to connect (1,4), (2,4), (3,4).
// An unacceptable construction is to connect (1,3), (2,3), (2,4), (3,4). Because when u=4, v=2, w=3, au≤av≤aw holds, and the respective edges exist.

// Tutorial
// We can first assume that no edge links two vertices with the same value, then for each vertex u and his neighbors Su, either au>maxv∈Suav or au<minv∈Suav. If au>maxv∈Suav, we paint u black, otherwise we paint it white. Then it's obviously that any edge connects two vertices with different color.
// So we can first determine the color of each vertex and then add as many as edges according to the color. (u,v) is addable only when u is black, v is white and au>av. If we paint u black and v white and au<av, we can swap the colors and the edges connecting to these 2 vertices and the answer will be no worse. So the best painting plan is determining a value A, painting u black if and only if au≥A, painting u white if and only if au<A.
// If we add an edge linking two vertices with the same value, then the two vertices can not connect other vertices. This plan will only be considered when a1=a2=⋯=an. When a1=a2=⋯=an, the answer is ⌊n/2⌋.

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
	var t, n, i, ans uint
	var a []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
			}

			sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
			ans = n / 2

			for i = 0; i+1 < n; i++ {
				if a[i] != a[i+1] {
					if ans < (i+1)*(n-i-1) {
						ans = (i + 1) * (n - i - 1)
					}
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
