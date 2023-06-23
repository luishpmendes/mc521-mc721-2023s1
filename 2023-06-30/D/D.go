// https://codeforces.com/problemset/problem/1811/A

// CodeForces 1811A - Insert Digit
// You have a positive number of length n and one additional digit.
// You can insert this digit anywhere in the number, including at the beginning or at the end.
// Your task is to make the result as large as possible.
// For example, you have the number 76543, and the additional digit is 4. Then the maximum number you can get is 765443, and it can be obtained in two ways — by inserting a digit after the 3th or after the 4th digit of the number.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of test cases.
// The descriptions of the test cases follow.
// The first line of the description of each test case contains two integers n and d (1≤n≤2⋅10^5; 0≤d≤9) — the length of the number and an additional digit, respectively.
// The second line of the description of each test case contains a string consisting of n digits — the number that you have initially. It is guaranteed that the number does not contain leading zeros.
// It is guaranteed that the sum of n for all test cases does not exceed 2⋅10^5.

// Output
// For each test case, output a string consisting of n+1 digits — the maximum possible number that can be obtained.

// Example
// input
// 11
// 5 4
// 76543
// 1 0
// 1
// 2 5
// 44
// 3 6
// 666
// 5 6
// 13579
// 5 8
// 97531
// 19 4
// 9876543210123456789
// 5 7
// 73737
// 8 1
// 20000000
// 7 0
// 7058959
// 12 1
// 828127127732
// output
// 765443
// 10
// 544
// 6666
// 613579
// 987531
// 98765443210123456789
// 773737
// 210000000
// 70589590
// 8281271277321

// Tutorial
// Note that numbers of the same length are compared lexicographically. That is, until some index the numbers will match, and then the digit in our number should be greater.
// Let's write out the numbers s1,s2,…si until si≥d. As soon as this condition is false or the line ends, insert the digit d.
// We got the lexicographically maximum number, which means just the maximum number.

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
	var t, n, d, i uint
	var s string
	var ans []byte

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &d, &s)
			ans = make([]byte, 0, n+1)

			for i = 0; i < n; i++ {
				if uint(s[i]-'0') >= d {
					ans = append(ans, s[i])
				} else {
					ans = append(ans, byte(d+48))
					ans = append(ans, s[i:]...)
					break
				}
			}

			if i == n {
				ans = append(ans, byte(d+48))
			}

			fmt.Fprintln(writer, string(ans))
		}
	}

	writer.Flush()
}
