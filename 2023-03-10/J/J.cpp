// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=8&page=show_problem&problem=549

// UVA 608 - Counterfeit Dollar

// Sally Jones has a dozen Voyageur silver dollars. However, only eleven of the coins are true silver
// dollars; one coin is counterfeit even though its color and size make it indistinguishable from the real
// silver dollars. The counterfeit coin has a different weight from the other coins but Sally does not know
// if it is heavier or lighter than the real coins.
// Happily, Sally has a friend who loans her a very accurate balance scale. The friend will permit Sally
// three weighings to find the counterfeit coin. For instance, if Sally weighs two coins against each other
// and the scales balance then she knows these two coins are true. Now if Sally weighs one of the true
// coins against a third coin and the scales do not balance then Sally knows the third coin is counterfeit
// and she can tell whether it is light or heavy depending on whether the balance on which it is placed
// goes up or down, respectively.
// By choosing her weighings carefully, Sally is able to ensure that she will find the counterfeit coin
// with exactly three weighings.

// Input

// The first line of input is an integer n (n > 0) specifying the number of cases to follow. Each case
// consists of three lines of input, one for each weighing. Sally has identified each of the coins with the
// letters A–L. Information on a weighing will be given by two strings of letters and then one of the words
// “up”, “down”, or “even”. The first string of letters will represent the coins on the left balance; the
// second string, the coins on the right balance. (Sally will always place the same number of coins on the
// right balance as on the left balance.) The word in the third position will tell whether the right side of
// the balance goes up, down, or remains even.

// Output

// For each case, the output will identify the counterfeit coin by its letter and tell whether it is heavy or
// light. The solution will always be uniquely determined.

// Sample Input

// 1
// ABCD EFGH even
// ABCI EFJK up
// ABIJ EFGH even

// Sample Output

// K is the counterfeit coin and it is light.

// classical problem

#include <iostream>
#include <string>
#include <vector>

bool is_heavy(const std::vector<std::string> & x, const std::vector<std::string> & y, const std::vector<std::string> & z, char c) {
    unsigned i, j;
    int l, r;

	for (i = 0; i < 3; i++) {
		l = r = 0;

		for (j = 0; j < x[i].size(); j++) {
			if (x[i][j] == c) {
                l += 100;
            } else {
                l++;
            }

			if (y[i][j] == c) {
                r += 100;
            } else {
                r++;
            }
		}

        if ((z[i][0] == 'e' && l != r) || (z[i][0] == 'u' && l <= r) || (z[i][0] == 'd' && l >= r)) {
            return false;
        }
	}

	return true;
}

bool is_light (const std::vector<std::string> & x, const std::vector<std::string> & y, const std::vector<std::string> & z, char c) {
    unsigned i, j;
    int l, r;

	for (i = 0; i < 3; i++) {
		l = r = 0;

		for (j = 0; j < x[i].size(); j++) {
			if (x[i][j] == c) {
                l -= 100;
            } else {
                l++;
            }

			if (y[i][j] == c) {
                r -= 100;
            } else {
                r++;
            }
		}

        if ((z[i][0] == 'e' && l != r) || (z[i][0] == 'u' && l <= r) || (z[i][0] == 'd' && l >= r)) {
            return false;
        }
	}

	return true;
}

int main() {
    unsigned n, i;
    char c;
    std::vector<std::string> x(3), y(3), z(3);
    bool flag;

	while (std::cin >> n) {
        while (n--) {
            for (i = 0; i < 3; i++) {
                std::cin >> x[i] >> y[i] >> z[i];
            }

            for (c = 'A', flag = false; c != 'M' && !flag; c++) {
                if (is_heavy(x, y, z, c)) {
                    std::cout << c << " is the counterfeit coin and it is heavy." << std::endl;
                    flag = true;
                } else if (is_light(x, y, z, c)) {
                    std::cout << c << " is the counterfeit coin and it is light." << std::endl;
                    flag = true;
                }
            }
        }
    }

    return 0;
}
