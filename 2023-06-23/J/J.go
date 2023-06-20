// https://codeforces.com/problemset/problem/1511/C

// CodeForces 1511C - Yet Another Card Deck
// You have a card deck of n cards, numbered from top to bottom, i. e. the top card has index 1 and bottom card — index n. Each card has its color: the i-th card has color ai.
// You should process q queries. The j-th query is described by integer tj. For each query you should:
//  • find the highest card in the deck with color tj, i. e. the card with minimum index;
//  • print the position of the card you found;
//  • take the card and place it on top of the deck.

// Input
// The first line contains two integers n and q (2≤n≤3⋅10^5; 1≤q≤3⋅10^5) — the number of cards in the deck and the number of queries.
// The second line contains n integers a1,a2,…,an (1≤ai≤50) — the colors of cards.
// The third line contains q integers t1,t2,…,tq (1≤tj≤50) — the query colors. It's guaranteed that queries ask only colors that are present in the deck.

// Output
// Print q integers — the answers for each query.

// Example
// input
// 7 5
// 2 1 1 4 3 3 1
// 3 2 1 1 4
// output
// 5 2 3 1 5

// Note
// Description of the sample:
//  1. the deck is [2,1,1,4,3–,3,1] and the first card with color t1=3 has position 5;
//  2. the deck is [3,2–,1,1,4,3,1] and the first card with color t2=2 has position 2;
//  3. the deck is [2,3,1–,1,4,3,1] and the first card with color t3=1 has position 3;
//  4. the deck is [1–,2,3,1,4,3,1] and the first card with color t4=1 has position 1;
//  5. the deck is [1,2,3,1,4–,3,1] and the first card with color t5=4 has position 5.

// Tutorial
// Let's look at one fixed color. When we search a card of such color, we take the card with minimum index and after we place it on the top of the deck it remains the one with minimum index.
// It means that for each color we take and move the same card — one card for each color. In other words, we need to keep track of only k cards, where k is the number of colors (k≤50). As a result, if posc is the position of a card of color c then we can simulate a query in the following way: for each color c such that posc<postj we increase posc by one (since the card will move down) and then set postj=1. Complexity is O(n+qk).
// But, if we look closely, we may note that we don't even need array posc. We can almost manually find the first card of color tj and move it to the first position either by series of swaps or, for example, using rotate function (present in C++) and it will work fast.
// Why? Let's look at one color c. For the first time it will cost O(n) operations to search the corresponding card and move it to the position 1. But after that, at any moment of time, the position of the card won't exceed k, since all cards before are pairwise different (due to the nature of queries). So, all next moves the color c costs only O(k) time.
// As a result, the complexity of such almost naive solution is O(kn+qk).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func find(a []uint, x uint) int {
	for i, v := range a {
		if v == x {
			return i
		}
	}
	return -1
}

func rotate(a []uint, p int) {
	x := a[p]
	copy(a[1:], a[:p])
	a[0] = x
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, p int
	var n, q, i, x uint
	var a []uint

	for t, _ = fmt.Fscan(reader, &n, &q); t == 2; t, _ = fmt.Fscan(reader, &n, &q) {
		a = make([]uint, n)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}

		for ; q > 0; q-- {
			fmt.Fscan(reader, &x)
			p = find(a, x)
			fmt.Fprint(writer, p+1, " ")
			rotate(a, p)
		}
	}

	writer.Flush()
}
