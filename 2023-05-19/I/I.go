// https://codeforces.com/problemset/problem/1793/A

// CodeForces 1793A - Yet Another Promotion

// The famous store "Second Food" sells groceries only two days a month. And the prices in each of days differ. You wanted to buy n kilos of potatoes for a month. You know that on the first day of the month 1 kilo of potatoes costs a coins, and on the second day b coins. In "Second Food" you can buy any integer kilograms of potatoes.
// Fortunately, "Second Food" has announced a promotion for potatoes, which is valid only on the first day of the month — for each m kilos of potatoes you buy, you get 1 kilo as a gift! In other words, you can get m+1 kilograms by paying for m kilograms.
// Find the minimum number of coins that you have to spend to buy at least n kilos of potatoes.

// Input
// Each test contains multiple test cases. The first line contains the number of test cases t (1≤t≤10000). Description of the test cases follows.
// The first line of each test case contains two integers a and b (1≤a,b≤10^9) — the prices of 1 kilo of potatoes on the first and second days, respectively.
// The second line contains two integers n and m (1≤n,m≤10^9) — the required amount of potatoes to buy and the amount of potatoes to use the promotion.

// Output
// For each test case print one integer — the minimum number of coins that you have to pay to buy at least n kilos of potatoes.

// Example
// input
// 5
// 5 4
// 3 1
// 5 4
// 3 2
// 3 4
// 3 5
// 20 15
// 10 2
// 1000000000 900000000
// 1000000000 8
// output
// 9
// 10
// 9
// 135
// 888888888900000000

// Note
// In the first test case, on the first day you buy 1 kilo and get 1 more for a promotion. On the second day, you can buy 1 kilo of potatoes. Thus, you will spend 5+4=9 coins in total.
// In the second test case, on the first day you buy 2 kilo and get another 1 more for a promotion. This way you will spend 2⋅5=10 coins.

// Tutorial
// Let n=(m+1)⋅q+r.
// Note that you need to use a promotion if a⋅m≤b⋅(m+1). In this case, we will buy potatoes q times for the promotion. The remaining potatoes (or all if the promotion is unprofitable) can be bought at min(a,b) per kilogram.
// Then the answer is:
// q⋅min(a⋅m,b⋅(m+1))+r⋅min(a,b)
// Thus this solution works in O(1)

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
	var t, a, b, n, m, q, r, ans uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &a, &b, &n, &m)
			q = n / (m + 1)
			r = n - q*(m+1)

			if a*m <= b*(m+1) {
				ans = q * a * m
			} else {
				ans = q * b * (m + 1)
			}

			if a < b {
				ans += r * a
			} else {
				ans += r * b
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
