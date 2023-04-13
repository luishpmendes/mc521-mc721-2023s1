// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=1071

// 10130 - SuperSale

// There is a SuperSale in a SuperHiperMarket. Every person can take only one object of each kind, i.e.
// one TV, one carrot, but for extra low price. We are going with a whole family to that SuperHiperMarket.
// Every person can take as many objects, as he/she can carry out from the SuperSale. We have given list
// of objects with prices and their weight. We also know, what is the maximum weight that every person
// can stand. What is the maximal value of objects we can buy at SuperSale?

// Input

// The input consists of T test cases. The number of them (1 ≤ T ≤ 1000) is given on the first line of
// the input file. Each test case begins with a line containing a single integer number N that indicates
// the number of objects (1 ≤ N ≤ 1000). Then follows N lines, each containing two integers: P and W.
// The first integer (1 ≤ P ≤ 100) corresponds to the price of object. The second integer (1 ≤ W ≤ 30)
// corresponds to the weight of object. Next line contains one integer (1 ≤ G ≤ 100) its the number of
// people in our group. Next G lines contains maximal weight (1 ≤ MW ≤ 30) that can stand this i-th
// person from our family (1 ≤ i ≤ G).

// Output

// For every test case your program has to determine one integer. Print out the maximal value of goods
// which we can buy with that family.

// Sample Input

// 2
// 3
// 72 17
// 44 23
// 31 24
// 1
// 26
// 6
// 64 26
// 85 22
// 52 4
// 99 18
// 39 13
// 54 9
// 4
// 23
// 20
// 20
// 26

// Sample Output

// 72
// 514

// discussed

#include <algorithm>
#include <iostream>
#include <limits>
#include <vector>

unsigned val(const unsigned & n, const std::vector<unsigned> & p, const std::vector<unsigned> & w, std::vector<std::vector<unsigned>> & memo, unsigned id, unsigned remW) {
    if (id == n || remW == 0) {
        return 0;
    }

    if (memo[id][remW] != std::numeric_limits<unsigned>::max()) {
        return memo[id][remW];
    }

    if (w[id] > remW) {
        return memo[id][remW] = val(n, p, w, memo, id + 1, remW);
    }

    return memo[id][remW] = std::max(val(n, p, w, memo, id + 1, remW), p[id] + val(n, p, w, memo, id + 1, remW - w[id]));
}

int main () {
    unsigned t, n, i, g, mw, ans;
    std::vector<unsigned> p, w;
    std::vector<std::vector<unsigned>> dp;

    while (std::cin >> t) {
        while(t--) {
            std::cin >> n;
            p.resize(n);
            w.resize(n);
            dp.resize(n+1, std::vector<unsigned>(31, std::numeric_limits<unsigned>::max()));
            dp.assign(n+1, std::vector<unsigned>(31, std::numeric_limits<unsigned>::max()));

            for (i = 0; i < n; i++) {
                std::cin >> p[i] >> w[i];
            }

            std::cin >> g;
            ans = 0;

            while (g--) {
                std::cin >> mw;
                ans += val(n, p, w, dp, 0, mw);
            }

            std::cout << ans << std::endl;
        }
    }

    return 0;
}
