// https://codeforces.com/problemset/problem/1776/H

// CodeForces 1776H - Beppa and SwerChat

// Beppa and her circle of geek friends keep up to date on a group chat in the instant messaging app SwerChatTM.
// The group has n members, excluding Beppa. Each of those members has a unique ID between 1 and n. When a user opens a group chat, SwerChatTM displays the list of other members of that group, sorted by decreasing times of last seen online (so the member who opened the chat most recently is the first of the list). However, the times of last seen are not displayed.
// Today, Beppa has been busy all day: she has only opened the group chat twice, once at 9:00 and once at 22:00. Both times, she wrote down the list of members in the order they appeared at that time. Now she wonders: what is the minimum number of other members that must have been online at least once between 9:00 and 22:00?
// Beppa is sure that no two members are ever online at the same time and no members are online when Beppa opens the group chat at 9:00 and 22:00.

// Input
// Each test contains multiple test cases. The first line contains an integer t (1≤t≤10000) — the number of test cases. The descriptions of the t test cases follow.
// The first line of each test case contains an integer n (1≤n≤10^5) — the number of members of the group excluding Beppa.
// The second line contains n integers a1,a2,…,an (1≤ai≤n) — the list of IDs of the members, sorted by decreasing times of last seen online at 9:00.
// The third line contains n integers b1,b2,…,bn (1≤bi≤n) — the list of IDs of the members, sorted by decreasing times of last seen online at 22:00.
// For all 1≤i<j≤n, it is guaranteed that ai≠aj and bi≠bj.
// It is also guaranteed that the sum of the values of n over all test cases does not exceed 10^5.

// Output
// For each test case, print the minimum number of members that must have been online between 9:00 and 22:00.

// Example
// input
// 4
// 5
// 1 4 2 5 3
// 4 5 1 2 3
// 6
// 1 2 3 4 5 6
// 1 2 3 4 5 6
// 8
// 8 2 4 7 1 6 5 3
// 5 6 1 4 8 2 7 3
// 1
// 1
// 1
// output
// 2
// 0
// 4
// 0

// Note
// In the first test case, members 4,5 must have been online between 9:00 and 22:00.
// In the second test case, it is possible that nobody has been online between 9:00 and 22:00.

// Tutorial
// Instead of finding the minimum number of members that must have been online at least once
// between 9:00 and 22:00, let us study the complement of that set: What is the maximum number of
// members that have never been online? We have the following observations:
// • Members that have never been online must be a suffix of the sequence b. It cannot happen
// that member bi has been online but bi−1 has not, for 2 ≤ i ≤ n.
// • Consider two members x and y who have never been online. Their relative order in a and b is
// the same. Indeed, if someone else goes online the relative order of x and y in the list does not
// change.
// From the above observations, we deduce that the members who have not been online form a suffix
// of b that is also a subsequence of a. Let p : {1, 2, . . . , n} → {1, 2, . . . , n} be the function such
// that ap(x) = x; we say that a sequence s of length m is a subsequence of a if p(si−1) < p(si), for
// 2 ≤ i ≤ m.
// It turns out that any “suffix of b that is a subsequence of a” can be a valid subset of members who
// have never been online as the following construction shows. For any such suffix bt
// , bt+1, . . . , bn, it
// could be that members bt−1, bt−2, . . . , b1 went online in this order between 9:00 and 22:00, leading
// to a valid list b of last seen online at 22:00.
// Thus, the answer to the problem is the length of the longest suffix of b that is also a subsequence
// of a. Let us explain how to find such length quickly.
// We first preprocess the sequence a to find the corresponding function p. Then, we iterate from bn
// to b1. If, for some i we discover that p(bi) < p(bi−1), then we know that bi−1 must have been online,
// and we denote this position i as r. In this way, br, br+1, . . . , bn is the sought longest suffix of b. The
// minimum number of members that must have been online at least once between 9:00 and 22:00 is
// then r − 1.

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
	var t, n, i, cnt, idx uint
	var a, b []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]uint, n)
			b = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
			}

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &b[i])
			}

			cnt = 0
			idx = n

			for i = n; i > 0 && idx > 0; i-- {
				cnt++

				for idx > 0 && a[idx-1] != b[i-1] {
					idx--
				}
			}

			if idx == 0 {
				fmt.Fprintln(writer, n-cnt+1)
			} else {
				fmt.Fprintln(writer, n-cnt)
			}
		}
	}

	writer.Flush()
}
