### Important Concepts

<br>
<br>

**Prefix, Postfix, Infix**
- Prefix to ANY:
    - Iterate backwards on input.
    - Reverse the output.

<br>
<br>
<br>

### Problems:

<br>
<br>

**Infix to postfix conversion**
- Use stack to store the operators, as per their priority order.
- Keep on adding operands to the output string.

<br>

**Postfix to prefix conversion**
- Use stack to store operands.
- When an operator arrives, replace last operand with `{operator}{left}{right}`.

<br>

**Postfix to infix conversion**
- Use stack to store operands.
- When an operator arrives, replace last operand with `({left}{operator}{right})`.

<br>

**Prefix to postfix conversion**
- Iterate from the back.
- Use stack to store operands.
- When an operator arrives, replace last operand with `{operator}{left}{right}`.
- Return **reversed** output.

<br>

**Trapping rain water**
- Create a decreasing monotonic stack. This stack will hold `[height, width]`
- When a smaller element appears, add `[element, 1]` to the stack.
- When a bigger number appears:
    - Calculate the water limit `min(stack[0][0], element)`.
    - Pop each pair from stack where `stack[-1][0] <= element`. Calculate the water stored.
    - At the same time, keep updating the `width` of the current element.
    - At last add `[element, width]` to the stack.
---
- O(1) space complexity approach
 