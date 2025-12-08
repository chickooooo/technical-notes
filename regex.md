# Regular Expressions

<br>
<br>
<br>
<br>
<br>

## Index

- [What is a Regex?](#what-is-a-regex)
- [Basic Matchers](#basic-matchers)
- [Dot `.`](#dot-)
- [Character set `[]`](#character-set-)
- [Negated character set](#negated-character-set)

<br>
<br>
<br>
<br>
<br>

### What is a Regex?

- A Regex is a string of characters that express a search pattern.
- It is used to find or replace words in texts.
- In addition, we can test whether a text complies with the rules we set.
- Example: To find all the PDF files, we use `^\w+\.pdf$` regex.

<br>
<br>
<br>

### Basic Matchers

- Basic matchers match the exact string.
- The character or word we want to find is written directly.
- It is similar to a normal search process.

```
curious
```
```
I have no special talents. I am only passionately curious.
                                                  -------
```

- Matches only the characters `curious`.

<br>
<br>
<br>

### Dot `.`

- The period `.` allows selecting any single character.
- Including special characters and spaces.

```
.
```
```
abcABC123 .:!?
--------------
```

<br>
<br>
<br>

### Character set `[]`

- A character set matches any one character from a specified group of characters.
- It is written using square brackets `[]`.

```
[abc]ad
```
```
bad dad cad zap
---     ---
```

- Matches `aad` or `bad` or `cad`.

<br>
<br>
<br>

### Negated character set

- If the first character inside the brackets is `^`, it means match anything except these characters.

```
[^abc]ad
```
```
bad dad cad zap
    ---
```

- Matches any word that ends with `ad` and does not starts with `a` or `b` or `c`.

<br>
<br>
<br>

### 

<br>
<br>
<br>
