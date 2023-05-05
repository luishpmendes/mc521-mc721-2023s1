// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=24&page=show_problem&problem=2171

// UVA 11230 - Annoying painting tool

// Maybe you wonder what an annoying painting tool is? First of all, the painting tool we speak of
// supports only black and white. Therefore, a picture consists of a rectangular area of pixels, which are
// either black or white. Second, there is only one operation how to change the colour of pixels:
// Select a rectangular area of r rows and c columns of pixels, which is completely inside the picture.
// As a result of the operation, each pixel inside the selected rectangle changes its colour (from black to
// white, or from white to black).
// Initially, all pixels are white. To create a picture, the operation described above can be applied
// several times. Can you paint a certain picture which you have in mind?

// Input

// The input contains several test cases. Each test case starts with one line containing four integers n, m,
// r and c. (1 ≤ r ≤ n ≤ 100, 1 ≤ c ≤ m ≤ 100), The following n lines each describe one row of pixels
// of the painting you want to create. The i-th line consists of m characters describing the desired pixel
// values of the i-th row in the finished painting (‘0’ indicates white, ‘1’ indicates black).
// The last test case is followed by a line containing four zeros.

// Output

// For each test case, print the minimum number of operations needed to create the painting, or ‘-1’ if it
// is impossible.

// Sample Input

// 3 3 1 1
// 010
// 101
// 010
// 4 3 2 1
// 011
// 110
// 011
// 110
// 3 4 2 2
// 0110
// 0111
// 0000
// 0 0 0 0

// Sample Output

// 4
// 6
// -1

#include <iostream>
#include <vector>
#include <string>

int main() {
    unsigned n, m, r, c, i, j, k, l;
    int ans;
    std::vector<std::vector<bool>> pixel;
    std::string s;

    while (std::cin >> n >> m >> r >> c && (n || m || r || c)) {
        pixel.resize(n, std::vector<bool>(m));
        pixel.assign(n, std::vector<bool>(m, false));

        for (i = 0; i < n; i++) {
            std::cin >> s;

            for (j = 0; j < m; j++) {
                pixel[i][j] = s[j] == '1';
            }
        }

        ans = 0;

        for (i = 0; i + r < n + 1; i++) {
            for (j = 0; j + c < m + 1; j++) {
                if (pixel[i][j]) {
                    for (k = 0; k < r; k++) {
                        for (l = 0; l < c; l++) {
                            pixel[i + k][j + l] = pixel[i + k][j + l] ^ true;
                        }
                    }

                    ans++;
                }
            }
        }

        for (i = 0; i < n && ans >= 0; i++) {
            for (j = 0; j < m && ans >= 0; j++) {
                if (pixel[i][j]) {
                    ans = -1;
                }
            }
        }

        std::cout << ans << std::endl;
    }

    return 0;
}
