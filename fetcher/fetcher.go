package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(u string) ([]byte, error) {
	//<-rateLimiter
	// Use proxy
	//var client *http.Client
	//if count%(len(proxyHttpUrl)+1)==len(proxyHttpUrl) {
	//	client = &http.Client{}
	//}else {
	//	proxyURL, _ := url.Parse(proxyHttpUrl[(count%2)])
	//	trans := &http.Transport{
	//		Proxy: http.ProxyURL(proxyURL),
	//	}
	//	client = &http.Client{Transport: trans}
	//}
	//resp, err := client.Get(u)

	request, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error:status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
