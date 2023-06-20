// https://codeforces.com/problemset/problem/1583/B

// CodeForces 1583B - Omkar and Heavenly Tree
// Lord Omkar would like to have a tree with n nodes (3≤n≤10^5) and has asked his disciples to construct the tree. However, Lord Omkar has created m (1≤m<n) restrictions to ensure that the tree will be as heavenly as possible.
// A tree with n nodes is an connected undirected graph with n nodes and n−1 edges. Note that for any two nodes, there is exactly one simple path between them, where a simple path is a path between two nodes that does not contain any node more than once.
// Here is an example of a tree:
// A restriction consists of 3 pairwise distinct integers, a, b, and c (1≤a,b,c≤n). It signifies that node b cannot lie on the simple path between node a and node c.
// Can you help Lord Omkar and become his most trusted disciple? You will need to find heavenly trees for multiple sets of restrictions. It can be shown that a heavenly tree will always exist for any set of restrictions under the given constraints.

// Input
// Each test contains multiple test cases. The first line contains the number of test cases t (1≤t≤10^4). Description of the test cases follows.
// The first line of each test case contains two integers, n and m (3≤n≤10^5, 1≤m<n), representing the size of the tree and the number of restrictions.
// The i-th of the next m lines contains three integers ai, bi, ci (1≤ai,bi,ci≤n, a, b, c are distinct), signifying that node bi cannot lie on the simple path between nodes ai and ci.
// It is guaranteed that the sum of n across all test cases will not exceed 10^5.

// Output
// For each test case, output n−1 lines representing the n−1 edges in the tree. On each line, output two integers u and v (1≤u,v≤n, u≠v) signifying that there is an edge between nodes u and v. Given edges have to form a tree that satisfies Omkar's restrictions.

// Example
// input
// 2
// 7 4
// 1 2 3
// 3 4 5
// 5 6 7
// 6 5 4
// 5 3
// 1 2 3
// 2 3 4
// 3 4 5
// output
// 1 2
// 1 3
// 3 5
// 3 4
// 2 7
// 7 6
// 5 1
// 1 3
// 3 2
// 2 4

// Note
// The output of the first sample case corresponds to the following tree:
// For the first restriction, the simple path between 1 and 3 is 1,3, which doesn't contain 2. The simple path between 3 and 5 is 3,5, which doesn't contain 4. The simple path between 5 and 7 is 5,3,1,2,7, which doesn't contain 6. The simple path between 6 and 4 is 6,7,2,1,3,4, which doesn't contain 5. Thus, this tree meets all of the restrictions.
// The output of the second sample case corresponds to the following tree:

// Tutorial
// Because the number of restrictions is less than n, there is guaranteed to be at least one value from 1 to n that is not a value of b for any of the restrictions. Find a value that is not b for all of the restrictions and construct a tree that is a "star" with that value in the middle. An easy way to do this is to make an edge from that value to every other number from 1 to n.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var q int
	var t, n, m, a, b, c, i, middle uint
	var set map[uint]struct{}

	for q, _ = fmt.Fscan(reader, &t); q == 1; q, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &m)
			set = make(map[uint]struct{})

			for i = 0; i < m; i++ {
				fmt.Fscan(reader, &a, &b, &c)
				set[b] = struct{}{}
			}

			middle = 0

			for i = 1; i <= n && middle == 0; i++ {
				if _, ok := set[i]; !ok {
					middle = i
				}
			}

			for i = 1; i <= n; i++ {
				if i != middle {
					fmt.Fprintln(writer, middle, i)
				}
			}
		}
	}

	writer.Flush()
}
