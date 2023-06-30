// https://codeforces.com/problemset/problem/1809/B

// CodeForces 1809B - Points on Plane
// You are given a two-dimensional plane, and you need to place n chips on it.
// You can place a chip only at a point with integer coordinates. The cost of placing a chip at the point (x,y) is equal to |x|+|y| (where |a| is the absolute value of a).
// The cost of placing n chips is equal to the maximum among the costs of each chip.
// You need to place n chips on the plane in such a way that the Euclidean distance between each pair of chips is strictly greater than 1, and the cost is the minimum possible.

// Input
// The first line contains one integer t (1≤t≤10^4) — the number of test cases. Next t cases follow.
// The first and only line of each test case contains one integer n (1≤n≤10^18) — the number of chips you need to place.

// Output
// For each test case, print a single integer — the minimum cost to place n chips if the distance between each pair of chips must be strictly greater than 1.

// Example
// input
// 4
// 1
// 3
// 5
// 975461057789971042
// output
// 0
// 1
// 2
// 987654321

// Note
// In the first test case, you can place the only chip at point (0,0) with total cost equal to 0+0=0.
// In the second test case, you can, for example, place chips at points (−1,0), (0,1) and (1,0) with costs |−1|+|0|=1, |0|+|1|=1 and |0|+|1|=1. Distance between each pair of chips is greater than 1 (for example, distance between (−1,0) and (0,1) is equal to √2). The total cost is equal to max(1,1,1)=1.
// In the third test case, you can, for example, place chips at points (−1,−1), (−1,1), (1,1), (0,0) and (0,2). The total cost is equal to max(2,2,2,0,2)=2.

// Tutorial
// Suppose, the answer is k. What's the maximum number of chips we can place? Firstly, the allowed points (x,y) to place chips are such that |x|+|y|≤k. We can group them by x-coordinate: for x=k there is only one y=0, for x=k−1 possible y are −1,0,1; for x=k−2 possible y are in segment [−2,…,2] and so on. For x=0 possible y are in [−k,…,k]. The negative x-s are the same.
// Let's calculate the maximum number of chips we can place at each "row": for x=k it's 1; for x=k−1 there are three y-s, but since we can't place chips at the neighboring y-s, we can place at most 2 chips; for x=k−2 we have 5 places, but can place only 3 chips; for x=0 we have 2k+1 places, but can occupy only k+1 points.
// In total, for x∈[0,…,k] we can place at most 1+2+⋯+(k+1)=((k+1)(k+2))/2 chips. Analogically, for x∈[−k,…,−1] we can place at most 1+2+⋯+k=(k(k+1))/2 chips.
// In total, we can place at most ((k+1)(k+2))/2+(k(k+1))/2=(k+1)^2 chips with cost at most k. Note that (k+1)^2 can actually be reached since the distance between chips on the different rows is greater than 1.
// So, to solve our task, it's enough to find minimum k such that (k+1)^2≥n that can be done with Binary Search.
// Or we can calculate k=⌈√n⌉-1. Note that √n can lose precision, since n is cast to double before taking the square root (for example, 975461057789971042 transforms to 9.754610577899711⋅1017=975461057789971100 when converted to double). So you should either cast long long to long double (that consists of 80 bits in some C++ compilers) or check value k+1 as a possible answer.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, ans uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			ans = uint(math.Sqrt(float64(n)))

			for ans*ans > n {
				ans--
			}

			for ans*ans < n {
				ans++
			}

			fmt.Fprintln(writer, ans-1)
		}
	}

	writer.Flush()
}
