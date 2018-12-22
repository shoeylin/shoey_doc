# [6. ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)

## 題目

The string "PATPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```text
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

```go
func convert(text string, nRows int) string
```

convert("PATPALISHIRING", 3) should return
"PAHNAPLSIIGYIR".

## 解題思路

輸入"ABCDEFGHIJKLMNOPQRSTUVWXYYZ"和參數5後 得到答案
"AGMSYBFHLNRTXZCEIKOQUWDJPV"
按照題目的擺放方法 可得

```test
A    I   Q   Y
B   HJ  PR  XZ
C  G K O S W
DF   LN  TV
E    M   U
```

可以看到 各行字符在原字符串中的索引號為

0行 0,    8,        16,       24
1行 1,  7,9,     15,17,    23,25
2行 2, 6, 10,  14,  18,  22,
3行 3,5,  11,13,    19,21,
4行 4,    12,       20,

令p=numRows*2-2 可以總結出以下規律

0行 0*p,1*p ....
r行 r,1*p-r,1*p+r,2*p-r,2*p+r ....
最後一行 numRow-1, numRow-1+1*p numRow-1+2*p ....

只需編程一次處理各行即可