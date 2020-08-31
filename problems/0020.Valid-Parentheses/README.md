# 0020. [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/description/)

## 題目

Given a string containing just the characters ‘(’, ‘)’, ‘{’, ‘}’, ‘[’ and ‘]’, determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets. Open brackets must be closed in the correct order. Note that an empty string is also considered valid.

*Example 1:*
```c
Input: "()"
Output: true
```

*Example 2:*
```c
Input: "()[]{}"
Output: true
```

*Example 3:*
```c
Input: "(]"
Output: false
```

*Example 4:*
```c
Input: "([)]"
Output: false
```

*Example 5:*
```c
Input: "{[]}"
Output: true
```

## 題目大意

給一字串，判斷字串內的括號是否有成對。

## 思路

可以用 stack 的方式在遇到成對的括號時，就 pop 出來，反之則 push 進去，最後再判斷 stack 的長度是否為 0。

