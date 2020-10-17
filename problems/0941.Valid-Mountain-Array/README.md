# 941. [Valid Mountain Array](https://leetcode.com/problems/valid-mountain-array/)

## 題目

Given an array `A` of integers, return `true` if and only if it is a valid mountain array.

Recall that A is a mountain array if and only if:

* `A.length >= 3`
* There exists some i with 0 < i < A.length - 1 such that:
  * `A[0] < A[1] < ... A[i-1] < A[i]`
  * `A[i] > A[i+1] > ... > A[A.length - 1]`

**Example 1:**

```c
Input: [2,1]
Output: false
```

**Example 2:**

```c
Input: [3,5,5]
Output: false
```

**Example 3:**

```c
Input: [0,3,2,1]
Output: true
```

**Note:**

1. `0 <= A.length <= 10000`
2. `0 <= A[i] <= 10000`

## 題目大意

給一整數陣列 `A`，判斷其是否符合山的形狀。
定義為：

* `A` 的長度要大於等於 3
* 其中的元素要符合:
  * `A[0] < A[1] < ... A[i-1] < A[i]`
  * `A[i] > A[i+1] > ... > A[A.length - 1]`

## 思路


