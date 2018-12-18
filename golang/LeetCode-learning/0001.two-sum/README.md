# [1. Two Sum](https://leetcode.com/problems/two-sum/)

##題目

Given an array of integers return indices of the two numbers
such that they add up to a specific target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice

Example:

```text
Given nums = [2, 7, 11, 15], target = 9 ,

Bacause nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1],
```


## 解題思路

```go
a + b = target 
```

也可以看成是
```go
a =target - b
```

在`map[整數]整數的序號`中 可以查詢到a的序號 這樣就不用嵌套兩個for循環