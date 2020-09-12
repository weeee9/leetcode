# 1025.[Divisor Game](https://leetcode.com/problems/divisor-game/)

## 題目

Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number N on the chalkboard.  On each player's turn, that player makes a move consisting of:

* Choosing any `x` with `0 < x < N` and `N % x == 0`.
* Replacing the number `N` on the chalkboard with `N - x`.
  
Also, if a player cannot make a move, they lose the game.

Return `True` if and only if Alice wins the game, assuming both players play optimally.

**Example 1:**

```c
Input: 2
Output: true
Explanation: Alice chooses 1, and Bob has no more moves.
```

**Example 2:**

```c
Input: 3
Output: false
Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.
```

**Note:**

1. `1 <= N <= 1000`

## 題目大意

Alice 和 Bob 兩人玩一遊戲除，由 Alice 先手。
遊戲規則是先給定一整數 `N`，接著兩人輪流做出下列步驟

* 選定一整數 `x` 滿足 `0 < x < N` 且 `N % x == 0`
* 使用 `N - x` 取代 `N`

當其中一人無法選出 `x` 時他就輸了。

判斷 Alice 會贏的話，回傳 `true`，反之回傳 `false`

## 思路

當 `N` 為質數時，該輪的玩家必輸，因為無法滿足 `0 < x < N` 且 `N % x == 0` 條件，所以策略就是讓對方面對 `N` 為質數的情況。

如果一開始 `N` 為偶數，先手的人就可以必贏，只要選 `x = 1` 一直讓對方在奇數的狀況下，最後一定會面臨到質數的情況。

反之如果一開始 `N` 為奇數，只要對手持續選  `x = 1`，先手的人就只能一直面臨到奇數，那先手的人一定必輸。

```c
N = 10

Alice: x = 1 -> N = 10 - 1 = 9
Bob: x = 3 ---> N = 9 - 3 = 6

Alice: x = 1 -> N = 6 - 1 = 5
Bob: LOSE
```

```c
N = 9

Alice: x = 3 -> N = 9 - 3 = 6
Bob: x = 1 ---> N = 6 - 1 = 5

Alice: LOSE
```
