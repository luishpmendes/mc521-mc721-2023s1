// https://codeforces.com/problemset/problem/1566/B

// CodeForces 1566B - MIN-MEX Cut

// A binary string is a string that consists of characters 0 and 1.
// Let MEX of a binary string be the smallest digit among 0, 1, or 2 that does not occur in the string. For example, MEX of 001011 is 2, because 0 and 1 occur in the string at least once, MEX of 1111 is 0, because 0 and 2 do not occur in the string and 0<2.
// A binary string s is given. You should cut it into any number of substrings such that each character is in exactly one substring. It is possible to cut the string into a single substring — the whole string.
// A string a is a substring of a string b if a can be obtained from b by deletion of several (possibly, zero or all) characters from the beginning and several (possibly, zero or all) characters from the end.
// What is the minimal sum of MEX of all substrings pieces can be?

// Input
// The input consists of multiple test cases. The first line contains a single integer t (1≤t≤10^4) — the number of test cases. Description of the test cases follows.
// Each test case contains a single binary string s (1≤|s|≤10^5).
// It's guaranteed that the sum of lengths of s over all test cases does not exceed 105.

// Output
// For each test case print a single integer — the minimal sum of MEX of all substrings that it is possible to get by cutting s optimally.

// Example
// input
// 6
// 01
// 1111
// 01100
// 101
// 0000
// 01010
// output
// 1
// 0
// 2
// 1
// 1
// 2

// Note
// In the first test case the minimal sum is MEX(0)+MEX(1)=1+0=1.
// In the second test case the minimal sum is MEX(1111)=0.
// In the third test case the minimal sum is MEX(01100)=2.

// Tutorial
// Hint 1
// The answer is never greater than 2.
// Hint 2
// If a string consists only of 1, then the answer is 0.
// Hint 3
// If all zeroes are consequent, then the answer is 1.
// Editorial
// The answer is never greater than 2, because MEX of the whole string is not greater than 2.
// The answer is 0 only if there are no zeroes in the string.
// Now we need to understand, when the answer is 1 or when it is 2. The sum of MEX is 1 only if all zeroes create a single segment without ones. Then we can cut this segment out, its MEX is 1, everything else is ones, their total MEX is 0.
// If the zeroes do not create a single segment without ones, the there are two such zeroes that there is a 1 between them. Then either these zeroes are in a single segment with the 1, so total MEX is not less than 2 or these zeroes are in different segments and the answer is still not less then 2.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, ans, zeroes, first, last uint
	var s string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &s)
			zeroes = uint(strings.Count(s, "0"))

			if zeroes == 0 {
				ans = 0
			} else {
				first = uint(strings.Index(s, "0"))
				last = uint(strings.LastIndex(s, "0"))

				if zeroes+first == last+1 {
					ans = 1
				} else {
					ans = 2
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
