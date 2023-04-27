// https://codeforces.com/problemset/problem/1245/C

// CodeForces 1245C - Constanze's Machine

// Constanze is the smartest girl in her village but she has bad eyesight.
// One day, she was able to invent an incredible machine! When you pronounce letters, the machine will inscribe them onto a piece of paper. For example, if you pronounce 'c', 'o', 'd', and 'e' in that order, then the machine will inscribe "code" onto the paper. Thanks to this machine, she can finally write messages without using her glasses.
// However, her dumb friend Akko decided to play a prank on her. Akko tinkered with the machine so that if you pronounce 'w', it will inscribe "uu" instead of "w", and if you pronounce 'm', it will inscribe "nn" instead of "m"! Since Constanze had bad eyesight, she was not able to realize what Akko did.
// The rest of the letters behave the same as before: if you pronounce any letter besides 'w' and 'm', the machine will just inscribe it onto a piece of paper.
// The next day, I received a letter in my mailbox. I can't understand it so I think it's either just some gibberish from Akko, or Constanze made it using her machine. But since I know what Akko did, I can just list down all possible strings that Constanze's machine would have turned into the message I got and see if anything makes sense.
// But I need to know how much paper I will need, and that's why I'm asking you for help. Tell me the number of strings that Constanze's machine would've turned into the message I got.
// But since this number can be quite large, tell me instead its remainder when divided by 109+7.
// If there are no strings that Constanze's machine would've turned into the message I got, then print 0.

// Input
// Input consists of a single line containing a string s (1≤|s|≤10^5) — the received message. s contains only lowercase Latin letters.

// Output
// Print a single integer — the number of strings that Constanze's machine would've turned into the message s, modulo 10^9+7.

// Examples

// input
// ouuokarinn
// output
// 4

// input
// banana
// output
// 1

// input
// nnn
// output
// 3

// input
// amanda
// output
// 0

// Note
// For the first example, the candidate strings are the following: "ouuokarinn", "ouuokarim", "owokarim", and "owokarinn".
// For the second example, there is only one: "banana".
// For the third example, the candidate strings are the following: "nm", "mn" and "nnn".
// For the last example, there are no candidate strings that the machine can turn into "amanda", since the machine won't inscribe 'm'.

// Tutorial
// If s has any 'm' or 'w', the answer is 0. Otherwise, we can do dp.
// Define dpi to be the number of strings that Constanze's machine would've turned into the first i characters of s. Then, dp0=dp1=1. For i>1, transition is dpi=dpi−1+dpi−2 if both si=si−1 and si is either 'u' or 'n', and simply dpi=dpi−1 otherwise. Answer is dp|s|.
// Alternatively, notice that the maximum cardinality segment consisiting of k letters 'u' or 'n' multiplies the answer by the k-th term of Fibonacci sequence. Thus, you can precalculate it and some sort of two iterators.

package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var s string
	var n, i uint
	var dp []uint

	for t, _ = fmt.Fscan(reader, &s); t == 1; t, _ = fmt.Fscan(reader, &s) {
		for _, x := range s {
			if x == 'w' || x == 'm' {
				fmt.Println(0)
				return
			}
		}

		n = uint(len(s))
		dp = make([]uint, n+1)
		dp[0] = 1
		dp[1] = 1

		for i = 2; i <= n; i++ {
			dp[i] = dp[i-1]

			if s[i-1] == s[i-2] && (s[i-1] == 'u' || s[i-1] == 'n') {
				dp[i] = (dp[i] + dp[i-2]) % MOD
			}
		}

		fmt.Println(dp[n])
	}

	writer.Flush()
}
