// https://codeforces.com/problemset/problem/1690/C

// CodeForces 1690C - Restoring the Duration of Tasks

// Recently, Polycarp completed n successive tasks.
// For each completed task, the time si is known when it was given, no two tasks were given at the same time. Also given is the time fi when the task was completed. For each task, there is an unknown value di (di>0) — duration of task execution.
// It is known that the tasks were completed in the order in which they came. Polycarp performed the tasks as follows:
//  • As soon as the very first task came, Polycarp immediately began to carry it out.
//  • If a new task arrived before Polycarp finished the previous one, he put the new task at the end of the queue.
//  • When Polycarp finished executing the next task and the queue was not empty, he immediately took a new task from the head of the queue (if the queue is empty — he just waited for the next task).
// Find di (duration) of each task.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of test cases.
// The descriptions of the input data sets follow.
// The first line of each test case contains one integer n (1≤n≤2⋅10^5).
// The second line of each test case contains exactly n integers s1<s2<⋯<sn (0≤si≤10^9).
// The third line of each test case contains exactly n integers f1<f2<⋯<fn (si<fi≤10^9).
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅10^5.

// Output
// For each of t test cases print n positive integers d1,d2,…,dn — the duration of each task.

// Example
// input
// 4
// 3
// 0 3 7
// 2 10 11
// 2
// 10 15
// 11 16
// 9
// 12 16 90 195 1456 1569 3001 5237 19275
// 13 199 200 260 9100 10000 10914 91066 5735533
// 1
// 0
// 1000000000
// output
// 2 7 1
// 1 1
// 1 183 1 60 7644 900 914 80152 5644467
// 1000000000

// Note
// First test case:
// The queue is empty at the beginning: []. And that's where the first task comes in. At time 2, Polycarp finishes doing the first task, so the duration of the first task is 2. The queue is empty so Polycarp is just waiting.
// At time 3, the second task arrives. And at time 7, the third task arrives, and now the queue looks like this: [7].
// At the time 10, Polycarp finishes doing the second task, as a result, the duration of the second task is 7.
// And at time 10, Polycarp immediately starts doing the third task and finishes at time 11. As a result, the duration of the third task is 1.

// Tutorial
// Accordingly, as it was said in the task, we put all tasks into the queue in the order of their arrival, then we fix the time at the beginning as cur_time=0. So, while there is a task in the queue, we proceed as follows:
//  1. Take the task from the queue.
//  2. Take as time the maximum from the current and from the arrival time of the task (cur_time=max(cur_time,s)).
//  3. We subtract the current time from the time when the task was done (f=d−cur_time).
//  4. Replace the current time with the time the task was done (cur_time=d)
//  5. If there is a task in the queue, go to item 1.

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
	var t, n, i, cur_time uint
	var s, f, d []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			s = make([]uint, n)
			f = make([]uint, n)
			d = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &s[i])
			}

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &f[i])
			}

			cur_time = 0

			for i = 0; i < n; i++ {
				if cur_time < s[i] {
					cur_time = s[i]
				}

				d[i] = f[i] - cur_time
				cur_time = f[i]
			}

			for i = 0; i < n-1; i++ {
				fmt.Fprint(writer, d[i], " ")
			}

			fmt.Fprintln(writer, d[n-1])
		}
	}

	writer.Flush()
}
