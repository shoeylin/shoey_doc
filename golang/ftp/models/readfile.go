package models

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

type MappingData struct {
	Str0 string
	Str1 string
	Str2 string
	Str3 string
	Str4 string
	Str5 string
	Str6 string
	Str7 string
}

func Opnefile(readFile string) []MappingData {
	// 開檔
	inputFile, Error := os.Open(readFile)
	// 判斷是否開檔錯誤
	if Error != nil {
		fmt.Println("開檔錯誤!")
		return nil
	}
	// 離開時自動執行關檔
	defer inputFile.Close()
	// 宣告行計數器
	cnt := 1
	//
	inputReader := bufio.NewReader(inputFile)
	//^^^^^^^^^^^         ^^^^^^^   ^^^^^^^
	// 緩衝輸入物件        建立函數   來源:已開啟檔案

	// 用迴圈讀取檔案內容
	var mappingDataAry []MappingData
	var md MappingData
	// mapMappingData := make(map[string]MappingData)

	for {
		// 讀取字串直到遇到跳行符號
		inputString, Error := inputReader.ReadString('\n')

		//big5 to UTF-8
		tempBig5, err := Decodebig5([]byte(inputString))
		CheckError(err)
		inputString = string(tempBig5)

		// 若到檔尾時分發生  io.EOF 錯誤
		// 根據此錯誤 判斷是否離開
		if Error == io.EOF {
			fmt.Println("已讀取到檔尾!!")
			return mappingDataAry
		}

		md = txtSplit(inputString)

		// 列印內容
		// fmt.Printf("第%2d行:%s\n", cnt, inputString)
		cnt++

		// fmt.Printf("第%2d行:%s\n", cnt, p)
		mappingDataAry = append(mappingDataAry, md)
	}
}

func txtSplit(inputstring string) MappingData {

	delim := ";"

	msgType := strings.Split(inputstring, delim)

	var md MappingData
	for i, v := range msgType {

		//2 4 5
		switch i {
		case 0:
			md.Str0 = strings.TrimSpace(v)
		case 1:
			fmt.Println("第", i, "個:", v)
			md.Str1 = strings.TrimSpace(v)
		case 2:
			fmt.Println("第", i, "個:", v)
			md.Str2 = strings.TrimSpace(v)
		case 3:
			fmt.Println("第", i, "個:", v)
			md.Str3 = strings.TrimSpace(v)
		case 4:
			fmt.Println("第", i, "個:", v)
			md.Str4 = strings.TrimSpace(v)
		case 5:
			fmt.Println("第", i, "個:", v)
			md.Str5 = strings.TrimSpace(v)
		case 6:
			fmt.Println("第", i, "個:", v)
			md.Str6 = strings.TrimSpace(v)
		case 7:
			fmt.Println("第", i, "個:", v)
			md.Str7 = strings.TrimSpace(v)
		default:
			continue
		}
	}

	return md
}

//convert BIG5 to UTF-8
func Decodebig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
