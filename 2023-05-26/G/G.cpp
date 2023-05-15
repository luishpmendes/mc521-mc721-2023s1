// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=25&page=show_problem&problem=2391

// UVA 11396 - Claw Decomposition

// A claw is defined as a pointed curved nail on the end of each toe in birds, some
// reptiles, and some mammals. However, if you are a graph theory enthusiast, you
// may understand the following special class of graph as shown in the following figure
// by the word claw.
// If you are more concerned about graph theory terminology, you may want to
// define claw as K1,3.
// Lets leave the definition for the moment and come to the problem. You are
// given a simple undirected graph in which every vertex has degree 3. You are to
// figure out whether the graph can be decomposed into claws or not.
// Just for the sake of clarity, a decomposition of a graph is a list of subgraphs such that each edge
// appears in exactly one subgraph in the list.

// Input

// There will be several cases in the input file. Each case starts with the number of vertices in the graph,
// V (4 ≤ V ≤ 300). This is followed by a list of edges. Every line in the list has two integers, a and b,
// the endpoints of an edge (1 ≤ a, b ≤ V ). The edge list ends with a line with a pair of ‘0’. The end of
// input is denoted by a case with V = 0. This case should not be processed.

// Output

// For every case in the input, print ‘YES’ if the graph can be decomposed into claws and ‘NO’ otherwise.

// Sample Input

// 4
// 1 2
// 1 3
// 1 4
// 2 3
// 2 4
// 3 4
// 0 0
// 6
// 1 2
// 1 3
// 1 6
// 2 3
// 2 5
// 3 4
// 4 5
// 4 6
// 5 6
// 0 0
// 0

// Sample Output

// NO
// NO

// it is just a bipartite graph check

#include <iostream>
#include <queue>
#include <vector>

int main() {
    unsigned n, a, b, i, u, v;
    std::vector<std::vector<unsigned>> adj_list;
    std::queue<unsigned> q;
    std::vector<unsigned> color; // color of each vertex
    bool is_bipartite;

    while (std::cin >> n && n > 0) {
        adj_list.resize(n);
        adj_list.assign(n, std::vector<unsigned>());
        q = std::queue<unsigned>();
        color.resize(n);
        color.assign(n, 3);

        while (std::cin >> a >> b && a > 0 && b > 0) {
            adj_list[a - 1].push_back(b - 1);
            adj_list[b - 1].push_back(a - 1);
        }

        // addition of one boolean flag, initially true
        is_bipartite = true;
        // start from source
        q.push(0);
        color[0] = 0;

        // similar to the original BFS routine
        while (!q.empty() && is_bipartite) {
            // queue: layer by layer!
            u = q.front();
            q.pop();

            for (i = 0; i < adj_list[u].size() && is_bipartite; i++) {
                // for each neighbor of u
                v = adj_list[u][i];

                if (color[v] == 3) { // if v.first is unvisited + reachable
                    // we color with another color
                    color[v] = 1 - color[u];
                    // enqueue v for the next iteration
                    q.push(v);
                } else if (color[v] == color[u]) { // u ans v have the same color
                    // we have a coloring conflict
                    is_bipartite = false;
                }
            }
        }

        if (is_bipartite) {
            std::cout << "YES" << std::endl;
        } else {
            std::cout << "NO" << std::endl;
        }
    }

    return 0;
}
