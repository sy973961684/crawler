package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"regexp"
)

var headerRe = regexp.MustCompile(`[\s\S]*?img alt="([\s\S]*?)"[\s\S]*?`)

func main() {
	url := "https://www.thepaper.cn"
	body, err := Fetch(url)

	if err != nil {
		fmt.Printf("fetch url error:%v\n", err)
		return
	}

	mathers := headerRe.FindAllSubmatch(body, -1)
	for _, m := range mathers {
		fmt.Println("fetch card news:", string(m[1]))
	}

}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeteminEncoding(bodyReader)
	uft8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return io.ReadAll(uft8Reader)
}

func DeteminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1023)

	if err != nil {
		fmt.Printf("fetch error:%v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
