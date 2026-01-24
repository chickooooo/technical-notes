# Graph

<br>
<br>
<br>
<br>
<br>

## Algorithms:

### Dijkstra's Algorithm:

- It is used to find the shortest path from the `start` node to all other nodes of the graph.
- It can be used on both directed and undirected graphs.
- All the edges of the graph should have **non-negative weights**.

<br>

#### Core concepts

- It uses **Greedy approach** to select the next node that will be visited.
- The node with the lowest next distance is selected to visit.
- It makes use of **Priority Queue (heap)** to get the next node with the lowest distance.

<br>

#### Algorithm

- Create an array `shortest` and initialize distance of all the nodes to `INF`.
- Create a priority queue that will hold `(distance, node)`.
- Insert `(0, start)` in the PQ.
- Pop the min item from the PQ.
    - If the `distance` is less than `shortest[node]`:
        - Update `shortest[node]` with `distance`.
        - Add neighbours of `node` in the PQ.
    -  If the `distance` is greater or equal to `shortest[node]`, skip that item.
- Keep repeating above step until the PQ is empty.
- Return the `shortest` array in the end.

<br>
<br>
<br>

### Minimum spanning tree

- A spanning tree is a tree with N nodes and N-1 edges.
- Every node in a spanning tree is reachable from every other node. Graph with a single component.
- A minimum spanning tree is a spanning tree whose sum of **edge weights** is minimum.

---

- This algorithm works only on **Undirected graphs**.
- Edge weights of this graph can be negative, zero, or positive.
- Algorithms include:
    - Prim's Algorithm
    - Kruskal's Algorithm

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Number of provinces

- Problem: You are given an nxn matrix isConnected where `isConnected[i][j] = 1` means city i and city j are directly connected. Cities connected directly or indirectly form a province. Return the number of provinces.
- Link: https://leetcode.com/problems/number-of-provinces/

---

- **Algorithm**: Connected components
- Maintain a `visited` array of n nodes.
- Start from node `0`. Increment `count`. Peform DFS and mark all the connected nodes as visited.
- Start from next unvisited node, increment `count` and again perform DFS.
- Keep doing so until all the nodes are visited.
- Here, `count` represents the no. of provinces.
- Each time we start from an unvisited node, we are visiting a new province.

---

- **Algorithm**: Disjoint set
- Use standard **UnionByRank** or **UnionBySize** algorithm and connect each node to its respective component.
- Then, traverse the `parent` array and count all the nodes that are parent of themself (root nodes).
- Each component will have 1 root node and hence no. of roots = no. of provinces.

<br>
<br>
<br>

### Rotting oranges

- Problem: Given a grid with 0 = empty, 1 = fresh orange, and 2 = rotten orange, each minute all fresh oranges next to rotten ones become rotten. Return the minimum minutes needed for all oranges to rot, or -1 if it's impossible.
- Link: https://leetcode.com/problems/rotting-oranges/

---

- **Algorithm**: Multi-source BFS
- All the cells with value `2` are staring point of BFS.
- Perform standard BFS to get the time taken to rot all possible oranges.
- Design BFS algorithm such that there is always 1 extra level of traversal.
- This extra level will handle duplicate entries if present in any level.
- If done correctly, when we start from `t=0`, we will get the final time as correct time.
- Finally, to check if all fresh oranges were covered, use a counter to keep track of addition & rotting of fresh oranges.

<br>
<br>
<br>

### Nearest 0 in matrix

- Problem: Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell. The distance between two cells sharing a common edge is 1.
- Link: https://leetcode.com/problems/01-matrix/

---

- **Algorithm**: Multi-source BFS
- All the `0` cells are the starting point of BFS.
- Set `distance=0`, increment `distance` after each level of BFS.
- Use the same matrix for storing distance without using extra space.

<br>
<br>
<br>

### Surrounded regions

- Problem: Given an m × n board of 'X' and 'O', flip all 'O' regions that are completely surrounded by 'X'. An 'O' region is surrounded if none of its cells touch the board's edge. Replace all such 'O' cells with 'X' in-place.
- Link: https://leetcode.com/problems/surrounded-regions/

---

- **Algorithm**: DFS
- Traverse the horizontal and vertical edges of the matrix.
- If we encounter an 'O', perform DFS and mark all the connected 'O's.
- Now traverse the complete matrix.
    - If we encounter a marked 'O', make that cell 'O'.
    - Rest all the other cells will become 'X'.

