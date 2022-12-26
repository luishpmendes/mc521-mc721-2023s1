// https://codeforces.com/problemset/problem/1765/E

// CodeForces 1765E - Exchange

// Monocarp is playing a MMORPG. There are two commonly used types of currency in this MMORPG — gold coins and silver coins. Monocarp wants to buy a new weapon for his character, and that weapon costs n silver coins. Unfortunately, right now, Monocarp has no coins at all.
// Monocarp can earn gold coins by completing quests in the game. Each quest yields exactly one gold coin. Monocarp can also exchange coins via the in-game trading system. Monocarp has spent days analyzing the in-game economy; he came to the following conclusion: it is possible to sell one gold coin for a silver coins (i. e. Monocarp can lose one gold coin to gain a silver coins), or buy one gold coin for b silver coins (i. e. Monocarp can lose b silver coins to gain one gold coin).
// Now Monocarp wants to calculate the minimum number of quests that he has to complete in order to have at least n silver coins after some abuse of the in-game economy. Note that Monocarp can perform exchanges of both types (selling and buying gold coins for silver coins) any number of times.

// Input
// The first line contains one integer t (1≤t≤10^4) — the number of test cases.
// Each test case consists of one line containing three integers n, a and b (1≤n≤10^7; 1≤a,b≤50).

// Output
// For each test case, print one integer — the minimum possible number of quests Monocarp has to complete.

// Example
// input
// 4
// 100 25 30
// 9999997 25 50
// 52 50 48
// 49 50 1
// output
// 4
// 400000
// 1
// 1

// Note
// In the first test case of the example, Monocarp should complete 4 quests, and then sell 4 gold coins for 100 silver coins.
// In the second test case, Monocarp should complete 400000 quests, and then sell 400000 gold coins for 10 million silver coins.
// In the third test case, Monocarp should complete 1 quest, sell the gold coin for 50 silver coins, buy a gold coin for 48 silver coins, and then sell it again for 50 coins. So, he will have 52 silver coins.
// In the fourth test case, Monocarp should complete 1 quest and then sell the gold coin he has obtained for 50 silver coins.

// Tutorial
// If a>b, then Monocarp can go infinite by obtaining just one gold coin: exchanging it for silver coins and then buying it back will earn him a−b silver coins out of nowhere. So, the answer is 1 no matter what n is.
// If a≤b, then it's suboptimal to exchange gold coins for silver coins and then buy the gold coins back. Monocarp should earn the minimum possible number of gold coins so that they all can be exchanged into at least n silver coins; so, the number of gold coins he needs is ⌈n/a⌉.
// One small note: you shouldn't use the functions like ceil to compute a fraction rounded up, since you may get some computation errors related to using floating-point numbers (and possibly getting precision loss). Instead, you should calculate ⌈n/a⌉ as ⌊(n+a−1)/a⌊ (this works for non-negative n and positive a), and every programming language provides an integer division operator which discards the fractional part and thus doesn't use floating-point computations at all.

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
	var t, n, a, b uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &a, &b)

			if a > b {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, (n-1)/a+1)
			}
		}
	}

	writer.Flush()
}
