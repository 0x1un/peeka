package component

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func Post(url string, reader *bytes.Reader) []byte {
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {

	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	return respBytes
}
