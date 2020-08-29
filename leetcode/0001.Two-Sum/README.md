# 0001. [Two Sum](https://leetcode.com/problems/two-sum/)

## 題目

Given an array of integers `nums` and and integer `target`, return the indices of the two numbers such that they add up to `target`.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

example:

```text
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1]
```

## 題目大意

在陣列中找到兩個數，它們的和等於給定的數字，並回傳那兩個數在陣列中的位置。

## 思路

首先宣告一 map 儲存陣列中的數 n 與位置 i 的關係 `( map[int]int{ n: i } )`。
掃過陣列中的每一個數 n，如果能在 map 中找到 target - n = 0 的數，
就可以直接回傳 target - n 的位置，與當前 n 的位置。
如果沒找到，則把當前數字 n 與其位置存入 map。

```text
以上面 example 為例掃過 nums:
n = 2 時: map[9-2] 找不到，把其值與位置存入 map (map[2] = 0)
n = 7 時: map[9-7] 有找到值，則可以直接回傳 {0, 1} (map[2] = 0 與當前 n 的位置 1)
```
