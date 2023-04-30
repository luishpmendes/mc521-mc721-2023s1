// https://codeforces.com/problemset/problem/1795/B

// CodeForces 1795B - Ideal Point

// You are given n one-dimensional segments (each segment is denoted by two integers — its endpoints).
// Let's define the function f(x) as the number of segments covering point x (a segment covers the point x if l≤x≤r, where l is the left endpoint and r is the right endpoint of the segment).
// An integer point x is called ideal if it belongs to more segments than any other integer point, i. e. f(y)<f(x) is true for any other integer point y.
// You are given an integer k. Your task is to determine whether it is possible to remove some (possibly zero) segments, so that the given point k becomes ideal.

// Input
// The first line contains one integer t (1≤t≤1000) — the number of test cases.
// The first line of each test case contains two integers n and k (1≤n,k≤50).
// Then n lines follow, i-th line of them contains two integers li and ri (1≤li,ri≤50; li≤ri) — the endpoints of the i-th segment.

// Output
// For each test case, print YES if it is possible to remove some (possibly zero) segments, so that the given point k becomes ideal, otherwise print NO.
// You may print each letter in any case (YES, yes, Yes will all be recognized as positive answer, NO, no and nO will all be recognized as negative answer).

// Example
// input
// 4
// 4 3
// 1 3
// 7 9
// 2 5
// 3 6
// 2 9
// 1 4
// 3 7
// 1 3
// 2 4
// 3 5
// 1 4
// 6 7
// 5 5
// output
// YES
// NO
// NO
// YES

// Note
// In the first example, the point 3 is already ideal (it is covered by three segments), so you don't have to delete anything.
// In the fourth example, you can delete everything except the segment [5,5].

// Tutorial
// First of all, let's delete all segments that do not cover the point k (because they increase the value of the function f at points other than k). If there are no segments left, then the answer is NO. Otherwise, all segments cover the point k. And it remains to check whether the point k is the only point which is covered by all segments. Note that it does not make sense to delete any of the remaining segments, because if there are several points with maximum value of f, then deleting segments can only increase their number.
// To check the number of points with the maximum value of f, you can iterate over x from 1 to 50 and calculate f(x), because of the small number of segments in the problem. A faster way is to check the size of the intersection of all segments. The left boundary of the intersection is L=maxi=1nli, and the right boundary is R=mini=1nri; if L=R, then the point k is ideal, otherwise it is not.

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
	var t, n, k, L, R, l, r uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &k)
			L = 0
			R = 55

			for ; n > 0; n-- {
				fmt.Fscan(reader, &l, &r)

				if l <= k && k <= r {
					if L < l {
						L = l
					}

					if R > r {
						R = r
					}
				}
			}

			if L == R {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}

	writer.Flush()
}
