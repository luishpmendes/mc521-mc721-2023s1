// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=243&page=show_problem&problem=3277

// UVA 12125 - March of the Penguins
// Somewhere near the south pole, a number of penguins are standing on a number of ice floes. Being social animals, the penguins
// would like to get together, all on the same floe. The penguins do
// not want to get wet, so they have use their limited jump distance
// to get together by jumping from piece to piece. However, temperatures have been high lately, and the floes are showing cracks,
// and they get damaged further by the force needed to jump to another floe. Fortunately the penguins are real experts on cracking
// ice floes, and know exactly how many times a penguin can jump
// off each floe before it disintegrates and disappears. Landing on
// an ice floe does not damage it. You have to help the penguins
// find all floes where they can meet.

// Input
// On the first line one positive number: the number of testcases, at
// most 100. After that per testcase:
// • One line with the integer N (1 ≤ N ≤ 100) and a floating-point number D (0 ≤ D ≤ 100000),
// denoting the number of ice pieces and the maximum distance a penguin can jump.
// • N lines, each line containing xi
// , yi
// , ni and mi
// , denoting for each ice piece its X and Y coordinate,
// the number of penguins on it and the maximum number of times a penguin can jump off this
// piece before it disappears (−10000 ≤ xi
// , yi ≤ 10000, 0 ≤ ni ≤ 10, 1 ≤ mi ≤ 200).

// Output
// Per testcase:
// • One line containing a space-separated list of 0-based indices of the pieces on which all penguins
// can meet. If no such piece exists, output a line with the single number ‘-1’.

// Sample Input
// 2
// 5 3.5
// 1 1 1 1
// 2 3 0 1
// 3 5 1 1
// 5 1 1 1
// 5 4 0 1
// 3 1.1
// -1 0 5 10
// 0 0 3 9
// 2 0 1 1

// Sample Output
// 1 2 4
// -1

// max flow modeling with vertex capacities; another interesting problem, similar level with UVa 11380

#include <iostream>
#include <limits>
#include <queue>
#include <string>
#include <utility>
#include <vector>

unsigned augment(std::vector<std::vector<unsigned>> & res,
                 const std::vector<unsigned> & p,
                 const unsigned & s,
                 unsigned v,
                 unsigned minEdge) {
    unsigned f;
    // traverse BFS spanning tree from s->t
    if (v == s) {
        // record minEdge in a global var f
        f = minEdge;
    } else if (p[v] != p.size()) {
        f = augment(res, p, s, p[v], std::min(minEdge, res[p[v]][v]));

        res[p[v]][v] -= f;
        res[v][p[v]] += f;
    }

    return f;
}

unsigned edmondsKarp(const std::vector<std::vector<unsigned>> & adj,
                     std::vector<std::vector<unsigned>> & res,
                     const unsigned & s,
                     const unsigned & t) {
    // mf stands for max_flow
    unsigned mf = 0, f, u;
    std::vector<bool> visited (adj.size(), false);
    std::vector<unsigned> p (adj.size(), adj.size());
    std::queue<unsigned> q;

    // O(VE^2) Edmonds Karp’s algorithm
    while(true) {
        f = 0;
        visited.assign(adj.size(), false);
        p.assign(adj.size(), adj.size());
        visited[s] = true;
        q.push(s);

        while (!q.empty()) {
            u = q.front();
            q.pop();

            if (u == t) {
                // immediately stop BFS if we already reach sink t
                while (!q.empty()) {
                    q.pop();
                }
            } else {
                for (const unsigned & v : adj[u]) {
                    if (res[u][v] > 0 && !visited[v]) {
                        visited[v] = true;
                        q.push(v);
                        p[v] = u;
                    }
                }
            }
        }

        if (!visited[t]) {
            // we did not reach sink t so there is no more augmenting path
            break;
        }

        // find the min edge weight ‘f’ in this path, if any
        f = augment(res, p, s, t, std::numeric_limits<unsigned>::max());

        if (f == 0) {
            // we cannot send any more flow (‘f’ = 0), terminate
            break;
        }

        // we can still send a flow, increase the max flow!
        mf += f;
    }

    // this is the max flow value
    return mf;
}

int main() {
    unsigned t, N, n, m, sum, source, sink, i, j;
    double D;
    std::vector<unsigned> x, y, ans;
    std::vector<std::vector<unsigned>> adj, adj_cpy, res, res_cpy;

    while (std::cin >> t) {
        while (t--) {
            std::cin >> N >> D;
            sum = 0;
            source = 0;
            sink = 2 * N + 1;
            
            x.resize(N + 1);
            y.resize(N + 1);
            ans.clear();
            adj.assign(2 * N + 2, std::vector<unsigned>());
            res.assign(2 * N + 2, std::vector<unsigned>(2 * N + 2, 0));

            for (i = 1; i <= N; i++) {
                std::cin >> x[i] >> y[i] >> n >> m;
                sum += n;
                adj[source].push_back(i);
                adj[i].push_back(source);
                res[source][i] = n;
                adj[i].push_back(N + i);
                adj[N + i].push_back(i);
                res[i][N + i] = m;
            }

            for (i = 1; i <= N; i++) {
                for (j = 1; j <= N; j++) {
                    if (i != j && (x[i] - x[j]) * (x[i] - x[j]) + (y[i] - y[j]) * (y[i] - y[j]) < D * D) {
                        adj[N + i].push_back(j);
                        adj[j].push_back(N + i);
                        res[N + i][j] = std::numeric_limits<unsigned>::max();
                    }
                }
            }

            adj_cpy = std::vector<std::vector<unsigned>>(adj);
            res_cpy = std::vector<std::vector<unsigned>>(res);

            for (i = 1; i <= N; i++) {
                adj = std::vector<std::vector<unsigned>>(adj_cpy);
                res = std::vector<std::vector<unsigned>>(res_cpy);

                adj[i].push_back(sink);
                adj[sink].push_back(i);
                res[i][sink] = std::numeric_limits<unsigned>::max();

                if (edmondsKarp(adj, res, source, sink) == sum) {
                    ans.push_back(i - 1);
                }

                res[i][sink] = 0;
            }

            if (ans.empty()) {
                std::cout << "-1" << std::endl;
            } else {
                for (i = 0; i + 1 < ans.size(); i++) {
                    std::cout << ans[i] << ' ';
                }

                std::cout << ans[i] << std::endl;
            }
        }
    }

	return 0;
}
