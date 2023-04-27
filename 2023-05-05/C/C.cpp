// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=117&page=show_problem&problem=2890

// UVA 11790 - Murcia's Skyline

// Murcia’s skyline is growing up very fast. Since the 15th century, it used to be dominated by the profile
// of its Baroque Cathedral. But nowadays, new skyscrapers are rising in Murcian huerta.
// Some people say that if look at the skyline from left to right, you can observe and increasing profile;
// but other people say the profile is decreasing.
// Looking at Murcia’s skyline from left to right, we have a series of N buildings. Each building has
// its own height and width. You have to discover if the skyline is increasing or decreasing.
// We say the skyline is increasing if the longest increasing subsequence of buildings is bigger or
// equal than the longest decreasing subsequence of buildings; in other case, we say it is decreasing. A
// subsequence is a subset of the original sequence, in the same order. The length of a subsequence of
// buildings is the sum of the widths of its elements.
// For example, assuming we have six buildings of heights: 10, 100, 50, 30, 80, 10; and widths: 50,
// 10, 10, 15, 20, 10; then we have an increasing subsequence of 3 buildings and total length 85, and a
// decreasing subsequence of 1 building and total length 50 (also, there is a decreasing subsequence of
// 4 buildings and length 45). So, in this case, we say that the skyline is increasing. You can see this
// example below.

// Input

// The first line of the input contains an integer indicating the number of test cases.
// For each test case, the first line contains a single integer, N, indicating the number of buildings of
// the skyline. Then, there are two lines, each with N integers separated by blank spaces. The first line
// indicates the heights of the buildings, from left to right. The second line indicates the widths of the
// buildings, also from left to right.

// Output

// For each test case, the output should contain a line. If the skyline is increasing, the format will be:
// Case i. Increasing (A). Decreasing (B).
// If the skyline is decreasing, the format will be:
// Case i. Decreasing (B). Increasing (A).
// where i is the number of the corresponding test case (starting with 1), A is the length of the longest
// increasing subsequence, and B is the length of the longest decreasing subsequence.

// Sample Input

// 3
// 6
// 10 100 50 30 80 10
// 50 10 10 15 20 10
// 4
// 30 20 20 10
// 20 30 40 50
// 3
// 80 80 80
// 15 25 20

// Sample Output

// Case 1. Increasing (85). Decreasing (50).
// Case 2. Decreasing (110). Increasing (50).
// Case 3. Increasing (25). Decreasing (25).

// combination of LIS+LDS, weighted

#include <algorithm>
#include <iostream>
#include <vector>

unsigned longest_increasing_subsequence(const std::vector<unsigned> & height, const std::vector<unsigned> & width) {
    std::vector<unsigned> lis(height.size());

    lis[0] = width[0];

    for (unsigned i = 1; i < height.size(); ++i) {
        lis[i] = width[i];

        for (unsigned j = 0; j < i; ++j) {
            if (height[j] < height[i] && lis[i] < lis[j] + width[i]) {
                lis[i] = lis[j] + width[i];
            }
        }
    }

    return *std::max_element(lis.begin(), lis.end());
}

int main() {
    unsigned t, c, n, lis, lds;
    std::vector<unsigned> height, width;

    while (std::cin >> t) {
        for (c = 1; c <= t; c++) {
            std::cin >> n;
            height.resize(n);
            width.resize(n);

            for (unsigned & h : height) {
                std::cin >> h;
            }

            for (unsigned & w : width) {
                std::cin >> w;
            }

            lis = longest_increasing_subsequence(height, width);
            std::reverse(height.begin(), height.end());
            std::reverse(width.begin(), width.end());
            lds = longest_increasing_subsequence(height, width);

            if (lis >= lds) {
                std::cout << "Case " << c << ". Increasing (" << lis << "). Decreasing (" << lds << ").\n";
            } else {
                std::cout << "Case " << c << ". Decreasing (" << lds << "). Increasing (" << lis << ").\n";
            }
        }
    }

    return 0;
}
