package remo_api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func GetResponseBody(url string) ([]byte, error) {
	TOKEN := os.Getenv("TOKEN")

	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+TOKEN)
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("nature remo api error: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	return body, nil
}
