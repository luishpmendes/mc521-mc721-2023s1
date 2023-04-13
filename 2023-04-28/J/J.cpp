// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=22&page=show_problem&problem=1967

// 11026 - A Grouping Problem

// You are given a set of N integers. You can take K different elements from them to make a group. Two
// groups will be different if there is at least one element which is not common to both. For example, if
// there are 4 elements a, b, c, d and you are asked to take two elements then ab, ad, bc, cd are all valid
// and different groups. A grouping system is complete if for a particular K, number of different groups
// is the maximum. In the former case, {ab, bc, cd, bd, ad, ac} is a complete grouping system.
// For a particular complete grouping system, the fitness is calculated in the following way
// 1. Each group of a grouping system contributes a part the multiplication of all numbers of that
// group
// 2. Contribution from all groups are added
// 3. The fitness is equivalent to T otal Contribution mod M, M is the bounding parameter
// In our example, for K = 2, the fitness is F2 = (ab + bc + cd + bd + ad + ac) mod M. If K = 1, then
// fitness is F1 = (a + b + c + d) mod M.
// Here, in this problem you have to find the complete grouping system with maximum fitness.

// Input

// Each test case starts with two positive integer N (2 ≤ N ≤ 1000) and M (1 ≤ M < 2^31). In next few
// lines there will be N positive integers. Each integer will be at best 1000. Input will be terminated by
// a case where N = M = 0.

// Output

// For each test case, print in a line the maximum fitness possible for a grouping system.

// Sample Input

// 4 10
// 1 2 3 4
// 4 100
// 1 2 3 4
// 4 6
// 1 2 3 4
// 0 0

// Sample Output

// 5
// 50
// 5

// DP, similar idea with binomial theorem in Section 5.4

// Summary:
// We are given a1,a2,…,an and a number K which defines the sum of products of permutation of length K, e.g. K=1 is a1+a2+…+an, K=2 is a1a2+a2a3+…+an−1an and so on and so forth. Find the maximum of these sum of products modM.

// Solution:
// Despite the technicality, this problem has a very nice recursive formulation:
// Let S(k,n) be the sum of products of permutation of length k amongst n elements. Then we have S(k,n)=an∗S(k−1,n−1)+S(k,n−1). This is actually derived from the factorization of an from all the terms that contain an. From here, we can do a bottom-up DP to arrive at the answer. Here is a sample implementation:

#include <iostream>
#include <vector>

int main () {
    unsigned long n, m, i, j, ans;
    std::vector<unsigned long> a;
    std::vector<std::vector<unsigned long>> dp;

    while (std::cin >> n >> m && (n || m)) {
        a.resize(n+1);
        dp.resize(n+1, std::vector<unsigned long>(n+1, 1));
        dp.assign(n+1, std::vector<unsigned long>(n+1, 1));

        for (i = 1; i <= n; i++) {
            std::cin >> a[i];
        }

        for (i = 1; i <= n; i++) {
            for (j = i; j <= n; j++) {
                dp[i][j] = (a[j] * dp[i - 1][j - 1]) % m;

                if (j > i) {
                    dp[i][j] += dp[i][j - 1];
                    dp[i][j] %= m;
                }
            }
        }

        ans = 0;

        for (i = 1; i <= n; i++) {
            if (ans < dp[i][n]) {
                ans = dp[i][n];
            }
        }

        std::cout << ans << std::endl;
    }

    return 0;
}
