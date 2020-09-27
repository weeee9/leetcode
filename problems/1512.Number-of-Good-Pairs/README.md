# 1512. [Number of Good Pairs](https://leetcode.com/problems/number-of-good-pairs/)

## 題目

Given an array of integers `nums`.

A pair `(i,j)` is called good if `nums[i]` == `nums[j]` and `i` < `j`.

Return the number of good pairs.

**Example 1:**

```c
Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
```

**Example 2:**

```c
Input: nums = [1,1,1,1]
Output: 6
Explanation: Each pair in the array are good.
```

**Example 3:**

```c
Input: nums = [1,2,3]
Output: 0
```

**Constraints:**

* `1 <= nums.length <= 100`
* `1 <= nums[i] <= 100`

## 題目大意

給一整數陣列 `nums`，判斷其中有多少個 **Good Pairs**。

**Good Pairs** 的定義為：
一對 *Good Pairs* `(i, j)`，且 `i`, `j` 滿足 `nums[i]` == `nums[j]` and `i` < `j`

## 思路

用一個 `map` 來紀錄 `nums[i]` 出現的次數，