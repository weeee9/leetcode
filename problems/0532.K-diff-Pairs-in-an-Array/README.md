# 0532.[K-diff Pairs in an Array](https://leetcode.com/problems/k-diff-pairs-in-an-array/)

## 題目

Given an array of integers and an integer **k**, you need to find the number of **unique** k-diff pairs in the array. Here a **k-diff** pair is defined as an integer pair (i, j), where **i** and **j** are both numbers in the array and their absolute difference is **k**.

**Example 1:**

```c
Input: [3, 1, 4, 1, 5], k = 2
Output: 2
Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
Although we have two 1s in the input, we should only return the number of unique pairs.
```

**Example 2:**

```c
Input:[1, 2, 3, 4, 5], k = 1
Output: 4
Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).
```

**Example 3:**

```c
Input: [1, 3, 1, 5, 4], k = 0
Output: 1
Explanation: There is one 0-diff pair in the array, (1, 1).
```

**Note:**

1. The pairs (i, j) and (j, i) count as the same pair.
2. The length of the array won't exceed 10,000.
3. All the integers in the given input belong to the range: [-1e7, 1e7].

## 題目大意

給一整數陣列以及一目標數 `k`，在陣列中找到出有幾對數字相減等於 `k` 的組合。
註： `(i, j)` 和 `(j, i)` 視為同一個。

## 思路

使用 `map` 來記錄每個數字出現的次數。
再掃過 `map` 來判斷該數字加上 `k` 之後是否在 `map` 中存在；或是 `k = 0` 時，該數字是否有出現兩次。