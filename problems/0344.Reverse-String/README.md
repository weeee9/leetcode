# 0344. [Reverse String](https://leetcode.com/problems/reverse-string/)

# 題目
Write a function that reverses a string. The input string is given as an array of characters char[].

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

You may assume all the characters consist of printable ascii characters.

*Example 1:*

```c
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
```

*Example 2:*

```c
Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
```

# 題目大意

實現一函數可以把字串反過來，但空間限制 ｀O(1)`。

# 思路

用兩個 index 分別從頭跟尾向中間出發，同時交換 index 位置的字元即可。