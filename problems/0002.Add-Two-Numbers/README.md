# 0001. [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

## 題目

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

example:

```text
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

## 題目大意

給兩個 linked list 代表兩個倒過來的數，將兩個數從低位開始相加，
得到的結果也使用 linked list 的方式倒過來表示。
