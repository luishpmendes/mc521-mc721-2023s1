// https://codeforces.com/problemset/problem/1792/A

// CodeForces 1792A - GamingForces

// Monocarp is playing a computer game. He's going to kill n monsters, the i-th of them has hi health.
// Monocarp's character has two spells, either of which he can cast an arbitrary number of times (possibly, zero) and in an arbitrary order:
//  • choose exactly two alive monsters and decrease their health by 1;
//  • choose a single monster and kill it.
// When a monster's health becomes 0, it dies.
// What's the minimum number of spell casts Monocarp should perform in order to kill all monsters?

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of testcases.
// The first line of each testcase contains a single integer n (1≤n≤100) — the number of monsters.
// The second line contains n integers h1,h2,…,hn (1≤hi≤100) — the health of each monster.
// The sum of n over all testcases doesn't exceed 2⋅10^4.

// Output
// For each testcase, print a single integer — the minimum number of spell casts Monocarp should perform in order to kill all monsters.

// Example
// input
// 3
// 4
// 1 2 1 2
// 3
// 2 4 2
// 5
// 1 2 3 4 5
// output
// 3
// 3
// 5

// Note
// In the first testcase, the initial health list is [1,2,1,2]. Three spells are casted:
//  • the first spell on monsters 1 and 2 — monster 1 dies, monster 2 has now health 1, new health list is [0,1,1,2];
//  • the first spell on monsters 3 and 4 — monster 3 dies, monster 4 has now health 1, new health list is [0,1,0,1];
//  • the first spell on monsters 2 and 4 — both monsters 2 and 4 die.
// In the second testcase, the initial health list is [2,4,2]. Three spells are casted:
//  • the first spell on monsters 1 and 3 — both monsters have health 1 now, new health list is [1,4,1];
//  • the second spell on monster 2 — monster 2 dies, new health list is [1,0,1];
//  • the first spell on monsters 1 and 3 — both monsters 1 and 3 die.
// In the third testcase, the initial health list is [1,2,3,4,5]. Five spells are casted. The i-th of them kills the i-th monster with the second spell. Health list sequence: [1,2,3,4,5] → [0,2,3,4,5] → [0,0,3,4,5] → [0,0,0,4,5] → [0,0,0,0,5] → [0,0,0,0,0].

// Tutorial
// The first spell looks pretty weak compared to the second spell. Feels like you almost always replace one with another. Let's show that you can totally avoid casting the spell of the first type twice or more on one monster.
// Let the two first spell casts be (i,j) and (i,k) for some monsters i,j and k. You can replace them by a cast of the second spell on i and a cast of the first spell on (j,k). That would deal even more damage to i and the same amount to j and k. The number of casts doesn't change.
// Thus, it only makes sense to use the first spell on monsters with 1 health. Calculate the number of them, kill the full pairs of them with the first spell, and use the second spell on the remaining monsters.
// Overall complexity: O(n) per testcase.

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
	var t, n, h, i, cnt1 uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			cnt1 = 0

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &h)

				if h == 1 {
					cnt1++
				}
			}

			fmt.Fprintln(writer, n-cnt1>>1)
		}
	}

	writer.Flush()
}
