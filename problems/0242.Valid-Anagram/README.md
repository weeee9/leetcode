# 0242. [Valid Anagram](https://leetcode.com/problems/valid-anagram/)

## 題目

Given two strings s and t , write a function to determine if t is an anagram of s.

**Example 1:**

```c
Input: s = "anagram", t = "nagaram"
Output: true
```

**Example 2:**

```c
Input: s = "rat", t = "car"
Output: false
```

**Note:**
You may assume the string contains only lowercase alphabets.

**Follow up:**
What if the inputs contain unicode characters? How would you adapt your solution to such case?

## 題目大意

給兩個字串 `s` 及 `t`，判斷 `t` 中所有的字母是否在 `s` 都存在，是則回傳 `true`，反之回傳 `false`。

## 思路

用 `map` 或是用`陣列`的方式記錄。

