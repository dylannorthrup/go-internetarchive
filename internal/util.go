package internal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/nektro/go-util/ansi/style"
	"github.com/valyala/fastjson"
)

// DieOnError kills the procss if err is not nil
func DieOnError(err error, args ...string) {
	if err != nil {
		LogError(err.Error())
		for _, item := range args {
			LogError(item)
		}
		os.Exit(1)
	}
}

// LogError does that
func LogError(err string) {
	fmt.Print(style.FgRed)
	log.Println(err)
	fmt.Print(style.ResetFgColor)
}

// Assert calls DieOnError is false
func Assert(b bool, msg string) {
	if !b {
		DieOnError(errors.New(msg))
	}
}

// GetBytes fetch urls and return []byte
func GetBytes(urls string, hdrs map[string]string) ([]byte, bool) {
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, false
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, false
	}
	if res.StatusCode >= 400 {
		return nil, false
	}
	byteSlice, err := io.ReadAll(res.Body)
	return byteSlice, err == nil
}

// GetDoc fetch and html document and parses it
func GetDoc(urls string, hdrs map[string]string) (*goquery.Document, []byte, bool) {
	byteSlice, ok := GetBytes(urls, hdrs)
	if !ok {
		return nil, byteSlice, false
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(byteSlice))
	if err != nil {
		return doc, byteSlice, false
	}
	return doc, byteSlice, true
}

// GetJSON is similar to GetDoc but returns a fastjson object
func GetJSON(urls string, hdrs map[string]string) (*fastjson.Value, []byte, bool) {
	byteSlice, ok := GetBytes(urls, hdrs)
	if !ok {
		return nil, byteSlice, false
	}
	val, err := fastjson.ParseBytes(byteSlice)
	return val, byteSlice, err == nil
}
