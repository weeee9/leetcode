# 0071. [Simplify Path](https://leetcode.com/problems/simplify-path/)

## 題目

Given an **absolute path** for a file (Unix-style), simplify it. Or in other words, convert it to the **canonical path**.

In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level.

Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. The last directory name (if it exists) **must not** end with a trailing /. Also, the canonical path must be the **shortest** string representing the absolute path.

**Example 1:**

```c
Input: "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
```

**Example 2:**

```c
Input: "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
```

**Example 3:**

```c
Input: "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
```

**Example 4:**

```c
Input: "/a/./b/../../c/"
Output: "/c"
```

**Example 5:**

```c
Input: "/a/../../b/../c//.//"
Output: "/c"
```

**Example 6:**

```c
Input: "/a//b////c/d//././/.."
Output: "/a/b/c"
```

## 題目大意

給一個字串代表 **Unix** 的路徑，回傳一簡化路徑的字串。

## 思路

使用 `stack` 記錄及判斷。
