// https://codeforces.com/problemset/problem/1538/C

// CodeForces 1538C - Number of Pairs

// You are given an array a of n integers. Find the number of pairs (i,j) (1≤i<j≤n) where the sum of ai+aj is greater than or equal to l and less than or equal to r (that is, l≤ai+aj≤r).
// For example, if n=3, a=[5,1,2], l=4 and r=7, then two pairs are suitable:
//  • i=1 and j=2 (4≤5+1≤7);
//  • i=1 and j=3 (4≤5+2≤7).

// Input
// The first line contains an integer t (1≤t≤10^4). Then t test cases follow.
// The first line of each test case contains three integers n,l,r (1≤n≤2⋅10^5, 1≤l≤r≤10^9) — the length of the array and the limits on the sum in the pair.
// The second line contains n integers a1,a2,…,an (1≤ai≤10^9).
// It is guaranteed that the sum of n overall test cases does not exceed 2⋅10^5.

// Output
// For each test case, output a single integer — the number of index pairs (i,j) (i<j), such that l≤ai+aj≤r.

// Example
// input
// 4
// 3 4 7
// 5 1 2
// 5 5 8
// 5 1 2 4 3
// 4 100 1000
// 1 1 1 1
// 5 9 13
// 2 5 5 1 1
// output
// 2
// 7
// 0
// 1

// Tutorial
// The problem can be divided into two classic ones:
//  • Count the number of pairs ai+aj≤r;
//  • Count the number of pairs ai+aj≤l−1.
// Let A — be the answer to the first problem, and B — be the answer to the second problem. Then A−B is the answer to the original problem.
// The new problem can be solved by binary search. Iterate over the first element of the pair. Then you need to count the number of elements such that aj≤r−ai. If you sort the array, this value can be calculated by running a single binary search.

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
	var t, n, l, r, i, ans, left, right uint
	var a []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &l, &r)
			a = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
			}

			sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

			ans = 0
			left = 0
			right = 0

			for i = n; i > 0; i-- {
				for right < n && a[right]+a[i-1] <= r {
					right++
				}

				for left < n && a[left]+a[i-1] < l {
					left++
				}

				ans += right - left

				if i >= left+1 && i < right+1 {
					ans--
				}
			}

			fmt.Fprintln(writer, ans/2)
		}
	}

	writer.Flush()
}
