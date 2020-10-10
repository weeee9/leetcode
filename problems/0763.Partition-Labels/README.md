# 0763. [Partition Labels](https://leetcode.com/problems/partition-labels/)

## 題目

A string `S` of lowercase English letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

**Example 1:**

```c
Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
``` 

**Note:**

* `S` will *have length in range `[1, 500]`.
* `S` will consist of lowercase English letters (`'a'` to `'z'`) only.

## 題目大意

給一字串 `S`，把它盡可能地拆分成多組字串，使得每組拆分後的字串中所有字元都只會出現在該組裡面。

## 思路

首先紀錄每個字元出現的最後位置，再根據位置紀錄去判斷拆分的時機。
