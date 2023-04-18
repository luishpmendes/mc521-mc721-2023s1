// https://codeforces.com/problemset/problem/1189/C

// CodeForces 1189C - Candies!

// Consider a sequence of digits of length 2k [a1,a2,…,a2k]. We perform the following operation with it: replace pairs (a2i+1,a2i+2) with (a2i+1+a2i+2)mod10 for 0≤i<2k−1. For every i where a2i+1+a2i+2≥10 we get a candy! As a result, we will get a sequence of length 2k−1.
// Less formally, we partition sequence of length 2k into 2k−1 pairs, each consisting of 2 numbers: the first pair consists of the first and second numbers, the second of the third and fourth …, the last pair consists of the (2k−1)-th and (2k)-th numbers. For every pair such that sum of numbers in it is at least 10, we get a candy. After that, we replace every pair of numbers with a remainder of the division of their sum by 10 (and don't change the order of the numbers).
// Perform this operation with a resulting array until it becomes of length 1. Let f([a1,a2,…,a2k]) denote the number of candies we get in this process.
// For example: if the starting sequence is [8,7,3,1,7,0,9,4] then:
// After the first operation the sequence becomes [(8+7)mod10,(3+1)mod10,(7+0)mod10,(9+4)mod10] = [5,4,7,3], and we get 2 candies as 8+7≥10 and 9+4≥10.
// After the second operation the sequence becomes [(5+4)mod10,(7+3)mod10] = [9,0], and we get one more candy as 7+3≥10.
// After the final operation sequence becomes [(9+0)mod10] = [9].
// Therefore, f([8,7,3,1,7,0,9,4])=3 as we got 3 candies in total.
// You are given a sequence of digits of length n s1,s2,…sn. You have to answer q queries of the form (li,ri), where for i-th query you have to output f([sli,sli+1,…,sri]). It is guaranteed that ri−li+1 is of form 2k for some nonnegative integer k.

// Input
// The first line contains a single integer n (1≤n≤10^5) — the length of the sequence.
// The second line contains n digits s1,s2,…,sn (0≤si≤9).
// The third line contains a single integer q (1≤q≤10^5) — the number of queries.
// Each of the next q lines contains two integers li, ri (1≤li≤ri≤n) — i-th query. It is guaranteed that ri−li+1 is a nonnegative integer power of 2.

// Output
// Output q lines, in i-th line output single integer — f([sli,sli+1,…,sri]), answer to the i-th query.

// Examples

// input
// 8
// 8 7 3 1 7 0 9 4
// 3
// 1 8
// 2 5
// 7 7
// output
// 3
// 1
// 0

// input
// 6
// 0 1 2 3 3 5
// 3
// 1 2
// 1 4
// 3 6
// output
// 0
// 0
// 1

// Note
// The first example illustrates an example from the statement.
// f([7,3,1,7])=1: sequence of operations is [7,3,1,7]→[(7+3)mod10,(1+7)mod10] = [0,8] and one candy as 7+3≥10 → [(0+8)mod10] = [8], so we get only 1 candy.
// f([9])=0 as we don't perform operations with it.

// Tutorial
// It's possible to solve the problem using dynamic programming.
// For each segment [sl,sl+1,…,sr], length of which is a power of two, we will save pair — digit, which will remain and the number of candies, which we get for this segment. For segment of length 1 this pair is (sl,0).
// Note that there are at most logn different lengths of segments, and the number of segments with fixed length is at most n. It follows that there are at most nlogn segments.
// Now solution is similar to building sparse table. We will calculate needed pairs for segments in the order of increasing of their length: firstly for segments of length 2, then for length 4, etc. It turns out that the answer for segment of length 2k can be calculated from smaller segments in constant time! Indeed, to get pair (last digit, number of candies) for [sl,sl+1,…,sl+2k−1], it's sufficient to know, how many candies we got in segments [sl,sl+1,…,sl+2k−1−1], [sl+2k−1,sl+2k−1+1,…,sl+2k−1], and also what digits dig1,dig2 were left last in this segments, then last digit on our segment is equal (dig1+dig2)mod10, also if dig1+dig2≥10, we get one more candy.
// So, transition for one segment is made in O(1), which gives asymptotics O(nlogn).

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	first  uint
	second uint
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, q, l, r, i, cur, deg, left1, left2, candies1, candies2, resCandies, resLeft, len uint
	var s []uint
	var dp [][]Pair

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		s = make([]uint, n)
		dp = make([][]Pair, 20)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &s[i])
			dp[0] = append(dp[0], Pair{s[i], 0})
		}

		cur = 1
		for deg = 1; deg < 20; deg++ {
			cur *= 2

			for i = 0; i+cur <= n; i++ {
				left1 = dp[deg-1][i].first
				left2 = dp[deg-1][i+cur>>1].first
				candies1 = dp[deg-1][i].second
				candies2 = dp[deg-1][i+cur>>1].second
				resCandies = candies1 + candies2
				resLeft = (left1 + left2) % 10

				if left1+left2 >= 10 {
					resCandies++
				}

				dp[deg] = append(dp[deg], Pair{resLeft, resCandies})
			}
		}

		fmt.Fscan(reader, &q)

		for ; q > 0; q-- {
			fmt.Fscan(reader, &l, &r)

			len = r + 1 - l
			deg = 0

			for len&1 == 0 {
				deg++
				len >>= 1
			}

			fmt.Fprintln(writer, dp[deg][l-1].second)
		}
	}

	writer.Flush()
}
