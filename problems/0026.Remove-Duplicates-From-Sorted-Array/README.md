# 0026. [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)

## 題目

Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifyng the input array **in-place with O(1) extra memory.**

**Example 1:**

```c
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
```

**Example 2:**

```c
Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
```

**Clarification:**

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

```c
// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```

## 題目大意

給一個排序過的陣列 `nums`，刪除 `nums` 中重複的元素，使每個元素在 `nums` 中只出現一次，並且回傳刪除重複元素後的 `nums` 長度。

## 思路

根據上面的 Clarification，這題的刪除並不是真的刪除，而是把重複的元素移到陣列後方，然後回傳刪除後 `nums` 的長度。
題目在評判的時候會讀取 `nums` 並根據回傳的長度輸出檢查。
