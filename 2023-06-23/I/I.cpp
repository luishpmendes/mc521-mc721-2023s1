// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=246&page=show_problem&problem=3601

// UVA 1160 - X-Plosives
// A secret service developed a new kind of explosive that attain its volatile property only when a specific
// association of products occurs. Each product is a mix of two different simple compounds, to which we
// call a binding pair. If N > 2, then mixing N different binding pairs containing N simple compounds
// creates a powerful explosive. For example, the binding pairs A+B, B+C, A+C (three pairs, three
// compounds) result in an explosive, while A+B, B+C, A+D (three pairs, four compounds) does not.
// You are not a secret agent but only a guy in a delivery agency with one dangerous problem: receive
// binding pairs in sequential order and place them in a cargo ship. However, you must avoid placing in
// the same room an explosive association. So, after placing a set of pairs, if you receive one pair that
// might produce an explosion with some of the pairs already in stock, you must refuse it, otherwise, you
// must accept it.
// An example. Lets assume you receive the following sequence: A+B, G+B, D+F, A+E, E+G,
// F+H. You would accept the first four pairs but then refuse E+G since it would be possible to make the
// following explosive with the previous pairs: A+B, G+B, A+E, E+G (4 pairs with 4 simple compounds).
// Finally, you would accept the last pair, F+H.
// Compute the number of refusals given a sequence of binding pairs.

// Input
// The input will contain several test cases, each of them as described below. Consecutive
// test cases are separated by a single blank line.
// Instead of letters we will use integers to represent compounds. The input contains several lines.
// Each line (except the last) consists of two integers (each integer lies between 0 and 105
// ) separated by
// a single space, representing a binding pair.
// Each test case ends in a line with the number ‘-1’. You may assume that no repeated binding pairs
// appears in the input.

// Output
// For each test case, the output must follow the description below.
// A single line with the number of refusals.

// Sample Input
// 1 2
// 3 4
// 3 5
// 3 1
// 2 3
// 4 1
// 2 6
// 6 5
// -1

// Sample Output
// 3

// count the number of edges not taken by Kruskal’s

#include <algorithm>
#include <iostream>
#include <limits>
#include <queue>
#include <vector>

#define MAX_N 100100

class UnionFind {
private:
    std::vector<int> p, rank;
public:
    UnionFind(unsigned n) {
        unsigned i;
        this->rank.assign(n, 0);
        this->p.assign(n, 0);
        for (i = 0; i < n; i++) {
            this->p[i] = i;
        }
    }

    UnionFind() : UnionFind(0) {}

    unsigned findSet(int i) {
        return (this->p[i] == i) ? i : (this->p[i] = this->findSet(this->p[i]));
    }

    bool isSameSet(int i, int j) {
        return this->findSet(i) == this->findSet(j);
    }

    void unionSet(int i, int j) {
        // if from different set
        if (!this->isSameSet(i, j)) {
            unsigned x = findSet(i),
                     y = findSet(j);
            
            if (this->rank[x] > this->rank[y]) {
                // rank keeps the tree short
                this->p[y] = x;
            } else {
                this->p[x] = y;
                
                if (this->rank[x] == this->rank[y]) {
                    this->rank[y]++;
                }
            }
        }
    }
};

int main() {
	int x, y, ans;
    UnionFind uf;

    while (std::cin >> x >> y) {
        uf = UnionFind(MAX_N);
        ans = 0;
        uf.unionSet(x, y);

        while (std::cin >> x && x != -1) {
            std::cin >> y;

            if (uf.isSameSet(x, y)) {
                ans++;
            } else {
                uf.unionSet(x, y);
            }
        }

        std::cout << ans << std::endl;
    }

	return 0;
}
