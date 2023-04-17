// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=119

// UVA 183 - Bit Maps

// The bitmap is a data structure that arises in many areas of computing. In the area of graphics, for
// example, a bitmap can represent an image by having a ‘1’ represent a black pixel and a ‘0’ represent a
// white pixel.
// Consider the following two ways of representing a rectangular bit map. In the first, it is simply
// represented as a two dimensional array of 1’s and 0’s. The second is based on a decomposition technique.
// First, the entire bit map is considered. If all bits within it are 1, a ‘1’ is output. If all bits within it are
// 0, a ‘0’ is output. Otherwise, a D is output, the bit map is divided into quarters (as described below),
// and each of those is processed in the same way as the original bit map. The quarters are processed in
// top left, top right, bottom left, bottom right order. Where a bit map being divided has an even number
// of rows and an even number of columns, all quarters have the same dimensions. Where the number of
// columns is odd, the left quarters have one more column than the right. Where the number of rows is
// odd the top quarters have one more row than the bottom. Note that if a region having only one row
// or one column is divided then two halves result, with the top half processed before the bottom where
// a single column is divided, and the left half before the right if a single row is divided.
// Write a program that will read in bitmaps of either form and transform them to the other form.

// Input
// Input will consist of a series of bit maps. Each bit map begins with a line giving its format (‘B’ or ‘D’)
// and its dimensions (rows and columns). Neither dimension will be greater than 200. There will be at
// least one space between each of the items of information. Following this line will be one or more lines
// containing the sequence of ‘1’, ‘0’ and ‘D’ characters that represent the bit map, with no intervening
// spaces. Each line (except the last, which may be shorter) will contain 50 characters. A ‘B’ type bitmap
// will be written left to right, top to bottom.
// The file will be terminated by a line consisting of a single ‘#’.

// Output
// Output will consist of a series of bitmaps, one for each bit map of the input. Output of each bit map
// begins on a new line and will be in the same format as the input. The width and height are to be
// output right justified in fields of width four.

// Sample Input
// B 3 4
// 001000011011
// D 2 3
// DD10111
// #

// Sample Output
// D 3 4
// D0D1001D101
// B 2 3
// 101111

#include <iostream>
#include <iomanip>
#include <string>
#include <vector>
#include <algorithm>

void BtoD(const std::vector<std::vector<int>> & bm, int Rs, int Re, int Cs, int Ce, int &p) {
    // Ignore 0-by-0 bitmap.
    if (Rs == Re || Cs == Ce) {
        return;
    }

    int zeros = 0;

    for (int r = Rs; r < Re; ++r) {
        zeros += static_cast<int>(count(bm[r].begin() + Cs, bm[r].begin() + Ce, 0));
    }

    // "Each line will contain 50 characters."
    if (p > 0 && p % 50 == 0) {
        std::cout << std::endl;
    }

    ++p;

    if (zeros == (Re - Rs) * (Ce - Cs)) {
        std::cout << "0";
    } else if (zeros == 0) {
        std::cout << "1";
    } else {
        std::cout << "D";
        int rHalf = (Rs + Re + 1) / 2;
        int cHalf = (Cs + Ce + 1) / 2;
        BtoD(bm, Rs, rHalf, Cs, cHalf, p);
        BtoD(bm, Rs, rHalf, cHalf, Ce, p);
        BtoD(bm, rHalf, Re, Cs, cHalf, p);
        BtoD(bm, rHalf, Re, cHalf, Ce, p);
    }
}

void DtoB(std::vector<std::vector<int>> & bm, int Rs, int Re, int Cs, int Ce) {
    if (Rs == Re || Cs == Ce) {
        return;
    }

    // Get 1 char from cin.
    int ch = std::cin.get();

    if (ch == '0' || ch == '1')
    {
        for (int r = Rs; r < Re; ++r) {
            for (int c = Cs; c < Ce; ++c) {
                bm[r][c] = ch - '0';
            }
        }

        return;
    } else {
        int rHalf = (Rs + Re + 1) / 2;
        int cHalf = (Cs + Ce + 1) / 2;
        DtoB(bm, Rs, rHalf, Cs, cHalf);
        DtoB(bm, Rs, rHalf, cHalf, Ce);
        DtoB(bm, rHalf, Re, Cs, cHalf);
        DtoB(bm, rHalf, Re, cHalf, Ce);
    }
}

int main() {  
    char type;

    while (std::cin >> type, type != '#') {
        int row, col;
        std::cin >> row >> col;
        std::cin.ignore();

        std::cout << (type == 'B'? "D" : "B") 
                  << std::right << std::setw(4) << row  
                  << std::right << std::setw(4) << col 
                  << std::endl;

        std::vector<std::vector<int> > bm(row, std::vector<int>(col));       

        if (type == 'B') {
            std::string s;

            while (int(s.size()) < row * col) {
                std::string tmp;
                std::getline(std::cin, tmp);
                s = s + tmp;
            }

            for (int r = 0; r < row; ++r) {
                for (int c = 0; c < col; ++c) {
                    bm[r][c] = s[r * col + c] - '0';
                }
            }
            // p indicates how many characters have been printed.
            int p = 0;
            BtoD(bm, 0, row, 0, col, p);
            std::cout << std::endl;
        } else {
            DtoB(bm, 0, row, 0, col);

            for (int r = 0; r < row; ++r) {
                for (int c = 0; c < col; ++c)
                {
                    // "Each line will contain 50 characters."
                    if (r + c > 0 && (r * col + c) % 50 == 0) {
                        std::cout << std::endl;
                    }

                    std::cout << bm[r][c];
                }
            }

            std::cout << std::endl;
        }
    }

    return 0;
}
