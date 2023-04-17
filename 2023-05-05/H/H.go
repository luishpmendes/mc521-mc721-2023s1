// https://codeforces.com/problemset/problem/1774/C

// CodeForces 1774C - Ice and Fire

// Little09 and his friends are playing a game. There are n players, and the temperature value of the player i is i.
// The types of environment are expressed as 0 or 1. When two players fight in a specific environment, if its type is 0, the player with a lower temperature value in this environment always wins; if it is 1, the player with a higher temperature value in this environment always wins. The types of the n−1 environments form a binary string s with a length of n−1.
// If there are x players participating in the game, there will be a total of x−1 battles, and the types of the x−1 environments will be the first x−1 characters of s. While there is more than one player left in the tournament, choose any two remaining players to fight. The player who loses will be eliminated from the tournament. The type of the environment of battle i is si.
// For each x from 2 to n, answer the following question: if all players whose temperature value does not exceed x participate in the game, how many players have a chance to win?

// Input
// Each test contains multiple test cases. The first line contains a single integer t (1≤t≤10^3)  — the number of test cases. The description of the test cases follows.
// The first line of each test case contains a single integer n (2≤n≤2⋅10^5)  — the number of players.
// The second line of each test case contains a binary string s with a length n−1.
// It is guaranteed that the sum of n over all test cases does not exceed 3⋅10^5.

// Output
// For each test case output n−1 integers  — for each x from 2 to n, output the number of players that have a chance to win.

// Example
// input
// 2
// 4
// 001
// 4
// 101
// output
// 1 1 3
// 1 2 3

// Note
// In the first test case, for x=2 and x=3, only the player whose temperature value is 1 can be the winner. For x=4, the player whose temperature value is 2,3,4 can be the winner.

// Tutorial
// We define fi to mean that the maximum x satisfies si−x+1=si−x+2=...=si.
// It can be proved that for x players, fx−1 players are bound to lose and the rest have a chance to win. So the answer to the first i battles is ansi=i−fi+1.
// Next, we prove this conclusion. Suppose there are n players and n−1 battles, and sn−1=1, and there are x consecutive 1 at the end. If x=n−1, then obviously only the n-th player can win. Otherwise, sn−1−x must be 0. Consider the following facts:
//  1. Players 1 to x have no chance to win. If the player i (1≤i≤x) can win, he must defeat the player whose temperature value is lower than him in the last x battles. However, in total, only the i−1 player's temperature value is lower than his. Because i−1<x, the i-th player cannot win.
//  2. Players from x+1 to n have a chance to win. For the player i (x+1≤i≤n), we can construct: in the first n−2−x battles, we let all players whose temperature value in [x+1,n] except the player i fight so that only one player will remain. In the (n−1−x)-th battle, we let the remaining player fight with the player 1. Since sn−1−x=0, the player 1 will win. Then there are only the first x players and the player i in the remaining x battles, so the player i can win.
// For sn−1=0, the situation is similar and it will not be repeated here.

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
	var t, n, cnt, i uint
	var s string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)

			cnt = 0

			for i = 0; i+1 < n; i++ {
				if i == 0 {
					cnt = 1
				} else if s[i-1] == s[i] {
					cnt++
				} else {
					cnt = 1
				}

				fmt.Fprint(writer, i+2-cnt, " ")
			}

			fmt.Fprintln(writer)
		}
	}

	writer.Flush()
}
