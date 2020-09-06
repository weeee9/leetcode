# 0925.[Long Pressed Name](https://leetcode.com/problems/long-pressed-name/)

## 題目

Your friend is typing his name into a keyboard.  Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard.  Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.

**Example 1:**

```c
Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.
```

**Example 2:**

```c
Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.
```

**Example 3:**

```c
Input: name = "leelee", typed = "lleeelee"
Output: true
```

**Example 4:**

```c
Input: name = "laiden", typed = "laiden"
Output: true
Explanation: It's not necessary to long press any character.
```

## 題目大意

比對正確的 `name` 與輸入的 `typed`，檢查輸入的某個字母是否有長壓按鍵的情況。

## 思路

運用滑動窗口的概念，兩個字串一起比較，如果是相同的字元，窗口就往後滑動。
直到遇見不同的字元，如果窗口的字元與前一個字元不符合的話，就可以直接回傳 `false`。
