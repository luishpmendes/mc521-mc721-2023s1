// https://codeforces.com/problemset/problem/1593/B

// CodeForces 1593B - Make it Divisible by 25

// It is given a positive integer n. In 1 move, one can select any single digit and remove it (i.e. one selects some position in the number and removes the digit located at this position). The operation cannot be performed if only one digit remains. If the resulting number contains leading zeroes, they are automatically removed.
// E.g. if one removes from the number 32925 the 3-rd digit, the resulting number will be 3225. If one removes from the number 20099050 the first digit, the resulting number will be 99050 (the 2 zeroes going next to the first digit are automatically removed).
// What is the minimum number of steps to get a number such that it is divisible by 25 and positive? It is guaranteed that, for each n occurring in the input, the answer exists. It is guaranteed that the number n has no leading zeros.

// Input
// The first line contains one integer t (1≤t≤10^4) — the number of test cases. Then t test cases follow.
// Each test case consists of one line containing one integer n (25≤n≤10^18). It is guaranteed that, for each n occurring in the input, the answer exists. It is guaranteed that the number n has no leading zeros.

// Output
// For each test case output on a separate line an integer k (k≥0) — the minimum number of steps to get a number such that it is divisible by 25 and positive.

// Example
// input
// 5
// 100
// 71345
// 3259
// 50555
// 2050047
// output
// 0
// 3
// 1
// 3
// 2

// Note
// In the first test case, it is already given a number divisible by 25.
// In the second test case, we can remove the digits 1, 3, and 4 to get the number 75.
// In the third test case, it's enough to remove the last digit to get the number 325.
// In the fourth test case, we can remove the three last digits to get the number 50.
// In the fifth test case, it's enough to remove the digits 4 and 7.

// Tutorial
// A number is divisible by 25 if and only if its last two digits represent one of the following strings: "00", "25", "50", "75".
// Let's solve for each string the following subtask: what is the minimum number of characters to be deleted so that the string becomes a suffix of the number. Then, choosing the minimum of the answers for all subtasks, we solve the whole problem.
// Let's solve the subtask for a string "X Y" where 'X' and 'Y' are digits. We can do it using the following algorithm: let's delete the last digit of the number until it is equal to 'Y', then the second to last digit of the number until it is equal to 'X'. If it is not possible, then this subtask has no solution (i.e. its result will not affect the answer).

package main

import (
	"bufio"
	"fmt"
	"os"
)

var subseqs = []string{"00", "25", "50", "75"}

func solve(s string, t string) uint {
	var sptr int
	var ans uint

	sptr = len(s) - 1
	ans = 0

	for sptr >= 0 && s[sptr] != t[1] {
		sptr--
		ans++
	}

	if sptr < 0 {
		ans = 100
	} else {
		sptr--

		for sptr >= 0 && s[sptr] != t[0] {
			sptr--
			ans++
		}

		if sptr < 0 {
			ans = 100
		}
	}

	return ans
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, ans, aux uint
	var n string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			ans = 100

			for _, e := range subseqs {
				aux = solve(n, e)

				if ans > aux {
					ans = aux
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
