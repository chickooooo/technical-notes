# Math & Geometry

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Largest area of square inside rectangles

- Problem: You are given several axis-aligned rectangles on a 2D plane. Each rectangle is defined by its bottom-left and top-right coordinates. Determine the maximum possible area of a square that can be placed entirely inside the overlapping region of at least two rectangles. If no such square exists, return 0.
- Link: https://leetcode.com/problems/find-the-largest-area-of-square-inside-two-rectangles/

---

- **Algorithm**: Maximum intersecting rectangle
- A maximum intersecting rectangle can be created using exactly 2 rectangles. Once we prove this, the solution becomes easy.
- Form all pairs of 2 rectangles and find the maximum rectangle.
- Maximum square inside a rectangle is formed using the minimum side of that rectangle.
- Find the max square and return its area.

<br>
<br>
<br>

### Largest magic square

- Problem: Given an m x n grid of integers, find the largest possible size k such that a k by k subgrid is a magic square, where all rows, columns, and both diagonals have the same sum. The numbers do not need to be distinct.
- Link: https://leetcode.com/problems/largest-magic-square/

---

- **Algorithm**: Prefix sum
- Create multiple prefix sum grids: row-wise, column-wise, diagonal-wise and anti-diagonal-wise.
- Now iterate through the main grid and find the largest magic square.
- Also implement early stopping wherever possible.

<br>
<br>
<br>

### 

<br>
<br>
<br>
