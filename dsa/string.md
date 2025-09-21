## Important Concepts

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Max depth of nested parentheses
- Simple iteration.
- Depth increases with `(` and decreases with `)`.

<br>
<br>
<br>

### Reverse the order of words
- Problem:
    ```
    input = "a good   example"
    output = "example good a"
    ```

---

- Algorithm: Reverse of reverse is the original.
- Reverse the whole string while handling the spaces.
- Reverse individual words.

<br>
<br>
<br>

### Roman number to integer

- Problem: Convert roman number string to its representative integer. The string only contains `I(1), V(5), X(10), L(50), C(100), D(500), M(1000))`.
    ```
    input = "MCMXCIV"
    output = 1994
    ```

---

- Algorithm: Simple string traversal.
- At each step, add current value to total.
- If current roman value is greater than previous: `total += -2 * previous`.
