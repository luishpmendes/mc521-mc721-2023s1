// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=17&page=show_problem&problem=1448

// UVA 10507 - Waking up brain

// Most recent researches on Neuroanatomy have permitted us of identifying a series of large electric signal
// transmision route connecting different areas of the brain. Even more, it has been found that if one slept
// area X of the brain is directly connected to at least three awake areas for a year, the X area will wake
// up. There is evidence of when an area X of the brain wakes up, it remains awake. Let A, B, C, …the
// different areas of the brain and let’s imagine a brain with some initially slept areas, interconnected one
// another. If three of these areas wake up by direct stimulation according to the previous researches, how
// many years will all the slept areas take to wake up?

// Input

// The input file contains several test cases, each of them as described below. There is a blank line between
// two consecutive inputs.
//  • The first line of the input is an integer, 3 ≤ N ≤ 26, that indicates the number of slept areas.
//  • The second line of the input is an integer M ≥ 0, that indicates the number of connections (if A
// is connected to B, then B is connected to A, but it counts only once).
//  • The third line consists of three characters that indicate which areas have wake up by direct
// stimulation.
//  • The remaining M lines consist of two characters each one, that indicate the different conections
// between areas of the brain, one line per conection.

// Output

// The output is only one line with one of the following text messages:
// ‘THIS BRAIN NEVER WAKES UP’
// ‘WAKE UP IN, n, YEARS’, where n is the number of the years all the brain has taken to wake up.

// Sample Input

// 6
// 11
// HAB
// AB
// AC
// AH
// BD
// BC
// BF
// CD
// CF
// CH
// DF
// FH

// Sample Output

// WAKE UP IN, 3, YEARS

// disjoint sets simplifies this problem

#include <iostream>
#include <map>
#include <string>
#include <vector>

class UnionFind {
    private:
        std::vector<int> p, rank;

    public:
        UnionFind(int N) {
            this->rank.assign(N, 0);
            this->p.assign(N, 0);
            
            for (int i = 0; i < N; i++) {
                this->p[i] = i;
            }
        }

        UnionFind() : UnionFind(0) {}

        int findSet(int i) {
            return (this->p[i] == i) ? i : (this->p[i] = this->findSet(this->p[i]));
        }

        bool isSameSet(int i, int j) {
            return this->findSet(i) == this->findSet(j);
        }

        void unionSet(int i, int j) {
            // if from different set
            if (!this->isSameSet(i, j)) {
                int x = this->findSet(i),
                    y = this->findSet(j);

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

int main () {
    unsigned n, m, i, j, years, prev, curr, awakeNeighbors;
    std::string awake, connection;
    std::map<char, unsigned> index;
    std::vector<std::vector<unsigned>> adj;
    std::vector<unsigned> neighborsOfAwake;
    UnionFind uf;

    while (std::cin >> n >> m >> awake) {
        index.clear();
        adj.assign(26, std::vector<unsigned>());
        uf = UnionFind(26);

        index[awake[0]] = 0;
        index[awake[1]] = 1;
        index[awake[2]] = 2;
        uf.unionSet(0, 1);
        uf.unionSet(1, 2);
        i = 3;

        while (m--) {
            std::cin >> connection;

            if (index.find(connection[0]) == index.end()) {
                index[connection[0]] = i++;
            }

            if (index.find(connection[1]) == index.end()) {
                index[connection[1]] = i++;
            }

            adj[index[connection[0]]].push_back(index[connection[1]]);
            adj[index[connection[1]]].push_back(index[connection[0]]);
        }

        years = 0;
        prev = 1;
        curr = 3;

        while (prev < curr) {
            years++;
            prev = curr;
            neighborsOfAwake.clear();

            for (i = 3; i < n; i++) {
                if (uf.isSameSet(0, i)) {
                    continue;
                }

                awakeNeighbors = 0;
                for (j = 0; j < adj[i].size(); j++) {
                    if (uf.isSameSet(0, adj[i][j])) {
                        awakeNeighbors++;
                    }
                }

                if (awakeNeighbors >= 3) {
                    neighborsOfAwake.push_back(i);
                }
            }

            for (i = 0; i < neighborsOfAwake.size(); i++) {
                uf.unionSet(0, neighborsOfAwake[i]);
            }

            curr = 0;

            for (i = 0; i < n; i++) {
                if (uf.isSameSet(0, i)) {
                    curr++;
                }
            }
        }

        if (curr == n) {
            std::cout << "WAKE UP IN, " << years-1 << ", YEARS" << std::endl;
        } else {
            std::cout << "THIS BRAIN NEVER WAKES UP" << std::endl;
        }
    }

    return 0;
}
