# 0167. [Two Sum II - Input array is sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)

## 題目

Given an array of integers that is already ***sorted in ascending order***, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

**Note:**

* Your returned answers (both index1 and index2) are not zero-based.
* You may assume that each input would have exactly one solution and you may not use the same element twice.

**Example 1:**

```c
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
```

**Example 2:**

```c
Input: numbers = [2,3,4], target = 6
Output: [1,3]
```

**Example 3:**

```c
Input: numbers = [-1,0], target = -1
Output: [1,2]
```

**Constraints:**

* `2 <= nums.length <= 3 * 104`
* `-1000 <= nums[i] <= 1000`
* `nums` is sorted in **increasing order**.
* `-1000 <= target <= 1000`

## 題目大意

給一排序過的整數陣列 `nums` 及一整數目標 `target`，在 `nums` 中找到兩個數的和等於 `target`，並回傳它們的位置。

## 思路

用雙指針的方式尋找。
