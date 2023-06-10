// https://codeforces.com/problemset/problem/1767/B

// CodeForces 1767B - Block Towers

// There are n block towers, numbered from 1 to n. The i-th tower consists of ai blocks.
// In one move, you can move one block from tower i to tower j, but only if ai>aj. That move increases aj by 1 and decreases ai by 1. You can perform as many moves as you would like (possibly, zero).
// What's the largest amount of blocks you can have on the tower 1 after the moves?

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of testcases.
// The first line of each testcase contains a single integer n (2≤n≤2⋅10^5) — the number of towers.
// The second line contains n integers a1,a2,…,an (1≤ai≤10^9) — the number of blocks on each tower.
// The sum of n over all testcases doesn't exceed 2⋅10^5.

// Output
// For each testcase, print the largest amount of blocks you can have on the tower 1 after you make any number of moves (possibly, zero).

// Example
// input
// 4
// 3
// 1 2 3
// 3
// 1 2 2
// 2
// 1 1000000000
// 10
// 3 8 6 7 4 1 2 4 10 1
// output
// 3
// 2
// 500000001
// 9

// Note
// In the first testcase, you can move a block from tower 2 to tower 1, making the block counts [2,1,3]. Then move a block from tower 3 to tower 1, making the block counts [3,1,2]. Tower 1 has 3 blocks in it, and you can't obtain a larger amount.
// In the second testcase, you can move a block from any of towers 2 or 3 to tower 1, so that it has 2 blocks in it.
// In the third testcase, you can 500000000 times move a block from tower 2 to tower 1. After that the block countes will be [500000001,500000000].

// Tutorial
// Notice that it never makes sense to move blocks between the towers such that neither of them is tower 1 as that can only decrease the heights. Moreover, it never makes sense to move blocks away from the tower 1. Thus, all operations will be moving blocks from some towers to tower 1.
// At the start, which towers can move at least one block to tower 1? Well, only such i that ai>a1. What happens after you move a block? Tower 1 becomes higher, some tower becomes lower. Thus, the set of towers that can share a block can't become larger.
// Let's order the towers by the number of blocks in them. At the start, the towers that can share a block are at the end (on some suffix) in this order. After one move is made, the towers get reordered, and the suffix can only shrink.
// Ok, but if that suffix shrinks, what's the first tower that will become too low? The leftmost one that was available before. So, regardless of what the move is, the first tower that might become unavailable is the leftmost available tower. Thus, let's attempt using it until it's not too late.
// The algorithm then is the following. Find the lowest tower that can move the block to tower 1, move a block, repeat. When there are no more towers higher than tower 1, the process stops.
// However, the constraints don't allow us to do exactly that. We'll have to make at most 109 moves per testcase.
// Ok, let's move the blocks in bulk every time. Since the lowest available tower will remain the lowest until you can't use it anymore, make all the moves from it at the same time. If the current number of blocks in tower 1 is x and the current number of blocks in that tower is y, ⌈y−x2⌉ blocks can be moved.
// You can also avoid maintaining the available towers by just iterating over the towers in the increasing order of their height.
// Overall complexity: O(nlogn) per testcase.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, i, x, y uint
	var a []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
			}

			x = a[0]
			a = a[1:]

			sort.Slice(a, func(i, j int) bool {
				return a[i] < a[j]
			})

			for _, y = range a {
				if y > x {
					x += (y - x + 1) / 2
				}
			}

			fmt.Fprintln(writer, x)
		}
	}

	writer.Flush()
}
