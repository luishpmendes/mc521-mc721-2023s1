// https://codeforces.com/problemset/problem/1741/D

// CodeForces 1741D - Masha and a Beautiful Tree

// The girl named Masha was walking in the forest and found a complete binary tree of height n and a permutation p of length m=2^n.
// A complete binary tree of height n is a rooted tree such that every vertex except the leaves has exactly two sons, and the length of the path from the root to any of the leaves is n. The picture below shows the complete binary tree for n=2.
// A permutation is an array consisting of n different integers from 1 to n. For example, [2,3,1,5,4] is a permutation, but [1,2,2] is not (2 occurs twice), and [1,3,4] is also not a permutation (n=3, but there is 4 in the array).
// Let's enumerate m leaves of this tree from left to right. The leaf with the number i contains the value pi (1≤i≤m).
// For example, if n=2, p=[3,1,4,2], the tree will look like this:
// Masha considers a tree beautiful if the values in its leaves are ordered from left to right in increasing order.
// In one operation, Masha can choose any non-leaf vertex of the tree and swap its left and right sons (along with their subtrees).
// For example, if Masha applies this operation to the root of the tree discussed above, it will take the following form:
// Help Masha understand if she can make a tree beautiful in a certain number of operations. If she can, then output the minimum number of operations to make the tree beautiful.

// Input
// The first line contains single integer t (1≤t≤10^4) — number of test cases.
// In each test case, the first line contains an integer m (1≤m≤262144), which is a power of two  — the size of the permutation p.
// The second line contains m integers: p1,p2,…,pm (1≤pi≤m) — the permutation p.
// It is guaranteed that the sum of m over all test cases does not exceed 3⋅10^5.

// Output
// For each test case in a separate line, print the minimum possible number of operations for which Masha will be able to make the tree beautiful or -1, if this is not possible.

// Example
// input
// 4
// 8
// 6 5 7 8 4 3 1 2
// 4
// 3 1 4 2
// 1
// 1
// 8
// 7 8 4 3 1 2 6 5
// output
// 4
// -1
// 0
// -1

// Note
// Consider the first test.
// In the first test case, you can act like this (the vertex to which the operation is applied at the current step is highlighted in purple):
// It can be shown that it is impossible to make a tree beautiful in fewer operations.
// In the second test case, it can be shown that it is impossible to make a tree beautiful.
// In the third test case, the tree is already beautiful.

// Tutorial
// Let some vertex be responsible for a segment of leaves [l..r]. Then her left son is responsible for the segment [l..(l+r−1)/2], and the right for the segment [(l+r+1)/2..r]. Note that if we do not apply the operation to this vertex, then it will not be possible to move some element from the right son's segment to the left son's segment. It remains to understand when we need to apply the operation to the vertex. Let the maximum on the segment [l..r] be max, the minimum on the same segment is min. Then if min lies in the right son, and max in the left, then we should obviously apply the operation, for the reason described above. In the case when min lies in the left son, and max in the right, the application of the operation will definitely not allow you to get a solution. Let's continue to act in a similar way recursively from the children of the current vertex. At the end, we should check whether we have received a sorted permutation. The above solution works for O(nm), since there are n levels in the tree and at each level, vertexes are responsible for m sheets in total. You can optimize this solution to O(m) if you pre-calculate the maximum and minimum for each vertex.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve(p *[]uint, l, r uint) uint {
	var mid, i, mal, mar, ans uint

	ans = 0

	if r != l+1 {
		mid = (l + r) >> 1
		mal = (*p)[l]
		mar = (*p)[mid]
		ans = 0

		for i = l + 1; i < mid; i++ {
			if mal < (*p)[i] {
				mal = (*p)[i]
			}
		}

		for i = mid + 1; i < r; i++ {
			if mar < (*p)[i] {
				mar = (*p)[i]
			}
		}

		if mal > mar {
			ans++

			for i = 0; i+l < mid; i++ {
				(*p)[l+i], (*p)[mid+i] = (*p)[mid+i], (*p)[l+i]
			}
		}

		ans += solve(p, l, mid)
		ans += solve(p, mid, r)
	}

	return ans
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, m, i, ans uint
	var p []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &m)
			p = make([]uint, m)

			for i = 0; i < m; i++ {
				fmt.Fscan(reader, &p[i])
			}

			ans = solve(&p, 0, m)

			if sort.SliceIsSorted(p, func(i, j int) bool { return p[i] < p[j] }) {
				fmt.Fprintln(writer, ans)
			} else {
				fmt.Fprintln(writer, -1)
			}
		}
	}

	writer.Flush()
}
