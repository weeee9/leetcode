# 1598. [Crawler Log Folder](https://leetcode.com/problems/crawler-log-folder/)

## 題目

The Leetcode file system keeps a log each time some user performs a *change folder* operation.

The operations are described below:

* `"../"` : Move to the parent folder of the current folder. (If you are already in the main folder, **remain in the same folder**).
* `"./"` : Remain in the same folder.
* `"x/"` : Move to the child folder named x (This folder is **guaranteed to always exist**).

You are given a list of strings `logs` where `logs[i]` is the operation performed by the user at the `ith` step.

The file system starts in the main folder, then the operations in `logs` are performed.

Return the *minimum number of operations* needed to go back to the main folder after the change folder operations.

**Example 1:**

```c
Input: logs = ["d1/","d2/","../","d21/","./"]
Output: 2
Explanation: Use this change folder operation "../" 2 times and go back to the main folder.
```

**Example 2:**

```c
Input: logs = ["d1/","d2/","./","d3/","../","d31/"]
Output: 3
```

**Example 3:**

```c
Input: logs = ["d1/","../","../","../"]
Output: 0
``` 

**Constraints:**

* `1 <= logs.length <= 103`
* `2 <= logs[i].length <= 10`
* `logs[i]` contains lowercase English letters, digits, `'.'`, and `'/'`.
* `logs[i]` follows the format described in the statement.
* Folder names consist of lowercase English letters and digits.

## 題目大意

給一字串陣列，裡面包含 **change folder** 的操作 ( `cd` command in linux)，根據陣列中的操作，
返回最後最少需要多少次操作才能回到 `main` 資料夾。

## 思路

用 stack 的方式實作。
