// https://codeforces.com/problemset/problem/1345/B

// CodeForces 1345B - Card Constructions

// A card pyramid of height 1 is constructed by resting two cards against each other. For h>1, a card pyramid of height h is constructed by placing a card pyramid of height h−1 onto a base. A base consists of h pyramids of height 1, and h−1 cards on top. For example, card pyramids of heights 1, 2, and 3 look as follows:
// You start with n cards and build the tallest pyramid that you can. If there are some cards remaining, you build the tallest pyramid possible with the remaining cards. You repeat this process until it is impossible to build another pyramid. In the end, how many pyramids will you have constructed?

// Input
// Each test consists of multiple test cases. The first line contains a single integer t (1≤t≤1000) — the number of test cases. Next t lines contain descriptions of test cases.
// Each test case contains a single integer n (1≤n≤10^9) — the number of cards.
// It is guaranteed that the sum of n over all test cases does not exceed 10^9.

// Output
// For each test case output a single integer — the number of pyramids you will have constructed in the end.

// Example
// input
// 5
// 3
// 14
// 15
// 24
// 1
// output
// 1
// 2
// 1
// 3
// 0

// Note
// In the first test, you construct a pyramid of height 1 with 2 cards. There is 1 card remaining, which is not enough to build a pyramid.
// In the second test, you build two pyramids, each of height 2, with no cards remaining.
// In the third test, you build one pyramid of height 3, with no cards remaining.
// In the fourth test, you build one pyramid of height 3 with 9 cards remaining. Then you build a pyramid of height 2 with 2 cards remaining. Then you build a final pyramid of height 1 with no cards remaining.
// In the fifth test, one card is not enough to build any pyramids.

// Tutorial
// Let's count the number of cards in a pyramid of height h. There are 2(1+2+3+⋯+h) cards standing up, and there are 0+1+2+⋯+(h−1) horizontal cards. So, there are 2*(h(h+1))/2+((h−1)h)/2=3/2*h^2+1/2*h cards total. Using this formula, we can quickly find the largest height h that uses at most n cards.
// The quadratic formula or binary search can be used here, but are unnecessary. Simply iterating through all h values works in O(√n) time per test.
// It's enough to see that this takes O(t√N) time overall, where N is the sum of n across all test cases. But interestingly, we can argue for a tighter bound of O(√(tN)) due to the Cauchy-Schwarz Inequality:
// ∑_(i=1)^(t)(1⋅√(ni))≤(∑_(i=1)^(t)1^2)(∑_(i=1)^(t)(√(n_i))^2)=√(tN)

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
	var t, n, h, ans uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			for ans = 0; n > 1; ans++ {
				h = 1

				for h*(3*h+1) <= 2*n {
					h++
				}

				n -= ((h - 1) * (3*h - 2)) / 2
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
