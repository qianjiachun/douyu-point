package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"unsafe"
)

func HttpGet(url string) string {
	//超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	CheckErr(err)
	response, err := client.Do(req)
	CheckErr(err)
	defer func() {
		_ = response.Body.Close()
	}()
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	return string(bytes)
}

func HttpPost(url string, data string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	CheckErrNoExit(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(req)
	CheckErrNoExit(err)
	defer func() {
		_ = response.Body.Close()
	}()
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErrNoExit(err)
	return string(bytes)
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

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func GetFieldValue(data string, field string) string {
	return GetStrMiddle(data, field+"@=", "/")
}

/*
	CheckErr
*/
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func CheckErrNoExit(err error) {
	if err != nil {
		log.Println(err)
	}
}
func CheckErrRollback(err error, tx *sql.Tx) bool {
	if err != nil {
		//log.Println(err)
		err = tx.Rollback()
		if err != nil {
			log.Println("tx.Rollback() Error:" + err.Error())
			return false
		}
		return false
	}
	return true
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func JsonToMap(jsonStr string) (map[string]int, error) {
	m := make(map[string]int)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	return m, nil
}
