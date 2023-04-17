// https://codeforces.com/problemset/problem/1732/B

// CodeForces 1732B - Ugu

// A binary string is a string consisting only of the characters 0 and 1. You are given a binary string s1s2…sn. It is necessary to make this string non-decreasing in the least number of operations. In other words, each character should be not less than the previous. In one operation, you can do the following:
//  • Select an arbitrary index 1≤i≤n in the string;
//  • For all j≥i, change the value in the j-th position to the opposite, that is, if sj=1, then make sj=0, and vice versa.
// What is the minimum number of operations needed to make the string non-decreasing?

// Input
// Each test consists of multiple test cases. The first line contains an integer t (1≤t≤10^4) — the number of test cases. The description of test cases follows.
// The first line of each test cases a single integer n (1≤n≤10^5) — the length of the string.
// The second line of each test case contains a binary string s of length n.
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅10^5.

// Output
// For each test case, output a single integer — the minimum number of operations that are needed to make the string non-decreasing.

// Example
// input
// 8
// 1
// 1
// 2
// 10
// 3
// 101
// 4
// 1100
// 5
// 11001
// 6
// 100010
// 10
// 0000110000
// 7
// 0101010
// output
// 0
// 1
// 2
// 1
// 2
// 3
// 1
// 5

// Note
// In the first test case, the string is already non-decreasing.
// In the second test case, you can select i=1 and then s=01.
// In the third test case, you can select i=1 and get s=010, and then select i=2. As a result, we get s=001, that is, a non-decreasing string.
// In the sixth test case, you can select i=5 at the first iteration and get s=100001. Then choose i=2, then s=111110. Then we select i=1, getting the non-decreasing string s=000001.

// Tutorial
// Let's mentally imagine the following array of length n−1: ai=0 if si=si+1, and 1 otherwise. Note that if we apply the operation to the index i, then all the values ​​of the array a do not change, except for ai−1. Let's look at this in more detail:
//  • For i≤j, note that the jth and (j+1)th elements invert their value, so aj does not change.
//  • For j<i−1, note that the j-th and (j+1)-th elements do not change their value, so aj does not change.
//  • For j=i−1, note that the jth element does not change its value, but the (j+1)th element does, so aj will change its value.
// If we look at the array a for a sorted binary string, we can see that this array does not contain more than one unit (you either have a string consisting of only zeros or only ones, or it looks like this — 0000…01…11111 ).
// Let s be the number of ones in the original array a. We have now shown that the answer is ≥max(s−1,0). In fact, if the string starts with 0, then the answer is max(s−1,0), otherwise it is s.
// Let's prove that if the string starts with 0, then we can get the answer max(s−1,0) (the case with 1 will be similar). Let's show a constructive proof using a small example s=0001110010:
//  • Choose i=3, then s=0000001101,
//  • Choose i=7, then s=0000000010,
//  • Choose i=9, then s=0000000001.

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
	var t, n, ans, a, b, i uint
	var s string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)
			ans, a, b = 0, 0, 0

			for i = 0; i < n && b == 0; i++ {
				b = uint(s[i] - '0')
			}

			for ; i < n; i++ {
				a = uint(s[i] - '0')

				if a^b == 1 {
					b = a
					ans++
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
