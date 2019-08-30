package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	//get
	ary := GetAPIFunction()
	fmt.Println(ary)

	//post
	res := SetAPIFunction("abcd", "id", "name", "email")
	fmt.Println(res)
}

type GetAPI struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SetAPI struct {
	Token string `json:"token"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAPIFunction() []GetAPI {
	url := "http://192.168.0.1:12345/api/getAPI"
	body := httpGet(url, "json")
	getAPIAry := []GetAPI{}
	err := json.Unmarshal(body, &getAPIAry)
	if err != nil {
		log.Println("getAPI Unmarshal err:", err, "body:", string(body))
		return getAPIAry
	}
	log.Printf("getAPI len:%v", len(getAPIAry))
	return getAPIAry
}

func SetAPIFunction(token, id, name, email string) string {
	url := "http://192.168.0.1:12345/api/postAPI"
	setAPI := SetAPI{}
	setAPI.Token = token
	setAPI.ID = id
	setAPI.Name = name
	setAPI.Email = email

	b, _ := json.Marshal(&setAPI)
	fmt.Println("b", string(b))

	res := httpPOST(url, "json", string(b))

	// fmt.Println("res", string(res))

	return string(res)
}

const httpCleintTimeOut = 5

func httpGet(url string, acceptType string) []byte {

	client := &http.Client{Timeout: time.Duration(httpCleintTimeOut * time.Second)}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("httpGet err: ", err, "url:", url)
		return nil
	}

	req.Header.Set("Accept", "application/"+acceptType)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("httpGet Do err: ", err, "url:", url)
		return nil
	}

	if resp != nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func httpPOST(url string, acceptType string, data string) []byte {

	client := &http.Client{Timeout: time.Duration(httpCleintTimeOut * time.Second)}

	// data := url.Values{"start": {"0"}, "offset": {"xxxx"}}
	body := strings.NewReader(data)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Println("httpGet err: ", err, "url:", url)
		return nil
	}

	req.Header.Set("Accept", "application/"+acceptType)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("httpPOST Do err: ", err, "url:", url)
		return nil
	}

	if resp != nil {
		defer resp.Body.Close()
	}
	bodyResp, _ := ioutil.ReadAll(resp.Body)
	return bodyResp
}
