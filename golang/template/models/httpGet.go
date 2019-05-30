package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const httpCleintTimeOut = 5

type GetAPI struct {
	User  string `json:"user"`
	Name  string `json:"name"`
	Email int    `json:"email"`
}
type GetAPIAry []GetAPI

func PrintAPI() {
	url := ""
	apiAry := GetFromAPI(url)
	fmt.Println(apiAry)
}

func GetFromAPI(url string) GetAPIAry {
	body := httpGet(url, "json")
	apiAry := GetAPIAry{}
	err := json.Unmarshal(body, &apiAry)
	if err != nil {
		log.Println("getRainStation Unmarshal err:", err, "body:", string(body))
		return apiAry
	}
	log.Printf("getRainStation len:%v", len(apiAry))
	return apiAry
}

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
