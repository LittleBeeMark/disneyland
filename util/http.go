package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(urlstring string, params map[string]string) ([]byte, error) {
	u, _ := url.Parse(urlstring)
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	//fmt.Println("request url :", u.String())
	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return result, err
}
