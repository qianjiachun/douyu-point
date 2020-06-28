package common

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"unsafe"
)

func HttpGet(url string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}

func GetStrMiddle(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start)
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func GetFieldValue(data string, field string) string {
	return GetStrMiddle(data, field + "@=", "/")
}