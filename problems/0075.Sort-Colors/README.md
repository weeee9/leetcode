# 0075. [Sort Colors](https://leetcode.com/problems/sort-colors/)

## 題目
Given an array `nums` with `n` objects colored red, white, or blue, sort them **in-place** so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

Here, we will use the integers `0`, `1`, and `2` to represent the color red, white, and blue respectively.

**Follow up:**

* Could you solve this problem without using the library's sort function?
* Could you come up with a one-pass algorithm using only `O(1)` constant space?

**Example 1:**

```c
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

**Example 2:**

```c
Input: nums = [2,0,1]
Output: [0,1,2]
```

**Example 3:**

```c
Input: nums = [0]
Output: [0]
```

**Example 4:**

```c
Input: nums = [1]
Output: [1]
```

## 題目大意

給一整數陣列，然後進行排序。

## 思路

陣列中只會出現 0, 1, 2，所以可以用指針的方式重新擺放。
0 擺在最前面，所以放一個 0 的時候 1 和 2 也要放；放 1 的時候 2 也要一起放；最後的放 2 只要移動指針就好。每次擺放指針的值都要加一。
