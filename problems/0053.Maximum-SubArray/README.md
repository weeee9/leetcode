# 0053. [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

##  題目

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

**Example:**
```c
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

## 題目大意

給一陣列 `nums`，在 `nums` 中找到一連續的子陣列，其所有的元素的和為 `nums` 中最大。

## 思路

依序掃過 `nums`，並把每一個元素 `nums[i]` 都當成子陣列的開頭，並和前面的子陣列的和做比對。
