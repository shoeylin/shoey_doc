# [13. Roman to Integer](https://leetcode.com/problems/roman-to-integer/)

## 題目
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.

## 解題思路
這題與 [12. Integer to Roman]的一個逆轉換 一樣的解題思路

此題最關鍵的訊息是
> 右加左減 左減數字必須為一位 如8寫成VIII 而非 IIX

1. 從右往左處理字符串
2. 當前字浮代表數字 小於右邊的時候 總體減去當前字符代表的數字
3. 否則 總體加上當前字符代表的數字

## 總結
抓住關鍵訊息 避免思維定式