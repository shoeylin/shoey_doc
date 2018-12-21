#[4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## 題目
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
```
nums1 = [1, 3]
nums2 = [2]
The median is 2.0
```
Example 2:
```
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5
```
## 解題思路
輸入兩個已經排序好的整數切片 合併兩個切片 返回合併後切片的中位數

所以題目的關鍵是 如何高校的合併切片

## 總結
分步驟完成程序