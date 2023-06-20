// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=12&page=show_problem&problem=989

// UVA 10048 - Audiophobia
// Consider yourself lucky! Consider yourself lucky to be still breathing and having fun participating in
// this contest. But we apprehend that many of your descendants may not have this luxury. For, as you
// know, we are the dwellers of one of the most polluted cities on earth. Pollution is everywhere, both in
// the environment and in society and our lack of consciousness is simply aggravating the situation.
// However, for the time being, we will consider only one type of pollution - the sound pollution. The
// loudness or intensity level of sound is usually measured in decibels and sound having intensity level 130
// decibels or higher is considered painful. The intensity level of normal conversation is 60 65 decibels and
// that of heavy traffic is 70 80 decibels.
// Consider the following city map where the edges refer to streets and the nodes refer to crossings.
// The integer on each edge is the average intensity level of sound (in decibels) in the corresponding street.
// To get from crossing A to crossing G you may follow the following path: A- C- F- G. In that case
// you must be capable of tolerating sound intensity as high as 140 decibels. For the paths A- B- E- G,
// A- B- D- G and A- C- F- D- G you must tolerate respectively 90, 120 and 80 decibels of sound intensity.
// There are other paths, too. However, it is clear that A- C- F- D- G is the most comfortable path since
// it does not demand you to tolerate more than 80 decibels.
// In this problem, given a city map you are required to determine the minimum sound intensity level
// you must be able to tolerate in order to get from a given crossing to another.

// Input
// The input may contain multiple test cases.
// The first line of each test case contains three integers C(≤ 100), S(≤ 1000) and Q(≤ 10000) where
// C indicates the number of crossings (crossings are numbered using distinct integers ranging from 1 to
// C), S represents the number of streets and Q is the number of queries.
// Each of the next S lines contains three integers: c1, c2 and d indicating that the average sound
// intensity level on the street connecting the crossings c1 and c2 (c1 ̸= c2) is d decibels.
// Each of the next Q lines contains two integers c1 and c2 (c1 ̸= c2) asking for the minimum sound
// intensity level you must be able to tolerate in order to get from crossing c1 to crossing c2.
// The input will terminate with three zeros form C, S and Q.

// Output
// For each test case in the input first output the test case number (starting from 1) as shown in the
// sample output. Then for each query in the input print a line giving the minimum sound intensity level
// (in decibels) you must be able to tolerate in order to get from the first to the second crossing in the
// query. If there exists no path between them just print the line “no path”.
// Print a blank line between two consecutive test cases.

// Sample Input
// 7 9 3
// 1 2 50
// 1 3 60
// 2 4 120
// 2 5 90
// 3 6 50
// 4 6 80
// 4 7 70
// 5 7 40
// 6 7 140
// 1 7
// 2 6
// 6 2
// 7 6 3
// 1 2 50
// 1 3 60
// 2 4 120
// 3 6 50
// 4 6 80
// 5 7 40
// 7 5
// 1 7
// 2 4
// 0 0 0

// Sample Output
// Case #1
// 80
// 60
// 60
// 
// Case #2
// 40
// no path
// 80

// minimax

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

std::vector<std::vector<std::pair<unsigned, unsigned>>> kruskal(unsigned n, std::vector<std::pair<unsigned, std::pair<unsigned, unsigned>>> & edges) {
    std::vector<std::vector<std::pair<unsigned, unsigned>>> adj(n, std::vector<std::pair<unsigned, unsigned>>());
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
            adj[edge.second.first].push_back(std::make_pair(edge.second.second, edge.first));
            adj[edge.second.second].push_back(std::make_pair(edge.second.first, edge.first));
        }
    }

    return adj;
}

std::vector<unsigned> dijkstra(const std::vector<std::vector<std::pair<unsigned, unsigned>>> & adj, unsigned s) {
    std::vector<unsigned> dist(adj.size(), std::numeric_limits<unsigned>::max() >> 2);
    std::priority_queue<std::pair<unsigned, unsigned>, std::vector<std::pair<unsigned, unsigned>>, std::greater<std::pair<unsigned, unsigned>>> pq;
    std::pair<unsigned, unsigned> front, v;
    unsigned d, u, j;

    dist[s] = 0;
    pq.push(std::make_pair(dist[s], s));

    // main loop
    while (!pq.empty()) {
        // greedy: get shortest unvisited vertex
        front = pq.top();
        pq.pop();
        d = front.first;
        u = front.second;

        // this is a very important check
        if (d <= dist[u]) {
            for (j = 0; j < adj[u].size(); j++) {
                // all outgoing edges from u
                v = adj[u][j];

                if (dist[v.first] > std::max(dist[u], v.second)) {
                    // relax operation
                    dist[v.first] = std::max(dist[u], v.second);
                    pq.push(std::make_pair(dist[v.first], v.first));
                }
            }
        }
    }

    return dist;
}

int main() {
    unsigned c, s, q, c1, c2, d, t;
    std::vector<std::pair<unsigned, std::pair<unsigned, unsigned>>> edges;
    std::vector<std::vector<std::pair<unsigned, unsigned>>> adj;
    std::vector<unsigned> dist;

    t = 0;

    while (std::cin >> c >> s >> q && (c || s || q)) {
        edges.clear();

        while (s--) {
            std::cin >> c1 >> c2 >> d;
            c1--;
            c2--;
            edges.push_back(std::make_pair(d, std::make_pair(c1, c2)));
        }

        adj = kruskal(c, edges);

        if (t) {
            std::cout << std::endl;
        }

        std::cout << "Case #" << ++t << std::endl;

        while (q--) {
            std::cin >> c1 >> c2;
            c1--;
            c2--;
            dist = dijkstra(adj, c1);

            if (dist[c2] >= std::numeric_limits<unsigned>::max() >> 2) {
                std::cout << "no path" << std::endl;
            } else {
                std::cout << dist[c2] << std::endl;
            }
        }
    }
	
	return 0;
}
