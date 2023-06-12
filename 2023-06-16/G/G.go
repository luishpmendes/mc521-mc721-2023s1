// https://codeforces.com/problemset/problem/1798/B

// CodeForces 1798B - Three Sevens

// Lottery "Three Sevens" was held for m days. On day i, ni people with the numbers ai,1,…,ai,ni participated in the lottery.
// It is known that in each of the m days, only one winner was selected from the lottery participants. The lottery winner on day i was not allowed to participate in the lottery in the days from i+1 to m.
// Unfortunately, the information about the lottery winners has been lost. You need to find any possible list of lottery winners on days from 1 to m or determine that no solution exists.

// Input
// Each test contains multiple test cases. The first line contains the number of test cases t (1≤t≤50000). The description of the test cases follows.
// The first line of each test case contains a single integer m (1≤m≤50000) — the number of days in which the lottery was held.
// Next, for each i from 1 to m, follows a two-line block of data.
// The first line of each block contains a single integer ni (1≤ni≤50000) — the number of lottery participants on day i.
// The second line of the block contains integers ai,1,…,ai,ni (1≤ai,j≤50000) — lottery participants on day i. It is guaranteed that all the numbers ai,1,…,ai,ni are pairwise distinct.
// It is guaranteed that the sum of ni over all blocks of all test cases does not exceed 50000.

// Output
// For each test case, if there is no solution, print a single integer −1.
// Otherwise, print m integers p1,p2,…,pm (1≤pi≤50000) — lottery winners on days from 1 to m. If there are multiple solutions, print any of them.

// Example
// input
// 3
// 3
// 4
// 1 2 4 8
// 3
// 2 9 1
// 2
// 1 4
// 2
// 2
// 1 2
// 2
// 2 1
// 4
// 4
// 1 2 3 4
// 1
// 1
// 1
// 4
// 1
// 3
// output
// 8 2 1
// -1
// 2 1 4 3

// Note
// In the first test case, one of the answers is [8,2,1] since the participant with the number 8 participated on day 1, but did not participate on days 2 and 3; the participant with the number 2 participated on day 2, but did not participate on day 3; and the participant with the number 1 participated on day 3. Note that this is not the only possible answer, for example, [8,9,4] is also a correct answer.
// In the second test case, both lottery participants participated on both days, so any possible lottery winner on the day 1 must have participated on the day 2, which is not allowed. Thus, there is no correct answer.
// In the third test case, only one participant participated on days 2, 3, 4, and on day 1 there is only one participant who did not participate in the lottery on days 2,3,4 — participant 2, which means [2,1,4,3] is the only correct answer to this test case.

// Hint 1
// Let the participant with the number X participate in the lottery in days i1<i2<…<ik. On what days could the participant X be chosen as the winner so as not to break the conditions?

// Hint 2
// If there are several candidates for the lottery winner on the day of i (who did not participate in the days from i+1 to m), does it matter which of them we choose as a winner?

// Tutorial
// Let's calculate the array last, where lastX is the last day of the lottery in which the person X participated. Then the only day when X could be a winner is the day lastX. Then on the day of i, only the person with lastX=i can be the winner. It is also clear that if there are several such participants for the day i, you can choose any of them as the winner, since these participants cannot be winners on any other days. In total, we need to go through all the days, if for some day there are no participants with last equal to this day, then the answer is −1. Otherwise, we choose any participant with lastX=i as the winner on the day of i.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, m, n, i, j, max_id int
	var a [][]int
	var last, ans []int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			max_id = 0
			fmt.Fscan(reader, &m)
			a = make([][]int, m)
			ans = make([]int, m)

			for i = 0; i < m; i++ {
				fmt.Fscan(reader, &n)
				a[i] = make([]int, n)

				for j = 0; j < len(a[i]); j++ {
					fmt.Fscan(reader, &a[i][j])

					if max_id < a[i][j] {
						max_id = a[i][j]
					}
				}

				ans[i] = -1
			}

			last = make([]int, max_id+1)

			for i = 0; i < m; i++ {
				for j = 0; j < len(a[i]); j++ {
					last[a[i][j]] = i
				}
			}

			for i = 0; i < m; i++ {
				for j = 0; j < len(a[i]); j++ {
					if last[a[i][j]] == i {
						ans[i] = a[i][j]
					}
				}

				if ans[i] == -1 {
					fmt.Fprintln(writer, -1)
					break
				}
			}

			if i == m {
				for i = 0; i < m; i++ {
					fmt.Fprint(writer, ans[i], " ")
				}

				fmt.Fprintln(writer)
			}
		}
	}

	writer.Flush()
}
