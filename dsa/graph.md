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

### 

<br>
<br>
<br>
