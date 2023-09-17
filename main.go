package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

func main() {
	url := "https://www.thepaper.cn"
	body, err := Fetch(url)

	if err != nil {
		fmt.Printf("fetch url error:%v\n", err)
		return
	}

	doc, err := htmlquery.Parse(bytes.NewReader(body))
	nodes := htmlquery.Find(doc, `/html/body/div/main/div[2]/div[3]/div[2]/div[1]/div[3]/div[1]/div/div[2]/div/div[2]/div/div/div/div[1]/a/h2`)
	for _, node := range nodes {
		fmt.Println("fetch card ", node.FirstChild.Data)
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
