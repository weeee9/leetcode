# 0844.[Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)

## 題目

Given two strings `S` and `T`, return if they are equal when both are typed into empty text editors. `#` means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

**Example 1:**

```c
Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
```

**Example 2:**

```c
Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
```

**Example 3:**

```c
Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
```

**Example 4:**

```c
Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
```

**Note:**

* `1 <= S.length <= 200`
* `1 <= T.length <= 200`
* `S` and `T` only contain lowercase letters and `'#'` characters.

## 題目大意

給兩個字串 `S` 和 `T`，如果遇到 `'#'` 就回退一個字元，判斷最後兩個字串是否相等。

## 思路

使用 `stack` 來記錄。