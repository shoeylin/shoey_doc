#[5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

# 題目
Given a string s, find the longest palindromic substring in s. You may asuume that the maximum length of s is 1000.

Example:
```
Input: "babad"
Optput: "bab"
Note: "aba" is also valid answer.
```
Exmaple:
```
Input: "cbbd"
Output: "bb"
```

## 解題思路
題目要求尋找字符中的最常回文
palindromic 是回文，也就是一個字串從前面念過去和從後面念過來是相同的，也就是左右對稱
當然 我們可以使用下面程序判斷字符串s[i:j+1]是否是回文
```go
func isPalindromic(s *string, i, j int)bool{
    for i< j{
        if (*s)[i] != (*s)[j]{
            return false
        }
        i++
        j--
    }
    return true
}
```
但是 這樣就沒有利用回文的一下特點 假定回文的長度為l x 為任意字符
1. 當l為奇數時 回文的`正中間段`只會是 "x" 或 "xxx" 或 "xxxxx" ...
2. 當l為偶數時 回文的`正中間段`只會是 "xx" 或 "xxxx" 或 "xxxxxx"
3. 同一個字符串的任一兩個回文substring的`正中間段` 不會重疊

## 總結充分利用查找對象的特點 可以加快查找速度