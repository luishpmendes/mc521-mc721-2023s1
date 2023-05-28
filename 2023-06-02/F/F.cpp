// https://codeforces.com/problemset/problem/1433/G
// You are a mayor of Berlyatov. There are n districts and m two-way roads between them. The i-th road connects districts xi and yi. The cost of travelling along this road is wi. There is some path between each pair of districts, so the city is connected.
// There are k delivery routes in Berlyatov. The i-th route is going from the district ai to the district bi. There is one courier on each route and the courier will always choose the cheapest (minimum by total cost) path from the district ai to the district bi to deliver products.
// The route can go from the district to itself, some couriers routes can coincide (and you have to count them independently).
// You can make at most one road to have cost zero (i.e. you choose at most one road and change its cost with 0).
// Let d(x,y) be the cheapest cost of travel between districts x and y.
// Your task is to find the minimum total courier routes cost you can achieve, if you optimally select the some road and change its cost with 0. In other words, you have to find the minimum possible value of ∑i=1kd(ai,bi) after applying the operation described above optimally.

// Input
// The first line of the input contains three integers n, m and k (2≤n≤1000; n−1≤m≤min(1000,n(n−1)2); 1≤k≤1000) — the number of districts, the number of roads and the number of courier routes.
// The next m lines describe roads. The i-th road is given as three integers xi, yi and wi (1≤xi,yi≤n; xi≠yi; 1≤wi≤1000), where xi and yi are districts the i-th road connects and wi is its cost. It is guaranteed that there is some path between each pair of districts, so the city is connected. It is also guaranteed that there is at most one road between each pair of districts.
// The next k lines describe courier routes. The i-th route is given as two integers ai and bi (1≤ai,bi≤n) — the districts of the i-th route. The route can go from the district to itself, some couriers routes can coincide (and you have to count them independently).

// Output
// Print one integer — the minimum total courier routes cost you can achieve (i.e. the minimum value ∑i=1kd(ai,bi), where d(x,y) is the cheapest cost of travel between districts x and y) if you can make some (at most one) road cost zero.

// Examples

// input
// 6 5 2
// 1 2 5
// 2 3 7
// 2 4 4
// 4 5 2
// 4 6 8
// 1 6
// 5 3
// output
// 22

// input
// 5 5 4
// 1 2 5
// 2 3 4
// 1 4 3
// 4 3 7
// 3 5 2
// 1 5
// 1 3
// 3 3
// 1 5
// output
// 13

// Note
// The picture corresponding to the first example:
// There, you can choose either the road (2,4) or the road (4,6). Both options lead to the total cost 22.
// The picture corresponding to the second example:
// There, you can choose the road (3,4). This leads to the total cost 13.

// Tutorial
// If we would naively solve the problem, we would just try to replace each edge's cost with zero and run Dijkstra algorithm n times to get the cheapest paths. But this is too slow.
// Let's try to replace each edge's cost with zero anyway but use some precalculations to improve the speed of the solution. Let's firstly run Dijkstra n times to calculate all cheapest pairwise paths. Then, let's fix which edge we "remove" (x,y).
// There are three cases for the path (a,b): this edge was not on the cheapest path before removing and is not on the cheapest path after removing. Then the cost of this path is d(a,b). The second case is when this edge was not on the cheapest path before removing but it is on the cheapest path after removing. Then the cost of this path is min(d(a,x)+d(y,b),d(a,y)+d(x,b)). So we are just going from a to x using the cheapest path, then going through the zero edge and then going from y to b using the cheapest path also (or vice versa, from a to y and from x to b). And the third case is when this edge was already on the cheapest path between a and b but this case is essentially the same as the second one.
// So, if we fix the edge (x,y), then the answer for this edge is ∑i=1kmin(d(ai,bi),d(ai,x)+d(y,bi),d(ai,y)+d(x,bi)). Taking the minimum over all edges, we will get the answer.
// The precalculating part works in O(nmlogn) and the second part works in O(km).

#include <iostream>
#include <queue>
#include <set>
#include <utility>
#include <vector>

void dijkstra(const std::vector<std::vector<std::pair<unsigned, unsigned>>> & adj, unsigned s, std::vector<unsigned> & dist) {
    unsigned u, v, w, d, i;
    std::priority_queue<std::pair<unsigned, unsigned>, std::vector<std::pair<unsigned, unsigned>>, std::greater<std::pair<unsigned, unsigned>>> pq;

    dist.resize(adj.size());
    dist.assign(adj.size(), 1e9);
	dist[s] = 0;
	pq.push(std::make_pair(dist[s], s));

	while (!pq.empty()) {
        d = pq.top().first;
		u = pq.top().second;
		pq.pop();

        if (d > dist[u]) {
            continue;
        }

        for (i = 0; i < adj[u].size(); i++) {
            v = adj[u][i].first;
            w = adj[u][i].second;

			if (dist[v] > dist[u] + w) {
                dist[v] = dist[u] + w;
                pq.push(std::make_pair(dist[v], v));
			}
		}
	}
}

int main() {
	unsigned n, m, k, u, v, w, a, b, ans, cur, i, j;
    std::vector<std::vector<std::pair<unsigned, unsigned>>> adj;
    std::vector<std::pair<unsigned, unsigned>> route;
    std::vector<std::vector<unsigned>> d;

	while (std::cin >> n >> m >> k) {
        adj.resize(n);
        adj.assign(n, std::vector<std::pair<unsigned, unsigned>>());
        route.resize(k);
        d.resize(n);
        ans = 1e9;

        for (i = 0; i < m; i++) {
            std::cin >> u >> v >> w;
            u--;
            v--;
            adj[u].push_back({v, w});
            adj[v].push_back({u, w});
        }
        
        for (i = 0; i < k; i++) {
            std::cin >> a >> b;
            a--;
            b--;
            route[i] = std::make_pair(a, b);
        }

        for (u = 0; u < n; u++) {
            dijkstra(adj, u, d[u]);
        }
        
        for (u = 0; u < n; u++) {
            for (i = 0; i < adj[u].size(); i++) {
                v = adj[u][i].first;
                w = adj[u][i].second;
                cur = 0;

                for (j = 0; j < route.size(); j++) {
                    a = route[j].first;
                    b = route[j].second;

                    if (d[a][b] <= d[a][u] + d[v][b] && d[a][b] <= d[a][v] + d[u][b]) {
                        cur += d[a][b];
                    } else if (d[a][u] + d[v][b] <= d[a][v] + d[u][b]) {
                        cur += d[a][u] + d[v][b];
                    } else {
                        cur += d[a][v] + d[u][b];
                    }
                }

                if (ans > cur) {
                    ans = cur;
                }
            }
        }
        
        std::cout << ans << std::endl;
    }
	
	return 0;
}
