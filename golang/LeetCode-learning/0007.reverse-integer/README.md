# [7. Reverse Integer](https://leetcode.com/problems/reverse-integer/)

## 題目
Given a 32-bit signed integer,Revers digits of an integer
```
Example1: x = 123 , return 321
Example2: x = -123 , return -321
```

## 解題思路
檢查正負存起來 x除以10的餘數去暫存至答案 解答再乘10回來 把正負還原
最後檢查有無超過int32