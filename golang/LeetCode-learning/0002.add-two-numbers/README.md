# [2. Add Two Numbers]
(https://leetcode.com/problems/add-two-numbers/)

## 題目

You are given two non-empty linked lists representing two
non-negative integers. The digits are stored in reverse order
and each of their nodes contain a single digit. Add the two
numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero,
except the number 0 itself.

```text
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
```

##解題思路

```text
(2 -> 4 -> 3)是 342

(5 -> 6 -> 4)是 465

(7 -> 0 -> 8)是 807

342+465=807
```

所以 題目的本意是 把整數算了一種表達方式 實現期加法

設計程序的時候 需要處理的點有

1.位上的加法 要處理進位問題
2.如何進入下一運算
3.按位相加結束后 也還需要處理進位問題

## 總結

讀懂提議後 按步驟實現題目要求