# 231.[Power of Two](https://leetcode.com/problems/power-of-two/)

## 題目

Given an integer, write a function to determine if it is a power of two.

**Example 1:**

```c
Input: 1
Output: true 
Explanation: 2^0 = 1
```

**Example 2:**

```c
Input: 16
Output: true
Explanation: 2^4 = 16
```

**Example 3:**

```c
Input: 218
Output: false
```
## 題目大意

給一整數，判斷其是不是 2 的冪次方。

## 思路

觀察 **2 的次方** 及 **2 的次方減 1** 轉 2 進制的關係：

```c
// 2 的次方          // 2 的次方減 1
1 ==> 0000 0001    (1-1 = 0) ===> 0000 0000
2 ==> 0000 0010    (2-1 = 1) ===> 0000 0001
4 ==> 0000 0100    (4-1 = 3) ===> 0000 0011
8 ==> 0000 1000    (8-1 = 7) ===> 0000 0111
16 => 0001 0000    (16-1 = 15) => 0000 1111
32 => 0010 0000    (32-1 = 31) => 0001 1111
...
```

可以發現 2 的次方和 2 的次方減一之後做 `AND` 運算會等於 0，了解之後實作的邏輯就很簡單。

另外當然也可以用迴圈的方式來實作。
