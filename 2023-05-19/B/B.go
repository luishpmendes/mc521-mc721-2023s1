// https://codeforces.com/problemset/problem/1792/B

// CodeForces 1792B - Stand-up Comedian

// Eve is a beginner stand-up comedian. Her first show gathered a grand total of two spectators: Alice and Bob.
// Eve prepared a1+a2+a3+a4 jokes to tell, grouped by their type:
//  • type 1: both Alice and Bob like them;
//  • type 2: Alice likes them, but Bob doesn't;
//  • type 3: Bob likes them, but Alice doesn't;
//  • type 4: neither Alice nor Bob likes them.
// Initially, both spectators have their mood equal to 0. When a spectator hears a joke he/she likes, his/her mood increases by 1. When a spectator hears a joke he/she doesn't like, his/her mood decreases by 1. If the mood of a spectator becomes negative (strictly below zero), he/she leaves.
// When someone leaves, Eve gets sad and ends the show. If no one leaves, and Eve is out of jokes, she also ends the show.
// Thus, Eve wants to arrange her jokes in such a way that the show lasts as long as possible. Help her to calculate the maximum number of jokes she can tell before the show ends.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of testcases.
// The only line of each testcase contains four integers a1,a2,a3,a4 (0≤a1,a2,a3,a4≤10^8; a1+a2+a3+a4≥1) — the number of jokes of each type Eve prepared.

// Output
// For each testcase, print a single integer — the maximum number of jokes Eve can tell before at least one of the spectators leaves or before she runs out of jokes.

// Example
// input
// 4
// 5 0 0 0
// 0 0 0 5
// 2 5 10 6
// 3 0 0 7
// output
// 5
// 1
// 15
// 7

// Note
// In the first testcase, Eve only has jokes of the first type. Thus, there's no order to choose. She tells all her jokes, both Alice and Bob like them. Their mood becomes 5. The show ends after Eve runs out of jokes.
// In the second testcase, Eve only has jokes of the fourth type. Thus, once again no order to choose. She tells a joke, and neither Alice, nor Bob likes it. Their mood decrease by one, becoming −1. They both have negative mood, thus, both leave, and the show ends.
// In the third testcase, first, Eve tells both jokes of the first type. Both Alice and Bob has mood 2. Then she can tell 2 jokes of the third type. Alice's mood becomes 0. Bob's mood becomes 4. Then 4 jokes of the second type. Alice's mood becomes 4. Bob's mood becomes 0. Then another 4 jokes of the third type. Alice's mood becomes 0. Bob's mood becomes 4. Then the remaining joke of the second type. Alice's mood becomes 1. Bob's mood becomes 3. Then one more joke of the third type, and a joke of the fourth type, for example. Alice's mood becomes −1, she leaves, and the show ends.
// In the fourth testcase, Eve should first tell the jokes both spectators like, then the jokes they don't. She can tell 4 jokes of the fourth type until the spectators leave.

// Tutorial
// First, let Eve tell the jokes of the first type — they will never do any harm. At the same time, let her tell the jokes of the fourth time at the very end — they will not do any good.
// Types two and three are kind of opposites of each other. If you tell jokes of each of them one after another, then the moods of both spectators don't change. Let's use that to our advantage. Tell the jokes of these types in pairs until one of them runs out. There's a little corner case here, though. If there were no jokes of the first type, then you can't use a single pair because of the spectators leaves after one joke.
// Finally, try to tell the remaining jokes of the same type before the fourth type. So the construction looks like 1,1,…,1,2,3,2,3,…,2,3,2,2,2,…,2,4,4,4,…,4 with 2 and 3 possibly swapped with each other.
// Let's recover the answer from that construction. After the first type, both moods are a1. After the alternating jokes, the moods are still the same. After that, one of the spectators will have his/her mood only decreasing until the end. Once it reaches −1, the show ends.
// Thus, Eve can tell a1+min(a2,a3)⋅2+min(a1+1,abs(a2−a3)+a4) jokes if a1≠0. Otherwise, it's always 1.
// Overall complexity: O(1).

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
	var t, a1, a2, a3, a4, ans uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &a1, &a2, &a3, &a4)

			if a1 == 0 {
				ans = 1
			} else {
				ans = a1

				if a2 < a3 {
					ans += 2 * a2
				} else {
					ans += 2 * a3
				}

				if a2 > a3 {
					if a1+1 < a2-a3+a4 {
						ans += a1 + 1
					} else {
						ans += a2 - a3 + a4
					}
				} else {
					if a1+1 < a3-a2+a4 {
						ans += a1 + 1
					} else {
						ans += a3 - a2 + a4
					}
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
