// https://codeforces.com/problemset/problem/1765/B

// CodeForces 1765B - Broken Keyboard

// Recently, Mishka started noticing that his keyboard malfunctions — maybe it's because he was playing rhythm games too much. Empirically, Mishka has found out that every other time he presses a key, it is registered as if the key was pressed twice. For example, if Mishka types text, the first time he presses a key, exactly one letter is printed; the second time he presses a key, two same letters are printed; the third time he presses a key, one letter is printed; the fourth time he presses a key, two same letters are printed, and so on. Note that the number of times a key was pressed is counted for the whole keyboard, not for each key separately. For example, if Mishka tries to type the word osu, it will be printed on the screen as ossu.
// You are given a word consisting of n lowercase Latin letters. You have to determine if it can be printed on Mishka's keyboard or not. You may assume that Mishka cannot delete letters from the word, and every time he presses a key, the new letter (or letters) is appended to the end of the word.

// Input
// The first line of the input contains one integer t (1≤t≤100) — the number of test cases.
// The first line of the test case contains one integer n (1≤n≤100) — the length of the word.
// The second line of the test case contains a string s consisting of n lowercase Latin letters — the word that should be checked.

// Output
// For each test case, print YES if the word s can be printed on Mishka's keyboard, and NO otherwise.

// Examples
// input
// 4
// 4
// ossu
// 2
// aa
// 6
// addonn
// 3
// qwe
// output
// YES
// NO
// YES
// NO

// Note
// In the first test case, Mishka can type the word as follows: press o (one letter o appears at the end of the word), then presses s (two letters s appear at the end of the word), and, finally, press u (one letter appears at the end of the word, making the resulting word ossu).
// In the second test case, Mishka can try typing the word as follows: press a (one letter a appears at the end of the word). But if he tries to press a one more time, two letters a will appear at the end of the word, so it is impossible to print the word using his keyboard.
// In the fourth test case, Mishka has to start by pressing q. Then, if he presses w, two copies of w will appear at the end of the word, but the third letter should be e instead of w, so the answer is NO.

// Tutorial
// There are many ways to solve this problem. Basically, we need to check two conditions.
// The first one is the condition on the number of characters: nmod3≠2, since after the first key press, we get the remainder 1 modulo 3, after the second key press, we get the remainder 0 modulo 3, then 1 again, then 0 — and so on, and we cannot get the remainder 2.
// Then we need to check that, in each pair of characters which appeared from the same key press, these characters are the same — that is, s2=s3, s5=s6, s8=s9, and so on.

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
	var s string
	var ans bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)
			ans = true

			if n%3 == 2 {
				ans = false
			}

			for i = 2; i < n && ans; i += 3 {
				if s[i] != s[i-1] {
					ans = false
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
