# [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

## 題目
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

就是說, x軸上在1,2,...,n點上有許多垂直的線段,長度依次是a1,a2,...,an. 找出兩條斷線,使他們和x抽圍成的面積最大.面積公式是 Min(ai,aj) X |j - i|

## 解題思路
窮舉法是O(n^2)的複雜度 會觸發leetcode的時間限制.

O(n)的複雜度的解法是，保持兩個指針i,j; 分別指向長度數組的首尾 如果ai小於aj 則移動i向後(i++) 反之 移動j向前(j--) 如果當前的area大於了所記錄的area 替換之. 這個想法的基礎是,如果i的長度小於j 無論如何移動j 短版在i 不可能找到比當前紀錄的area更大的值了,只能通過移動i來找到新的可能得更大面積