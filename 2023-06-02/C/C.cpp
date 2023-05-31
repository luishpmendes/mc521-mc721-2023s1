// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=25&page=show_problem&problem=2352

// UVA 11367 - Full Tank?
// After going through the receipts from your car trip through Europe this summer,
// you realised that the gas prices varied between the cities you visited. Maybe you
// could have saved some money if you were a bit more clever about where you filled
// your fuel?
// To help other tourists (and save money yourself next time), you want to write
// a program for finding the cheapest way to travel between cities, filling your tank
// on the way. We assume that all cars use one unit of fuel per unit of distance, and
// start with an empty gas tank.

// Input
// The first line of input gives 1 ≤ n ≤ 1000 and 0 ≤ m ≤ 10000, the number of cities and roads. Then
// follows a line with n integers 1 ≤ pi ≤ 100, where pi
// is the fuel price in the ith city. Then follow m lines
// with three integers 0 ≤ u, v < n and 1 ≤ d ≤ 100, telling that there is a road between u and v with
// length d. Then comes a line with the number 1 ≤ q ≤ 100, giving the number of queries, and q lines
// with three integers 1 ≤ c ≤ 100, s and e, where c is the fuel capacity of the vehicle, s is the starting
// city, and e is the goal.

// Output
// For each query, output the price of the cheapest trip from s to e using a car with the given capacity, or
// ‘impossible’ if there is no way of getting from s to e with the given car

// Sample Input
// 5 5
// 10 10 20 12 13
// 0 1 9
// 0 2 8
// 1 2 1
// 1 3 11
// 2 3 7
// 2
// 10 0 3
// 20 1 4

// Sample Output
// 170
// impossible

#include <iostream>
#include <limits>
#include <queue>
#include <set>
#include <utility>
#include <tuple>
#include <vector>

unsigned dijkstra(const std::vector<unsigned> & p,
                  const std::vector<std::vector<std::pair<unsigned, unsigned>>> & adj,
                  unsigned c, unsigned s, unsigned e) {
    std::priority_queue<std::tuple<unsigned, unsigned, unsigned>,
                        std::vector<std::tuple<unsigned, unsigned, unsigned>>,
                        std::greater<std::tuple<unsigned, unsigned, unsigned>>> pq;
    std::vector<std::vector<unsigned>> memo(adj.size(),
            std::vector<unsigned>(c + 1, std::numeric_limits<unsigned>::max()));
    unsigned result, v, d, f;

    result = std::numeric_limits<unsigned>::max();
    pq.push(std::make_tuple(0, 0, s));

    while (!pq.empty() && result == std::numeric_limits<unsigned>::max()) {
        v = std::get<2>(pq.top());
        f = std::get<1>(pq.top());
        d = std::get<0>(pq.top());
        pq.pop();

        if (v == e) {
            result = d;
        } else if (memo[v][f] == std::numeric_limits<unsigned>::max()) {
            memo[v][f] = d;

            if (f < c) {
                pq.push(std::make_tuple(d + p[v], f + 1, v));
            }

            for (const auto & u : adj[v]) {
                if (f >= u.second) {
                    pq.push(std::make_tuple(d, f - u.second, u.first));
                }
            }
        }
    }

    return result;
}

int main() {
    unsigned n, m, u, v, d, q, c, s, e, ans, i;
    std::vector<unsigned> p;
    std::vector<std::vector<std::pair<unsigned, unsigned>>> adj;

    while (std::cin >> n >> m) {
        p.resize(n);
        adj.resize(n);
        adj.assign(n, std::vector<std::pair<unsigned, unsigned>>());

        for (i = 0; i < n; i++) {
            std::cin >> p[i];
        }

        for (i = 0; i < m; i++) {
            std::cin >> u >> v >> d;
            adj[u].push_back(std::make_pair(v, d));
            adj[v].push_back(std::make_pair(u, d));
        }

        std::cin >> q;

        while (q--) {
            std::cin >> c >> s >> e;
            ans = dijkstra(p, adj, c, s, e);

            if (ans != std::numeric_limits<unsigned>::max()) {
                std::cout << ans << std::endl;
            } else {
                std::cout << "impossible" << std::endl;
            }
        }
    }
	
	return 0;
}
