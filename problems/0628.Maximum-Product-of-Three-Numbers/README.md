# 0628.[Maximum Product of Three Numbers](https://leetcode.com/problems/maximum-product-of-three-numbers/)

## 題目

Given an integer array, find three numbers whose product is maximum and output the maximum product.

**Example 1:**

```c
Input: [1,2,3]
Output: 6
```

**Example 2:**

```c
Input: [1,2,3,4]
Output: 24
```

**Note:**


1. The length of the given array will be in range [3,10^4] and all elements are in the range [-1000, 1000].
2. Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.


## 題目大意

給一陣列，找陣列中的三個數，他們的乘積是最大的。

## 思路

乘積最大的數有兩種情況：

1. 三個正數 
2. 一正數和兩個負數

所以我們只要找到最大的三個數，和最小的兩個數，然後再進行乘積的比對即可。
這題可以先排序再找，但時間複雜度會比較高；也可以設五個變數分別代表最大的三個數及最小的兩個數，時間複雜度就可以是 `O(n)`。