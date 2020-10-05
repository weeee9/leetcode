#

## 題目

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

**Note:** You may not slant the container and n is at least 2.

![question_11](./question_11.jpg)

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.



**Example:**

```c
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```

## 題目大意

給一整數陣列，選擇兩個元素其在陣列中的矩形面積是最大的。

其中元素的值當作矩形的高，元素間的距離當作矩形的寬。

## 思路

使用兩個指針分別從頭跟尾及算即可。