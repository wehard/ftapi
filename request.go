package ftapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func DoFTRequest(endpoint string, accessToken string) []byte {
	var data struct {
		Token string `json:"access_token"`
	}
	data.Token = accessToken
	b, err := json.Marshal(data)
	req, err := http.NewRequest("GET", "https://api.intra.42.fr"+endpoint, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ft request:", err)
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Println("ft request error: too many requests!")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	time.Sleep(500 * time.Millisecond)
	return (body)
}
