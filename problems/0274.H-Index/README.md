# 0274. [H-Index](https://leetcode.com/problems/h-index/)

## 題目

Given an array of citations (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.

According to the `definition of h-index on Wikipedia`: "A scientist has index h if h of his/her N papers have **at least** h citations each, and the other N − h papers have **no more than** h citations each."

**Example:**

```c
Input: citations = [3,0,6,1,5]
Output: 3
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had
             received 3, 0, 6, 1, 5 citations respectively.
             Since the researcher has 3 papers with at least 3 citations each and the remaining
             two with no more than 3 citations each, her h-index is 3.
```

**Note:** If there are several possible values for h, the maximum one is taken as the h-index.

## 題目大意

求 H-index。
一個人的 N 篇論文中至少有 h 次被引用，且其他 N-h 篇論文被引用的次數不超過 h，該 h 就是那個人的 H-index。

## 思路

先排序論文被引用次數 `[3, 0, 6, 1, 5] -> [0, 1, 3, 5, 6]`。
根據 `H-index = 3` 可以反推排序過後，第 i 篇論文的被引用次數至少要大於等於 `N-i`。
