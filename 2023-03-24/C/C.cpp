// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=441&page=show_problem&problem=3977

// UVA 12532 - Interval Product

// It’s normal to feel worried and tense the day before a programming contest. To relax, you went out for
// a drink with some friends in a nearby pub. To keep your mind sharp for the next day, you decided to
// play the following game. To start, your friends will give you a sequence of N integers X1, X2, . . . , XN.
// Then, there will be K rounds; at each round, your friends will issue a command, which can be:
//  • a change command, when your friends want to change one of the values in the sequence; or
//  • a product command, when your friends give you two values I, J and ask you if the product
// XI × XI+1 × . . . × XJ−1 × XJ is positive, negative or zero.
// Since you are at a pub, it was decided that the penalty for a wrong answer is to drink a pint of
// beer. You are worried this could affect you negatively at the next day’s contest, and you don’t want
// to check if Ballmer’s peak theory is correct. Fortunately, your friends gave you the right to use your
// notebook. Since you trust more your coding skills than your math, you decided to write a program to
// help you in the game.

// Input

// Each test case is described using several lines. The first line contains two integers N and K, indicating
// respectively the number of elements in the sequence and the number of rounds of the game (1 ≤
// N, K ≤ 10^5). The second line contains N integers Xi that represent the initial values of the sequence
// (−100 ≤ Xi ≤ 100 for i = 1, 2, . . . , N). Each of the next K lines describes a command and starts with
// an uppercase letter that is either ‘C’ or ‘P’. If the letter is ‘C’, the line describes a change command, and
// the letter is followed by two integers I and V indicating that XI must receive the value V (1 ≤ I ≤ N
// and −100 ≤ V ≤ 100). If the letter is ‘P’, the line describes a product command, and the letter
// is followed by two integers I and J indicating that the product from XI to XJ , inclusive must be
// calculated (1 ≤ I ≤ J ≤ N). Within each test case there is at least one product command.

// Output

// For each test case output a line with a string representing the result of all the product commands in
// the test case. The i-th character of the string represents the result of the i-th product command. If the
// result of the command is positive the character must be ‘+’ (plus); if the result is negative the character
// must be ‘-’ (minus); if the result is zero the character must be ‘0’ (zero).

// Sample Input

// 4 6
// -2 6 0 -1
// C 1 10
// P 1 4
// C 3 7
// P 2 2
// C 4 -5
// P 1 4
// 5 9
// 1 5 -2 4 3
// P 1 2
// P 1 5
// C 4 -5
// P 1 5
// P 4 5
// C 3 0
// P 1 5
// C 4 -5
// C 4 -5

// Sample Output

// 0+-
// +-+-0

// clever usage of Fenwick/Segment Tree

#include <iostream>
#include <string>
#include <vector>

// the segment tree is stored like a heap array
class SegmentTree {
    private:
        std::vector<char> st;
        unsigned n;

        // same as binary heap operations
        int left (int p) {
            return p << 1;
        } 

        int right (int p) {
            return (p << 1) + 1;
        }

        char combine (char i, char j) {
            if (i == '0' || j == '0') {
                return '0';
            }

            if (i == j) {
                return '+';
            }

            return '-';
        }

        // O(n)
        char build(unsigned p, unsigned L, unsigned R, const std::vector<int> & A) {
            // as L == R, either one is fine
            if (L == R) {
                // store the sign
                this->st[p] = (A[L] < 0 ? '-' : (A[L] > 0 ? '+' : '0'));
            } else {
                // recursively compute the values
                char left_sign = this->build(this->left(p) , L , (L + R) / 2, A);
                char right_sign = this->build(this->right(p), ((L + R) / 2) + 1, R, A);

                this->st[p] = this->combine(left_sign, right_sign);
            }

            return this->st[p];
        }

        // O(log n)
        char query(unsigned p, unsigned L, unsigned R, unsigned i, unsigned j) {
            if (i > R || j < L) {
                // current segment outside query range
                return -1;
            }

            if (L >= i && R <= j) {
                // inside query range
                return this->st[p];
            }

            // compute the min position in the left and right part of the interval
            char left_sign = this->query(this->left(p) , L , (L + R) / 2, i, j);
            char right_sign = this->query(this->right(p), ((L + R) / 2) + 1, R , i, j);

            if (left_sign == -1) {
                // if we try to access segment outside query
                return right_sign;
            }

            if (right_sign == -1) {
                // same as above
                return left_sign;
            }

            // as in build routine
            return this->combine(left_sign, right_sign);
        }

        // O(log n)
        char update(unsigned p, unsigned L, unsigned R, unsigned i, int v) {
            if (i < L || i > R) {
                return this->st[p];
            }

            if (L == i && R == i) {
                this->st[p] = (v < 0 ? '-' : (v > 0 ? '+' : '0'));
                return this->st[p];
            }

            char left_sign = this->update(this->left(p) , L , (L + R) / 2, i, v);
            char right_sign = this->update(this->right(p), (L + R) / 2 + 1, R , i, v);

            this->st[p] = this->combine(left_sign, right_sign);
            return this->st[p];
        }

    public:
        SegmentTree(const std::vector<int> & A) {
            // copy content for local usage
            this->n = A.size();
            // create large enough vector of zeroes
            this->st.assign(4 * this->n, 0);
            // recursive build
            this->build(1, 0, this->n - 1, A);
        }

        // overloading
        char query(unsigned i, unsigned j) {
            return this->query(1, 0, this->n - 1, i, j);
        }

        // overloading
        char update(unsigned i, unsigned v) {
            return this->update(1, 0, this->n - 1, i, v);
        }
};

int main () {
    unsigned n, k, i, j;
    int v;
    std::vector<int> a;
    char command;
    std::string ans;

    while (std::cin >> n >> k) {
        a.resize(n);

        for (int & x : a) {
            std::cin >> x;
        }

        SegmentTree st(a);
        ans = "";

        while (k--) {
            std::cin >> command;

            if (command == 'C') {
                std::cin >> i >> v;
                st.update(i - 1, v);
            } else { // command == 'P'
                std::cin >> i >> j;
                ans += st.query(i - 1, j - 1);
            }
        }

        std::cout << ans << std::endl;
    }

    return 0;
}
