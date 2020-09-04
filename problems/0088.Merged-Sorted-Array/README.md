# 0088.[Merge-Sorted-Array](https://leetcode.com/problems/merge-sorted-array/description/)

## 題目

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.

**Example:**

```c
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
```

## 題目大意

給兩個排序過的陣列 `nums1` 及 `nums2`，把第二個陣列 `nums2` 合併到第一個陣列 `nums1` 中。

## 思路

依序選取兩個陣列中最大的數，從最後面開始放
