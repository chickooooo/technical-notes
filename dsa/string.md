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

### Remove outermost parentheses

- Problem: Given a valid parentheses string s, split it into its primitive components. For each primitive part, remove its outermost parentheses, then return the resulting string.
- Link: https://leetcode.com/problems/remove-outermost-parentheses/

---

- **Algorithm**: Depth counter
- Maintain a counter `depth` that indicates how deep we are in a nested parentheses.
- While iterating over the string `s`, ignore `(` at depth 1 and `)` at depth 0.
- For rest of the parentheses, add them to a new string and return it.

<br>
<br>
<br>

### Rotate string

- Problem: Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s. A shift on s consists of moving the leftmost character of s to the rightmost position.
- Link: https://leetcode.com/problems/rotate-string/

---

- **Algorithm**: Rotated string
- When searching in a rotated string or array, do `double = s+s`.
- Iterate over `double`, the rotated part starting from index `i` will be `double[i:i+n]`.

<br>
<br>
<br>

### Count and say

- Problem: The count-and-say sequence starts with "1". Each next term is formed by run-length encoding the previous term (counting consecutive digits and saying the count followed by the digit). Given a positive integer n, return the n-th term of the count-and-say sequence.
- Link: https://leetcode.com/problems/count-and-say/

---

- **Algorithm**: Count of consecutive same elements
- Start with `running_str = "1"` and iteratively loop till `n` is greater than 1.
- At each step, find the count of consecutive same elements and add `"{count}{num}"` to the `temp` string.
- After the loop, update `running_str` with `temp`. Return `running_str` at the end.

<br>
<br>
<br>

### 

<br>
<br>
<br>
