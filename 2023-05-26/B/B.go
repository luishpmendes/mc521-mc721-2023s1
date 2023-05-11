// https://codeforces.com/problemset/problem/1560/A

// CodeForces 1560A - Dislike of Threes

// Polycarp doesn't like integers that are divisible by 3 or end with the digit 3 in their decimal representation. Integers that meet both conditions are disliked by Polycarp, too.
// Polycarp starts to write out the positive (greater than 0) integers which he likes: 1,2,4,5,7,8,10,11,14,16,…. Output the k-th element of this sequence (the elements are numbered from 1).

// Input
// The first line contains one integer t (1≤t≤100) — the number of test cases. Then t test cases follow.
// Each test case consists of one line containing one integer k (1≤k≤1000).

// Output
// For each test case, output in a separate line one integer x — the k-th element of the sequence that was written out by Polycarp.

// Example
// input
// 10
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 1000
// output
// 1
// 2
// 4
// 5
// 7
// 8
// 10
// 11
// 14
// 1666

// Tutorial
// The solution is simple: let's create an integer variable (initially set to 0) that will contain the number of considered liked integers. Let's iterate over all positive integers starting with 1. Let's increase the variable only when the considered number is liked. If the variable is equal to k, let's stop the iteration and output the last considered number.
// Since the answer for k=1000 is x=1666, the count of considered numbers is at most 1666 so the solution will work on the specified limitations fast enough.

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
	var t, k, i uint
	var v []uint

	v = make([]uint, 0)
	v = append(v, 0)

	for i = 1; len(v) < 1010; i++ {
		if i%3 != 0 && i%10 != 3 {
			v = append(v, i)
		}
	}

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &k)
			fmt.Fprintln(writer, v[k])
		}
	}

	writer.Flush()
}
