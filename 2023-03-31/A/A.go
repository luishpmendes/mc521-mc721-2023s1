// https://codeforces.com/problemset/problem/1520/A

// CodeForces 1520A - Do Not Be Distracted!

// Polycarp has 26 tasks. Each task is designated by a capital letter of the Latin alphabet.
// The teacher asked Polycarp to solve tasks in the following way: if Polycarp began to solve some task, then he must solve it to the end, without being distracted by another task. After switching to another task, Polycarp cannot return to the previous task.
// Polycarp can only solve one task during the day. Every day he wrote down what task he solved. Now the teacher wants to know if Polycarp followed his advice.
// For example, if Polycarp solved tasks in the following order: "DDBBCCCBBEZ", then the teacher will see that on the third day Polycarp began to solve the task 'B', then on the fifth day he got distracted and began to solve the task 'C', on the eighth day Polycarp returned to the task 'B'. Other examples of when the teacher is suspicious: "BAB", "AABBCCDDEEBZZ" and "AAAAZAAAAA".
// If Polycarp solved the tasks as follows: "FFGZZZY", then the teacher cannot have any suspicions. Please note that Polycarp is not obligated to solve all tasks. Other examples of when the teacher doesn't have any suspicious: "BA", "AFFFCC" and "YYYYY".
// Help Polycarp find out if his teacher might be suspicious.

// Input
// The first line contains an integer t (1≤t≤1000). Then t test cases follow.
// The first line of each test case contains one integer n (1≤n≤50) — the number of days during which Polycarp solved tasks.
// The second line contains a string of length n, consisting of uppercase Latin letters, which is the order in which Polycarp solved the tasks.

// Output
// For each test case output:
//  • "YES", if the teacher cannot be suspicious;
//  • "NO", otherwise.
// You may print every letter in any case you want (so, for example, the strings yEs, yes, Yes and YES are all recognized as positive answer).

// Example
// input
// 5
// 3
// ABA
// 11
// DDBBCCCBBEZ
// 7
// FFGZZZY
// 1
// Z
// 2
// AB
// output
// NO
// NO
// YES
// YES
// YES

// Tutorial
// The simplest solution — go through the problem, because of which the teacher might have suspicions. Now you can find the first day when Polycarp solved this problem and the last such day. Between these two days, all problems should be the same. If this is not the case, the answer is "NO".

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
	var t, n, i uint
	var tasks string
	var last_task byte
	var began map[byte]bool
	var ans bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &tasks)
			began = make(map[byte]bool)
			last_task = tasks[0]
			began[last_task] = true
			ans = true

			for i = 1; i < n && ans; i++ {
				if tasks[i] != last_task {
					if began[tasks[i]] {
						ans = false
					} else {
						began[tasks[i]] = true
						last_task = tasks[i]
					}
				}
			}

			if ans {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}

	writer.Flush()
}
