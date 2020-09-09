# 202.[Happy Number](https://leetcode.com/problems/happy-number/)

## 題目

Write an algorithm to determine if a number n is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Return True if n is a happy number, and False if not.

**Example:**
```c
Input: 19
Output: true
Explanation: 
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
```

## 題目大意

判斷一個數是不是快樂數。快樂樹的定義是，不斷的把這個數的每個數字的平方加起來，如果最後等於 1，那就是快樂數。

## 思路

按照題目意思實作即可。