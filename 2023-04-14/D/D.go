// https://codeforces.com/problemset/problem/1154/A

// CodeForces 1154A - Restoring Three Numbers

// Polycarp has guessed three positive integers a, b and c. He keeps these numbers in secret, but he writes down four numbers on a board in arbitrary order — their pairwise sums (three numbers) and sum of all three numbers (one number). So, there are four numbers on a board in random order: a+b, a+c, b+c and a+b+c.
// You have to guess three numbers a, b and c using given numbers. Print three guessed integers in any order.
// Pay attention that some given numbers a, b and c can be equal (it is also possible that a=b=c).

// Input
// The only line of the input contains four positive integers x1,x2,x3,x4 (2≤xi≤10^9) — numbers written on a board in random order. It is guaranteed that the answer exists for the given number x1,x2,x3,x4.

// Output
// Print such positive integers a, b and c that four numbers written on a board are values a+b, a+c, b+c and a+b+c written in some order. Print a, b and c in any order. If there are several answers, you can print any. It is guaranteed that the answer exists.

// Examples

// input
// 3 6 5 4
// output
// 2 1 3

// input
// 40 40 40 60
// output
// 20 20 20

// input
// 201 101 101 200
// output
// 1 100 100

// Tutorial
// Let x1=a+b, x2=a+c, x3=b=c and x4=a+b+c. Then we can construct the following answer: c=x4−x1, b=x4−x2 and a=x4−x3.
// Because all numbers in the answer are positive, we can assume that the maximum element of x is a+b+c. So let's sort the input array x consisting of four elements and just print x4−x1,x4−x2 and x4−x3.

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
	var t int
	var x [4]uint

	for t, _ = fmt.Fscan(reader, &x[0], &x[1], &x[2], &x[3]); t == 4; t, _ = fmt.Fscan(reader, &x[0], &x[1], &x[2], &x[3]) {
		sort.Slice(x[:], func(i, j int) bool { return x[i] < x[j] })
		fmt.Fprintln(writer, x[3]-x[0], x[3]-x[1], x[3]-x[2])
	}

	writer.Flush()
}