<br>
<br>
<br>

### Course schedule

- Problem: There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai. Return true if you can finish all courses. Otherwise, return false.
- Link: https://leetcode.com/problems/course-schedule/

---

- **Algorithm**: Topological sort
- First create an adjacency list: `{node: [dependant]}`. Here `dependant` courses should be taken before taking node.
- Maintain 2 arrays:
    - `visited`: To keep track of nodes whose course is taken.
    - `temp`: Will hold nodes in current recursive stack.
- Iterate from `0` to `n-1` and perform DFS. Make use of `temp` array to check for cycles.
- If a cycle is found, stop and return `False`. Otherwise return `True`.

<br>
<br>
<br>

### Network delay time

- Problem: Given a directed, weighted graph with n nodes labeled 1 to n, where each edge (u, v, w) represents a signal travel time w from node u to node v, determine the minimum time required for a signal sent from node k to reach all nodes. If any node cannot be reached, return -1.
- Link: https://leetcode.com/problems/network-delay-time/

---

- **Algorithm**: Dijkstra's Algorithm
- Follow standard algorithm.
- Use `count: int` to keep track of visited nodes.
- Use `max_time: int` to keep track of max delay time.

<br>
<br>
<br>

### Shortest path in binary matrix

- Problem: Given an n × n binary matrix, find the length of the shortest path from the top-left cell (0,0) to the bottom-right cell (n−1,n−1) such that all visited cells contain 0 and movement is allowed in 8 directions (horizontal, vertical, and diagonal). If no such clear path exists, return −1.
- Link: https://leetcode.com/problems/shortest-path-in-binary-matrix/

---

- **Algorithm**: Dijkstra's Algorithm
- Standard Dijkstra's Algorithm just in 8 directions.
- If the top-left or bottom-right cell is 1, return `-1` early on.

<br>
<br>
<br>

### Number of ways to arrive at destination

- Problem: You are given a connected city with n intersections (0 to n−1) and bidirectional roads, each with a travel time. Starting from intersection 0, find how many different ways you can reach intersection n−1 using the shortest possible travel time.
- Link: https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/

---

- **Algorithm**: Dijkstra's Algorithm
- Use standard Dijkstra's Algorithm.
- While traversing keep track of the parent node.
- At each step, the count of shortest paths is current count + parent's shortest path count.

<br>
<br>
<br>

### Find eventual safe states

- Problem: Given a directed graph, find all nodes from which every possible path eventually leads to a terminal node (a node with no outgoing edges). Return these safe nodes in ascending order.
- Link: https://leetcode.com/problems/find-eventual-safe-states/

---

- **Algorithm**: Cycle detection in directed graph
- Each node will have 3 states:
    - `0`: not visited
    - `1`: visited
    - `2`: part of cycle
- Perform DFS traversal and find the state of each node.
- Return all the nodes that are not part of a cycle.

<br>
<br>
<br>

### Number of enclaves

- Problem: Given a binary grid where 1 is land and 0 is sea, count how many land cells are completely enclosed; meaning you cannot reach the grid boundary by moving up, down, left, or right through land cells.
- Link: https://leetcode.com/problems/number-of-enclaves/

---

- **Algorithm**: DFS on grid boundary
- Move along the boundary of the grid.
- Mark all the lands and their cells connected to the boundary as `0`.
- Now iterate over the grid and count all the `1` cells.

<br>
<br>
<br>

### Operations to make network connected

- Problem: Given n computers and a list of existing connections, where you can remove a cable and reconnect it between any two disconnected computers, find the minimum number of operations needed to connect all computers. Return -1 if it's impossible.
- Link: https://leetcode.com/problems/number-of-operations-to-make-network-connected/

---

- **Algorithm**: Disjoint set
- Note: To connect `n` nodes in a graph, we require atleast `n-1` edges.
- If number of connection is less than `n-1`, return `-1`. Else we have enough connections to connect all the computers.
- Suppose the graph has `m` components, then we need to move `m-1` connections to connect all components.
- Use Disjoint Set or DFS traversal to find the number of components `m` and return `m-1`.

<br>
<br>
<br>

### 

<br>
<br>
<br>
