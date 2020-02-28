package net

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Post(url string, reader *bytes.Reader) []byte {
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	return respBytes
}

type contextKey int

// type ProxyRequest struct {
// 	Purl   string
// 	Turl   string
// 	Method string
// 	Data   *io.Reader
// 	Client *http.Client
// }

const (
	proxyURLKey contextKey = 0
)

func proxy(req *http.Request) (*url.URL, error) {
	iproxy := req.Context().Value(proxyURLKey)
	if iproxy == nil {
		fmt.Printf("no proxy found\n")
		return nil, nil
	}
	proxyURL := iproxy.(*url.URL)
	fmt.Printf("proxy found:%s\n", proxyURL.String())
	return proxyURL, nil
}

func RequestWithProxy(method, purl, turl string) (*http.Response, error) {
	transport := &http.Transport{
		Proxy: proxy,
	}
	client := http.Client{
		Transport: transport,
	}
	{
		proxyURL, _ := url.Parse(purl)
		ctx := context.WithValue(context.Background(), proxyURLKey, proxyURL)
		requestWithProxy, err := http.NewRequestWithContext(ctx, method, turl, nil)
		if err != nil {
			return nil, err
		}
		response, _ := client.Do(requestWithProxy)
		return response, nil
	}
}
