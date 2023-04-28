// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=21&page=show_problem&problem=1884

// UVA 10943 - How do you add?

// Larry is very bad at math — he usually uses a calculator, which worked
// well throughout college. Unforunately, he is now struck in a deserted
// island with his good buddy Ryan after a snowboarding accident.
// They’re now trying to spend some time figuring out some good
// problems, and Ryan will eat Larry if he cannot answer, so his fate is
// up to you!
// It’s a very simple problem — given a number N, how many ways
// can K numbers less than N add up to N?
// For example, for N = 20 and K = 2, there are 21 ways:
// 0+20
// 1+19
// 2+18
// 3+17
// 4+16
// 5+15
// ...
// 18+2
// 19+1
// 20+0

// Input

// Each line will contain a pair of numbers N and K. N and K will both be an integer from 1 to 100,
// inclusive. The input will terminate on 2 0’s.

// Output

// Since Larry is only interested in the last few digits of the answer, for each pair of numbers N and K,
// print a single number mod 1,000,000 on a single line.

// Sample Input

// 20 2
// 20 2
// 0 0

// Sample Output

// 21
// 21

// discussed

#include <iostream>
#include <vector>

int main() {
    unsigned n, k, i, j, split;
    std::vector<std::vector<unsigned>> dp(101, std::vector<unsigned>(101));

    // these are the base case
    for (i = 0; i <= 100; i++) {
        dp[i][1] = 1;
    }

    // these three nested loops form the correct topological order
    for (j = 1; j < 100; j++) {
        for (i = 0; i <= 100; i++) {
            for (split = 0; split + i <= 100; split++) {
                dp[i + split][j + 1] += dp[i][j];
                dp[i + split][j + 1] %= 1000000;
            }
        }
    }

    while (std::cin >> n >> k && (n || k)) {
        std::cout << dp[n][k] << std::endl;
    }

    return 0;
}
