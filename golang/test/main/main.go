package main

import (
	// "encoding/hex"
	"fmt"
	"time"
	// "github.com/axgle/mahonia"
)

func main() {

	t := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Println("123" + t)
	p := t[0:13] + t[14:16] + t[17:23]
	fmt.Println(p)
	fmt.Println()
	// strrrr := "a578bf6eb971"
	// strrrr1 := "abeda5cdabfc"

	// b1 := byte(0x0a)
	// b2 := byte(0x05)

	// b1 = b1 << 4
	// c1 := b1 + b2

	// b11 := byte(0x07)
	// b22 := byte(0x08)

	// b11 = b11 << 4
	// c11 := b11 + b22

	// fmt.Println("c1", c1, b1, b2)

	// src := []byte{c1, c11}

	// fmt.Println("byte src", string(src))

	// encodedStr := hex.EncodeToString(src)

	// fmt.Println("encodedStr", encodedStr)

	// decodeStr, _ := hex.DecodeString(encodedStr)
	// fmt.Println("decodeStr", decodeStr)
	// decodestrrrr, _ := hex.DecodeString(strrrr)
	// fmt.Println("decodestrrrr", decodestrrrr, " ", string(decodestrrrr))
	// decodestrrrr2, _ := hex.DecodeString(strrrr1)
	// fmt.Println("decodestrrrr2", decodestrrrr2, " ", string(decodestrrrr2))

	// str := "\xa5\x78\xbf\x6e\xb9\x71"
	// s := "\xb6\xd4\xb6\xc0\xc1\xa2\xd1\xa7\xd4\xba\xbf\xc9\xb3\xd6\xd0\xf8\xb7\xa2\xd5\xb9\xce\xca\xcc\xe2"
	// s2 := "\x42\x44\x49\xab\xfc\xbc\xc6"

	// dec := mahonia.NewDecoder("BIG5")
	// fmt.Println(dec.ConvertString(str))
	// fmt.Println(dec.ConvertString(s2))
	// fmt.Println(dec.ConvertString(s))
	// fmt.Println(dec.ConvertString(string(src)))
	// fmt.Println("decodestrrrr", dec.ConvertString(string(decodestrrrr)))
	// fmt.Println("decodestrrrr2", dec.ConvertString(string(decodestrrrr2)))
	// str333 := ""
	// str666 := ""
	// for i := 0; i < len(strrrr); i = i + 2 {
	// 	str333 = "\\x" + strrrr[i:]
	// 	//	fmt.Println("str333", str333)
	// 	str666 = str666 + str333[:4]
	// 	//	fmt.Println("str666", str666)
	// }

	// fmt.Println("str333", str333)
	// fmt.Println("str666", str666)
	// fmt.Println("str666 dec ", dec.ConvertString(str666))
	// fmt.Println("str", str)
	// fmt.Println("str dec ", dec.ConvertString(str))
	// fmt.Println("strrrr", strrrr)
	// fmt.Println("src dec ", dec.ConvertString(string(src)))

	// fmt.Printf("str--- %x", str)
	// fmt.Println("")
	// fmt.Printf("str666--- %x", str666)
	// //	strconv.
	// fmt.Println("")
	// fmt.Printf("strrrr--- %x", strrrr)
	// fmt.Println("")
	// fmt.Println("123", hex.EncodeToString([]byte(strrrr)))
	// // fmt.Println(hex.EncodeToString(str666))
	// fmt.Println("")
	// tttooo := hex.EncodeToString([]byte(str))
	// fmt.Println("tttooo", tttooo)

}
