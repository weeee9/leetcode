# 0028.[Implement strStr()](https://leetcode.com/problems/implement-strstr/)

## 題目

Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

**Example 1:**

```c
Input: haystack = "hello", needle = "ll"
Output: 2
```

**Example 2:**

```c
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

**Clarification:**

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

## 題目大意

實作 `strStr()` 函數，返回字串中第一個遇到 `needle` 字串的位置，如果沒有則返回 -1。
此外，當 `needle` 為空字串時，返回 0。

## 思路

依序跑過 `haystack` 字串找 `needle` 即可。
