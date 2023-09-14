package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("fetch url error:%v\n", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error close io:%v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("read content failed:%v", err)
		return
	}

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links!\n", numLinks)

	numLinks = bytes.Count(body, []byte("<a"))
	fmt.Printf("homepage has %d links!\n", numLinks)

	exist := strings.Contains(string(body), "疫情")
	fmt.Printf("是否存在疫情:%v\n", exist)

	exist = bytes.Contains(body, []byte("疫情"))
	fmt.Printf("是否存在疫情:%v\n", exist)
}
