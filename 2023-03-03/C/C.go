// https://codeforces.com/problemset/problem/1765/D

// CodeForces 1765D - Watch the Videos

// Monocarp wants to watch n videos. Each video is only one minute long, but its size may be arbitrary. The i-th video has the size ai megabytes. All videos are published on the Internet. A video should be downloaded before it can be watched. Monocarp has poor Internet connection — it takes exactly 1 minute to download 1 megabyte of data, so it will require ai minutes to download the i-th video.
// Monocarp's computer has a hard disk of m megabytes. The disk is used to store the downloaded videos. Once Monocarp starts the download of a video of size s, the s megabytes are immediately reserved on a hard disk. If there are less than s megabytes left, the download cannot be started until the required space is freed. Each single video can be stored on the hard disk, since ai≤m for all i. Once the download is started, it cannot be interrupted. It is not allowed to run two or more downloads in parallel.
// Once a video is fully downloaded to the hard disk, Monocarp can watch it. Watching each video takes exactly 1 minute and does not occupy the Internet connection, so Monocarp can start downloading another video while watching the current one.
// When Monocarp finishes watching a video, he doesn't need it on the hard disk anymore, so he can delete the video, instantly freeing the space it occupied on a hard disk. Deleting a video takes negligible time.
// Monocarp wants to watch all n videos as quickly as possible. The order of watching does not matter, since Monocarp needs to watch all of them anyway. Please calculate the minimum possible time required for that.

// Input
// The first line contains two integers n and m (1≤n≤2⋅10^5; 1≤m≤10^9) — the number of videos Monocarp wants to watch and the size of the hard disk, respectively.
// The second line contains n integers a1,a2,…,an (1≤ai≤m) — the sizes of the videos.

// Output
// Print one integer — the minimum time required to watch all n videos.

// Examples

// input
// 5 6
// 1 2 3 4 5
// output
// 16

// input
// 5 5
// 1 2 3 4 5
// output
// 17

// input
// 4 3
// 1 3 2 3
// output
// 12

// Tutorial
// In this solution, we assume that the sequence a1,a2,…,an is sorted (if it is not — just sort it before running the solution).
// Suppose we download and watch the videos in some order. The answer to the problem is n+∑ai, reduced by 1 for every pair of adjacent videos that can fit onto the hard disk together (i. e. their total size is not greater than m), because for every such pair, we can start downloading the second one while watching the first one. So, we need to order the videos in such a way that the number of such pairs is the maximum possible.
// Suppose we want to order them so that every pair of adjacent videos is "good". We need to pick the ordering that minimizes the maximum sum of adjacent elements. There are multiple ways to construct this ordering; one of them is [an,a1,an−1,a2,an−2,a3,…], and the maximum sum of adjacent elements will be maxni=1,2i≠nai+an+1−i.
// Proof that this ordering is optimal starts here
// Suppose j is such value of i that ai+an+1−i is the maximum, and j<n+1−j. Let's prove that we cannot make the maximum sum of adjacent elements less than aj+an+1−j. There are at most j−1 values in a which are less than aj (let's call them Group A), and at least j values that are not less than aj (let's call them Group B). If we want each sum of adjacent elements to be less than aj+an+1−j, we want the elements of Group B to be adjacent only to the elements of the Group A. The elements of Group B have at least 2j−2 links to the neighbors (since at most two of them will have only one neighbor), the elements of Group A have at most 2(j−1) links to the neighbors, but we cannot link Group A with Group B without any outside links since it will disconnect them from all other elements of the array. So, we cannot get the maximum sum less than aj+an+1−j.
// Proof that this ordering is optimal ends here
// Okay, now what if we cannot make all pairs of adjacent elements "good"? We can run binary search on the number of pairs that should be "bad". Let this number be k, then if we need to make at most k pairs "bad", we can check that it's possible by removing k maximum elements from a and checking that now we can make the array "good". So, in total, our solution will work in O(nlogn).

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
	var t, n, m, i, j, ans int
	var a []int

	for t, _ = fmt.Fscan(reader, &n, &m); t == 2; t, _ = fmt.Fscan(reader, &n, &m) {
		a = make([]int, n)
		ans = 0

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			ans += a[i]
		}

		ans++
		sort.Ints(a)

		for i, j = 0, n-1; i < j; i, j = i+1, j-1 {
			for j > i && a[i]+a[j] > m {
				j--
				ans++
			}
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
