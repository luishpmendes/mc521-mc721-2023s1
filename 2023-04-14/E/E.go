// https://codeforces.com/problemset/problem/1385/D

// CodeForces 1385D - a-Good String

// You are given a string s[1…n] consisting of lowercase Latin letters. It is guaranteed that n=2k for some integer k≥0.
// The string s[1…n] is called c-good if at least one of the following three conditions is satisfied:
//  • The length of s is 1, and it consists of the character c (i.e. s1=c);
//  • The length of s is greater than 1, the first half of the string consists of only the character c (i.e. s1=s2=⋯=sn2=c) and the second half of the string (i.e. the string sn2+1sn2+2…sn) is a (c+1)-good string;
//  • The length of s is greater than 1, the second half of the string consists of only the character c (i.e. sn2+1=sn2+2=⋯=sn=c) and the first half of the string (i.e. the string s1s2…sn2) is a (c+1)-good string.
// For example: "aabc" is 'a'-good, "ffgheeee" is 'e'-good.
// In one move, you can choose one index i from 1 to n and replace si with any lowercase Latin letter (any character from 'a' to 'z').
// Your task is to find the minimum number of moves required to obtain an 'a'-good string from s (i.e. c-good string for c= 'a'). It is guaranteed that the answer always exists.
// You have to answer t independent test cases.
// Another example of an 'a'-good string is as follows. Consider the string s="cdbbaaaa". It is an 'a'-good string, because:
//  • the second half of the string ("aaaa") consists of only the character 'a';
//  • the first half of the string ("cdbb") is 'b'-good string, because:
//     • the second half of the string ("bb") consists of only the character 'b';
// 	• the first half of the string ("cd") is 'c'-good string, because:
// 	   • the first half of the string ("c") consists of only the character 'c';
// 	   • the second half of the string ("d") is 'd'-good string.

// Input
// The first line of the input contains one integer t (1≤t≤2⋅10^4) — the number of test cases. Then t test cases follow.
// The first line of the test case contains one integer n (1≤n≤131072) — the length of s. It is guaranteed that n=2^k for some integer k≥0. The second line of the test case contains the string s consisting of n lowercase Latin letters.
// It is guaranteed that the sum of n does not exceed 2⋅10^5 (∑n≤2⋅10^5).

// Output
// For each test case, print the answer — the minimum number of moves required to obtain an 'a'-good string from s (i.e. c-good string with c= 'a'). It is guaranteed that the answer exists.

// Example
// input
// 6
// 8
// bbdcaaaa
// 8
// asdfghjk
// 8
// ceaaaabb
// 8
// bbaaddcc
// 1
// z
// 2
// ac
// output
// 0
// 7
// 4
// 5
// 1
// 1

// Tutorial
// Consider the problem in 0-indexation. Define the function calc(l,r,c) which finds the minimum number of changes to make the string s[l…r) c-good string. Let mid=(l+r)/2. Then let cntl=(r−l)/2−count(s[l…mid),c)+calc(mid,r,c+1) and cntr=(r−l)/2−count(s[mid…r),c)+calc(l,mid,c+1), where count(s,c) is the number of occurrences of the character c in s. We can see that cntl describes the second condition from the statement and cntr describes the third one. So, calc(l,r,c) returns min(cntl,cntr) except one case. When r−l=1, we need to return 1 if sl≠c and 0 otherwise. This function works in O(nlogn) (each element of s belongs to exactly logn segments, like segment tree). You can get the answer if you run calc(0,n, ′a′).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func calc(s *string, l uint, r uint, c byte) uint {
	var mid, i, cntl, cntr uint

	if l == r {
		if (*s)[l] != c {
			return 1
		} else {
			return 0
		}
	} else {
		mid = (l + r) >> 1

		for i = l; i <= r; i++ {
			if (*s)[i] != c {
				if i <= mid {
					cntl++
				} else {
					cntr++
				}
			}
		}

		cntl += calc(s, mid+1, r, c+1)
		cntr += calc(s, l, mid, c+1)

		if cntl < cntr {
			return cntl
		} else {
			return cntr
		}
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n uint
	var s string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)
			fmt.Fprintln(writer, calc(&s, 0, n-1, 'a'))
		}
	}

	writer.Flush()
}
