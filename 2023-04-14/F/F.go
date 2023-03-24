// https://codeforces.com/problemset/problem/1343/B

// CodeForces 1343B - Balanced Array

// You are given a positive integer n, it is guaranteed that n is even (i.e. divisible by 2).
// You want to construct the array a of length n such that:
//  • The first n/2 elements of a are even (divisible by 2);
//  • the second n/2 elements of a are odd (not divisible by 2);
//  • all elements of a are distinct and positive;
//  • the sum of the first half equals to the sum of the second half (∑_(i=1)^(n/2)(ai)=∑_(i=n/2+1)^(n)(ai)).
// If there are multiple answers, you can print any. It is not guaranteed that the answer exists.
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤10^4) — the number of test cases. Then t test cases follow.
// The only line of the test case contains one integer n (2≤n≤2⋅10^5) — the length of the array. It is guaranteed that that n is even (i.e. divisible by 2).
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅10^5 (∑n≤2⋅10^5).

// Output
// For each test case, print the answer — "NO" (without quotes), if there is no suitable answer for the given test case or "YES" in the first line and any suitable array a1,a2,…,an (1≤ai≤10^9) satisfying conditions from the problem statement on the second line.

// Example
// input
// 5
// 2
// 4
// 6
// 8
// 10
// output
// NO
// YES
// 2 4 1 5
// NO
// YES
// 2 4 6 8 1 3 5 11
// NO

// Tutorial
// Firstly, if n is not divisible by 4 then the answer is "NO" because the parities of halves won't match. Otherwise, the answer is always "YES". Let's construct it as follows: firstly, let's create the array [2,4,6,…,n,1,3,5,…,n−1]. This array is almost good except one thing: the sum in the right half is exactly n/2 less than the sum in the left half. So we can fix it easily: just add n/2 to the last element.

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

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			n /= 2

			if n&1 == 0 {
				fmt.Fprintln(writer, "YES")

				for i = 1; i <= n; i++ {
					fmt.Fprint(writer, 2*i, " ")
				}

				for i = 1; i < n; i++ {
					fmt.Fprint(writer, 2*i-1, " ")
				}

				fmt.Fprintln(writer, 3*n-1)
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}

	writer.Flush()
}
