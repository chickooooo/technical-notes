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

- **Algorithm**: Reverse of reverse is the original.
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

- **Algorithm**: Simple string traversal.
- At each step, add current value to total.
- If current roman value is greater than previous: `total += -2 * previous`.

<br>
<br>
<br>

### Sum of frequency difference of all substrings

- Problem: Given a string s, return the sum of "difference in frequencies between the most frequent and least frequent characters" of all of its substrings.

---

- **Algorithm**: 2 `for` loops and frequency array.
- Iterate the string using 2 for loops in `O(n^2)` time complexity.
- At each step, use frequency array to calculate the difference between max & min frequency.
- Total time complexity: `O(n*n*26)`

<br>
<br>
<br>

### Longest palindromic substring

- Problem: Given a string s, return the longest palindromic substring in s.
- Link: https://leetcode.com/problems/longest-palindromic-substring/

---

- **Algorithm**: Check Palindrome from middle
- Iterate the given string `s`.
- At each step, find the longest palindromic substring with centre as `s[i]`. Odd length palindromic substring.
- Also, find the longest palindromic substring with centre as `s[i], s[i+1]`. Even length palindromic substring.
- Keep track of longest palindromic substring and return it.

<br>
<br>
<br>

### Sort characters by frequency

- Problem: Given a string s, sort it in decreasing order based on the frequency of the characters. Return the sorted string. If there are multiple answers, return any of them.
- Link: https://leetcode.com/problems/sort-characters-by-frequency/

---

- **Algorithm**: String + hashmap
- Calculate frquency of characters using `Counter(s)`.
- Create an array of `(-freq, key)` and sort the array.
- Iterate through the sorted array and construct the output.

<br>
<br>
<br>

###
