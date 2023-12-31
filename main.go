package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
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

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println("htmlquery.Parse failed:%v", err)
	}

	doc.Find("div a[target=_blank] h2").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)

	})
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
