// https://codeforces.com/problemset/problem/1315/B

// CodeForces 1315B - Homecoming

// After a long party Petya decided to return home, but he turned out to be at the opposite end of the town from his home. There are n crossroads in the line in the town, and there is either the bus or the tram station at each crossroad.
// The crossroads are represented as a string s of length n, where si=A, if there is a bus station at i-th crossroad, and si=B, if there is a tram station at i-th crossroad. Currently Petya is at the first crossroad (which corresponds to s1) and his goal is to get to the last crossroad (which corresponds to sn).
// If for two crossroads i and j for all crossroads i,i+1,…,j−1 there is a bus station, one can pay a roubles for the bus ticket, and go from i-th crossroad to the j-th crossroad by the bus (it is not necessary to have a bus station at the j-th crossroad). Formally, paying a roubles Petya can go from i to j if st=A for all i≤t<j.
// If for two crossroads i and j for all crossroads i,i+1,…,j−1 there is a tram station, one can pay b roubles for the tram ticket, and go from i-th crossroad to the j-th crossroad by the tram (it is not necessary to have a tram station at the j-th crossroad). Formally, paying b roubles Petya can go from i to j if st=B for all i≤t<j.
// For example, if s="AABBBAB", a=4 and b=3 then Petya needs:
//  • buy one bus ticket to get from 1 to 3,
//  • buy one tram ticket to get from 3 to 6,
//  • buy one bus ticket to get from 6 to 7.
// Thus, in total he needs to spend 4+3+4=11 roubles. Please note that the type of the stop at the last crossroad (i.e. the character sn) does not affect the final expense.
// Now Petya is at the first crossroad, and he wants to get to the n-th crossroad. After the party he has left with p roubles. He's decided to go to some station on foot, and then go to home using only public transport.
// Help him to choose the closest crossroad i to go on foot the first, so he has enough money to get from the i-th crossroad to the n-th, using only tram and bus tickets.

// Input
// Each test contains one or more test cases. The first line contains the number of test cases t (1≤t≤10^4).
// The first line of each test case consists of three integers a,b,p (1≤a,b,p≤10^5) — the cost of bus ticket, the cost of tram ticket and the amount of money Petya has.
// The second line of each test case consists of one string s, where si=A, if there is a bus station at i-th crossroad, and si=B, if there is a tram station at i-th crossroad (2≤|s|≤10^5).
// It is guaranteed, that the sum of the length of strings s by all test cases in one test doesn't exceed 10^5.

// Output
// For each test case print one number — the minimal index i of a crossroad Petya should go on foot. The rest of the path (i.e. from i to n he should use public transport).

// Example
// input
// 5
// 2 2 1
// BB
// 1 1 1
// AB
// 3 2 8
// AABBBBAABB
// 5 3 4
// BBBBB
// 2 1 1
// ABABAB
// output
// 2
// 1
// 3
// 1
// 6

// Tutorial
// The first thing you should do in this problem — you should understand the problem statement (which could be not very easy), and get the right answers to the sample test cases.
// Petya needs to find the minimal i such that he has enough money to get from i to n (not n+1, he doesn't need to use the transport from the last crossroad. This was a rather common mistake in misunderstanding the statement) using only public transport. We can see that if Petya can get from i to n using only public transport, he can also get from any j>i to n, using only public transport (because he will need fewer tickets).
// Let's iterate the candidates for i from n to 1 and try to find the minimal possible i. Of course, Petya can go from n to n using only public transport (he doesn't need to buy any ticket). Suppose Petya can get from j to n for some j, and it would cost him tj money. How much money he would need to get from j−1 to n? He definitely should be able to buy a ticket at station j−1. So, if it is the same ticket he should buy at station j, he will need the same amount of money, because he doesn't need to buy two consecutive equal tickets, it has no sense. Otherwise, he should buy one more ticket.
// So, a minimal amount of money tj−1 to get from j−1 to j is tj, if j<n and sj−1=sj, and tj+cost(sj−1) otherwise (cost(sj−1) is a if sj−1=A, and b otherwise). If this value is greater than p, he should go to j by foot, otherwise, we should resume the process because he can go to j−1 or even some less-numbered crossroads.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var q int
	var t, a, b, p, now, i, pre uint
	var s string
	var c []uint

	for q, _ = fmt.Fscan(reader, &t); q == 1; q, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &a, &b, &p, &s)
			s = " " + s
			c = []uint{a, b}
			now = uint(len(s)) - 2
			i = uint(len(s)) - 2

			for {
				for i > 0 && s[i] == s[now] {
					i--
				}

				pre = i + 1

				if i == 0 {
					break
				}

				if s[now] == 'A' {
					if p >= a {
						p -= a
					} else {
						break
					}
				} else if s[now] == 'B' {
					if p >= b {
						p -= b
					} else {
						break
					}
				}

				now = i
			}

			if p < c[s[now]-'A'] {
				fmt.Fprintln(writer, now+1)
			} else {
				fmt.Fprintln(writer, pre)
			}
		}
	}

	writer.Flush()
}
