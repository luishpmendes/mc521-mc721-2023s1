// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=23&page=show_problem&problem=2098

// UVA 11157 - Dynamic Frog

// With the increased use of pesticides, the local streams and rivers have
// become so contaminated that it has become almost impossible for the
// aquatic animals to survive.
// Frog Fred is on the left bank of such a river. N rocks are arranged
// in a straight line from the left bank to the right bank. The distance
// between the left and the right bank is D meters. There are rocks of
// two sizes. The bigger ones can withstand any weight but the smaller
// ones start to drown as soon as any mass is placed on it. Fred has to go
// to the right bank where he has to collect a gift and return to the left
// bank where his home is situated.
// He can land on every small rock at most one time, but can use the
// bigger ones as many times as he likes. He can never touch the polluted
// water as it is extremely contaminated.
// Can you plan the itinerary so that the maximum distance of a single leap is minimized?

// Input

// The first line of input is an integer T (T < 100) that indicates the number of test cases. Each case
// starts with a line containing two integers N (0 ≤ N ≤ 100) and D (1 ≤ D ≤ 1000000000). The next
// line gives the description of the N stones. Each stone is defined by S − M. S indicates the type Big(B)
// or Small(S) and M (0 < M < D) determines the distance of that stone from the left bank. The stones
// will be given in increasing order of M.

// Output

// For every case, output the case number followed by the minimized maximum leap.

// Sample Input

// 3
// 1 10
// B-5
// 1 10
// S-5
// 2 10
// B-3 S-6

// Sample Output

// Case 1: 5
// Case 2: 10
// Case 3: 7

#include <iostream>
#include <vector>
#include <string>
#include <utility>

int main() {
    unsigned t, c, n, d, i, m, ans, dist;
    char s, aux;
    std::vector<std::pair<char, unsigned>> stones;

    while (std::cin >> t) {
        for (c = 1; c <= t; c++) {
            std::cin >> n >> d;
            stones.resize(n+1);
            stones[0] = std::make_pair('B', 0);
            ans = 0;

            for (i = 1; i <= n; i++) {
                std::cin >> s >> aux >> m;
                stones[i] = std::make_pair(s, m);

                if (stones[i-1].first == 'B') {
                    dist = m - stones[i-1].second;
                } else {
                    dist = m - stones[i-2].second;
                }

                if (ans < dist) {
                    ans = dist;
                }
            }

            if (stones[n].first == 'B') {
                dist = d - stones[n].second;
            } else {
                dist = d - stones[n-1].second;
            }

            if (ans < dist) {
                ans = dist;
            }

            std::cout << "Case " << c << ": " << ans << std::endl;
        }
    }

    return 0;
}
