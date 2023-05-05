// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=279&page=show_problem&problem=3836

// UVA 12405 - Scarecrow

// Taso owns a very long field. He plans to grow different types of crops in
// the upcoming growing season. The area, however, is full of crows and
// Taso fears that they might feed on most of the crops. For this reason,
// he has decided to place some scarecrows at different locations of the
// field.
// The field can be modeled as a 1 × N grid. Some parts of the field
// are infertile and that means you cannot grow any crops on them. A
// scarecrow, when placed on a spot, covers the cell to its immediate left
// and right along with the cell it is on.
// Given the description of the field, what is the minimum number of
// scarecrows that needs to be placed so that all the useful section of the
// field is covered? Useful section refers to cells where crops can be grown.

// Input

// Input starts with an integer T (≤ 100), denoting the number of test cases.
// Each case starts with a line containing an integer N (0 < N < 100). The next line contains N
// characters that describe the field. A dot (.) indicates a crop-growing spot and a hash (#) indicates an
// infertile region.

// Output

// For each case, output the case number first followed by the number of scarecrows that need to be placed.

// Sample Input

// 3
// 3
// .#.
// 11
// ...##....##
// 2
// ##

// Sample Output

// Case 1: 1
// Case 2: 3
// Case 3: 0

// simpler interval covering problem

#include <iostream>
#include <string>

int main() {
    unsigned t, c, n, scarecrows, i;
    std::string s;

    while (std::cin >> t) {
        for (c = 1; c <= t; c++) {
            std::cin >> n >> s;

            scarecrows = 0;

            for (i = 0; i < n; i++) {
                if (s[i] == '.') {
                    scarecrows++;
                    i += 2;
                }
            }

            std::cout << "Case " << c << ": " << scarecrows << std::endl;
        }
    }

    return 0;
}
