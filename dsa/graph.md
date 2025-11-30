# Graph

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

### 

<br>
<br>
<br>
