// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=117&page=show_problem&problem=2847

// UVA 11747 - Heavy Cycle Edges
// Given an undirected graph with edge
// weights, a minimum spanning tree is a
// subset of edges of minimum total weight
// such that any two nodes are connected
// by some path containing only these edges.
// A popular algorithm for finding the minimum spanning tree T in a graph proceeds
// as follows:
// • let T be initially empty
// • consider the edges e1, . . . , em in increasing order of weight
// – add ei to T if the endpoints of
// ei are not connected by a path
// in T
// An alternative algorithm is the following:
// • let T be initially the set of all edges
// • while there is some cycle C in T
// – remove edge e from T where e has the heaviest weight in C
// Your task is to implement a function related to this algorithm. Given an undirected graph G with
// edge weights, your task is to output all edges that are the heaviest edge in some cycle of G.

// Input
// The first input of each case begins with integers n and m with 1 ≤ n ≤ 1, 000 and 0 ≤ m ≤ 25, 000
// where n is the number of nodes and m is the number of edges in the graph. Following this are m
// lines containing three integers u, v, and w describing a weight w edge connecting nodes u and v where
// 0 ≤ u, v < n and 0 ≤ w < 2
// 31. Input is terminated with a line containing n = m = 0; this case should
// not be processed. You may assume no two edges have the same weight and no two nodes are directly
// connected by more than one edge

// Output
// Output for an input case consists of a single line containing the weights of all edges that are the heaviest
// edge in some cycle of the input graph. These weights should appear in increasing order and consecutive
// weights should be separated by a space. If there are no cycles in the graph then output the text ‘forest’
// instead of numbers.

// Sample Input
// 3 3
// 0 1 1
// 1 2 2
// 2 0 3
// 4 5
// 0 1 1
// 1 2 2
// 2 3 3
// 3 1 4
// 0 2 0
// 3 1
// 0 1 1
// 0 0

// Sample Output
// 3
// 2 4
// forest

// sum the edge weights of the chords

#include <algorithm>
#include <iostream>
#include <limits>
#include <queue>
#include <vector>

class UnionFind {
private:
    std::vector<unsigned> p, rank;
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

    unsigned findSet(unsigned i) {
        return (this->p[i] == i) ? i : (this->p[i] = this->findSet(this->p[i]));
    }

    bool isSameSet(unsigned i, unsigned j) {
        return this->findSet(i) == this->findSet(j);
    }

    void unionSet(unsigned i, unsigned j) {
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

std::vector<std::pair<unsigned, std::pair<unsigned, unsigned>>> kruskal(unsigned n, std::vector<std::pair<unsigned, std::pair<unsigned, unsigned>>> & edges) {
    std::vector<std::pair<unsigned, std::pair<unsigned, unsigned>>> mst(n - 1);
    unsigned i = 0;
    // all vertices are disjoint sets initially
    UnionFind uf(n);

    // sort by edge weight O(E log E)
    std::sort(edges.begin(), edges.end());

    // for each edge, O(E)
    for (const std::pair<unsigned, std::pair<unsigned, unsigned>> & edge : edges) {
        // check
        if (!uf.isSameSet(edge.second.first, edge.second.second)) {
            // link them
            uf.unionSet(edge.second.first, edge.second.second);
            // add edge to MST
            mst[i++] = edge;
        }
    }

    return mst;
}

int main() {
    unsigned n, m, i;
    std::vector<std::pair<unsigned, std::pair<unsigned, unsigned>>> edges, mst, ans;

    while (std::cin >> n >> m && (n || m)) {
        edges.resize(m);
        ans.clear();

        for (std::pair<unsigned, std::pair<unsigned, unsigned>> & e : edges) {
            std::cin >> e.second.first >> e.second.second >> e.first;
        }

        mst = kruskal(n, edges);

        for (const std::pair<unsigned, std::pair<unsigned, unsigned>> & e : edges) {
            if (std::find(mst.begin(), mst.end(), e) == mst.end()) {
                ans.push_back(e);
            }
        }

        if (ans.empty()) {
            std::cout << "forest" << std::endl;
        } else {
            for (i = 0; i + 1 < ans.size(); i++) {
                std::cout << ans[i].first << ' ';
            }
            
            std::cout << ans.back().first << std::endl;
        }
    }
	
	return 0;
}
