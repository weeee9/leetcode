# 0720. [Longest Word in Dictionary](https://leetcode.com/problems/longest-word-in-dictionary/)

## 題目

Given a list of strings `words` representing an English Dictionary, find the longest word in `words` that can be built one character at a time by other words in `words`. If there is more than one possible answer, return the longest word with the smallest lexicographical order.

If there is no answer, return the empty string.

**Example 1:**

```c
Input: 
words = ["w","wo","wor","worl", "world"]
Output: "world"
Explanation: 
The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
```

**Example 2:**

```c
Input: 
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
Output: "apple"
Explanation: 
Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
```

**Note:**

* All the strings in the input will only contain lowercase letters.
* The length of `words` will be in the range `[1, 1000]`.
* The length of `words[i]` will be in the range `[1, 30]`.

## 題目大意

給一字串陣列 `words` 當作字典，找出其中最長的單字，該單字是由 `words` 中其他元素組成。
若有多個答案，回傳在字典中排序最小的那個；若無則回傳空字串。

## 思路

先把 `words` 排序，再用 `map` 紀錄。

依序掃過 `words` 中的每個元素 `word`，檢查在 `map` 中是否可以找到 `wword[:-1]` 的值。
