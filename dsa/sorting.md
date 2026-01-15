# Sorting

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Maximize area of square hole in grid

- Problem: Given a grid formed by horizontal and vertical bars, where some specified bars can be removed, determine the maximum possible area of a square-shaped hole that can be formed by removing any number of the allowed bars.
- Link: https://leetcode.com/problems/maximize-area-of-square-hole-in-grid/

---

- **Algorithm**: Maximum consecutive elements count
- Intersection of 2 bars form a point. If we remove points along a diagonal, the size of square keep on increasing.
- That means, longest diagonal line will give the square with the largest area.
- Longest diagonal line can be found using:
    - If `h_count` is max number of consecutive elements in `hBars`.
    - If `v_count` is max number of consecutive elements in `vBars`.
    - Max diagonal = `min(h_count, v_count)`.
    - Once we get the number of points on the longest diagonal, we get the largest area.
    - Note: Handle adjustments for off by 1 error.

###

<br>
<br>
<br>
