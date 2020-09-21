# 263.[Ugly Number](https://leetcode.com/problems/ugly-number/)

## 題目

Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

**Example 1:**
```c
Input: 6
Output: true
Explanation: 6 = 2 × 3
```

**Example 2:**

```c
Input: 8
Output: true
Explanation: 8 = 2 × 2 × 2
```

**Example 3:**

```c
Input: 14
Output: false 
Explanation: 14 is not ugly since it includes another prime factor 7.
```

Note:

1. `1` is typically treated as an ugly number.
2. Input is within the 32-bit signed integer range: [−2^31,  2^31 − 1].

## 題目大意

給一整數，判斷其是不是 `Ugly Number`。
`Ugly Number` 的定義是其質因數只包含 `2, 3, 5`。

## 思路

按照題目要求實作即可