// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=246&page=show_problem&problem=3553

// UVA 1112 - Mice and Maze
// A set of laboratory mice is being trained to escape a maze. The maze is made up of cells, and each cell
// is connected to some other cells. However, there are obstacles in the passage between cells and therefore
// there is a time penalty to overcome the passage Also, some passages allow mice to go one-way, but not
// the other way round.
// Suppose that all mice are now trained and, when placed in an arbitrary cell in the maze, take a
// path that leads them to the exit cell in minimum time.
// We are going to conduct the following experiment: a mouse is placed in each cell of the maze and
// a count-down timer is started. When the timer stops we count the number of mice out of the maze.
// Write a program that, given a description of the maze and the time limit, predicts the number of
// mice that will exit the maze. Assume that there are no bottlenecks is the maze, i.e. that all cells have
// room for an arbitrary number of mice.

// Input
// The input begins with a single positive integer on a line by itself indicating the number of the cases
// following, each of them as described below. This line is followed by a blank line, and there is also a
// blank line between two consecutive inputs.
// The maze cells are numbered 1, 2, . . . , N, where N is the total number of cells. You can assume
// that N ≤ 100.
// The first three input lines contain N, the number of cells in the maze, E, the number of the exit
// cell, and the starting value T for the count-down timer (in some arbitrary time unit).
// The fourth line contains the number M of connections in the maze, and is followed by M lines, each
// specifying a connection with three integer numbers: two cell numbers a and b (in the range 1, . . . , N)
// and the number of time units it takes to travel from a to b.
// Notice that each connection is one-way, i.e., the mice can’t travel from b to a unless there is another
// line specifying that passage. Notice also that the time required to travel in each direction might be
// different.

// Output
// For each test case, the output must follow the description below. The outputs of two consecutive cases
// will be separated by a blank line.
// The output consists of a single line with the number of mice that reached the exit cell E in at most
// T time units.

// Sample Input
// 1
// 
// 4
// 2
// 1
// 8
// 1 2 1
// 1 3 1
// 2 1 1
// 2 4 1
// 3 1 1
// 3 4 1
// 4 2 1
// 4 3 1

// Sample Output
// 3

// LA 2425, SouthwesternEurope01, run Dijkstra’s from destination

#include <iostream>
#include <limits>
#include <queue>
#include <utility>
#include <vector>

std::vector<unsigned> dijkstra(
    const std::vector<std::vector<std::pair<unsigned, unsigned>>> & adj,
    unsigned s) {
    std::priority_queue<std::pair<unsigned, unsigned>,
                        std::vector<std::pair<unsigned, unsigned>>,
                        std::greater<std::pair<unsigned, unsigned>>> pq;
    std::vector<unsigned> dist(adj.size(), std::numeric_limits<unsigned>::max());
    unsigned u, d;

    pq.push(std::make_pair(0, s));
    dist[s] = 0;

    while (!pq.empty()) {
        u = pq.top().second;
        d = pq.top().first;
        pq.pop();

        if (d <= dist[u]) {
            for (const std::pair<unsigned, unsigned> & v : adj[u]) {
                if (dist[v.first] > dist[u] + v.second) {
                    dist[v.first] = dist[u] + v.second;
                    pq.push(std::make_pair(dist[v.first], v.first));
                }
            }
        }
    }

    return dist;
}

int main() {
    unsigned c, n, t, e, m, a, b, w, ans;
    std::vector<std::vector<std::pair<unsigned, unsigned>>> adj;
    std::vector<unsigned> dist;

    while (std::cin >> c) {
        while (c--) {
            std::cin >> n >> e >> t >> m;

            adj.resize(n);
            adj.assign(n, std::vector<std::pair<unsigned, unsigned>>());

            while (m--) {
                std::cin >> a >> b >> w;
                adj[b - 1].push_back(std::make_pair(a - 1, w));
            }

            dist = dijkstra(adj, e - 1);
            ans = 0;

            for (const unsigned & d : dist) {
                if (d <= t) {
                    ans++;
                }
            }

            std::cout << ans << std::endl;

            if (c) {
                std::cout << std::endl;
            }
        }
    }
	
	return 0;
}
