# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 題目
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b" with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

##　解題思路
利用s[left:i+1]來表示s[:i+1]中的包含s[i]的最長字符串。
location[s[i]]是字符串s[i]在s[:i+1]中倒數第二次出現的序列號。
當left < location[s[i]]的時候 說明字符s[i]出現了兩次。需要設置 left = location[s[i]] + 1,
保證字符s[i]只出現一次。

補充說明:
substring 子字符串: 連續的字符
subsequence 子序列: 可包含不連續的字符
題目是找不重複字母最長的子字符串 不是找最長子序列

## 總結
利用Location保存字符上次出現的序列號 避免了查詢工作。location和[Two Sum](./algorithms/0001.two-sum)中的m是一樣的作用。
```go
// m 負責保存map[整數]整數的序列號
    m := make(map[int]int, len(nums))
```