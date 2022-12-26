// https://codeforces.com/problemset/problem/1374/C

// CodeForces 1374C - Move Brackets

// You are given a bracket sequence s of length n, where n is even (divisible by two). The string s consists of n2 opening brackets '(' and n2 closing brackets ')'.
// In one move, you can choose exactly one bracket and move it to the beginning of the string or to the end of the string (i.e. you choose some index i, remove the i-th character of s and insert it before or after all remaining characters of s).
// Your task is to find the minimum number of moves required to obtain regular bracket sequence from s. It can be proved that the answer always exists under the given constraints.
// Recall what the regular bracket sequence is:
//  • "()" is regular bracket sequence;
//  • if s is regular bracket sequence then "(" + s + ")" is regular bracket sequence;
//  • if s and t are regular bracket sequences then s + t is regular bracket sequence.
// For example, "()()", "(())()", "(())" and "()" are regular bracket sequences, but ")(", "()(" and ")))" are not.
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤2000) — the number of test cases. Then t test cases follow.
// The first line of the test case contains one integer n (2≤n≤50) — the length of s. It is guaranteed that n is even. The second line of the test case containg the string s consisting of n/2 opening and n/2 closing brackets.

// Output
// For each test case, print the answer — the minimum number of moves required to obtain regular bracket sequence from s. It can be proved that the answer always exists under the given constraints.

// Example
// input
// 4
// 2
// )(
// 4
// ()()
// 8
// ())()()(
// 10
// )))((((())
// output
// 1
// 0
// 1
// 3

// Note
// In the first test case of the example, it is sufficient to move the first bracket to the end of the string.
// In the third test case of the example, it is sufficient to move the last bracket to the beginning of the string.
// In the fourth test case of the example, we can choose last three openning brackets, move them to the beginning of the string and obtain "((()))(())".

// Tutorial
// Let's go from left to right over characters of s maintaining the current bracket balance (for the position i the balance is the number of opening brackets on the prefix till the i-th character minus the number of closing brackets on the same prefix).
// If the current balance becomes less than zero, then let's just take some opening bracket after the current position (it obviously exists because the number of opening equals the number of closing brackets) and move it to the beginning (so the negative balance becomes zero again and the answer increases by one). Or we can move the current closing bracket to the end of the string because it leads to the same result.
// Time complexity: O(n).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, bal int
	var t, n, ans uint
	var s string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)
			bal, ans = 0, 0

			for _, ch := range s {
				if ch == '(' {
					bal++
				} else {
					bal--

					if bal < 0 {
						ans++
						bal = 0
					}
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
