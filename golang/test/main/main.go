package main

import (
	// "encoding/hex"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	// "github.com/axgle/mahonia"
)

func main() {

	t := time.Now().Format("20060102")
	fmt.Println("123")

	fmt.Println(t)

	now := time.Now()
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1.Format("20060102"))

	//E:\\mygo\\src\\shoey_doc\\golang\\test\\main\\
	// //E:\\mygo\\src\\shoey_doc\\golang\\test\\main\\
	// err := DeCompress("./20190121.zip", "./1", d1.Format("20060102"))
	// fmt.Println("err", err)

	// r, err := zip.OpenReader("./20190121.zip")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer r.Close()

	// // Iterate through the files in the archive,
	// // printing some of their contents.
	// for _, f := range r.File {
	// 	fmt.Printf("Contents of %s:\n", f.Name)
	// 	rc, err := f.Open()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	_, err = io.CopyN(os.Stdout, rc, 68)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	rc.Close()
	// 	fmt.Println("123")
	// }
	filename := d1.Format("20060102") + ".zip"

	DeCompress(filename)

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

//解压
func DeCompress2(zipFile, dest, name string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "./"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func DeCompress(filename string) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.MkdirAll(filename[:len(filename)-4], 0644)
	fmt.Println("error: ", err)
	err = os.Chdir(filename[:len(filename)-4])
	fmt.Println("error2: ", err)
	for _, k := range r.Reader.File {
		if k.FileInfo().IsDir() {
			err := os.MkdirAll(k.Name, 0644)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		r, err := k.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("unzip: ", k.Name)
		defer r.Close()
		NewFile, err := os.Create(k.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.Copy(NewFile, r)
		NewFile.Close()
	}
}
