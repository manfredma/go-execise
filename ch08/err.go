package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime"
)

func main() {
	// 哨兵错误
	SentinelErr()
	// 自定义错误
	customException()
	// 包装错误
	wrapException()
	// Error AS
	errorAs()
	// panic and recover
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}

func customException() {
	data, err := LoginAndGetData("1", "2", "3")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i, datum := range data {
		fmt.Println(i, datum)
	}
}

func wrapException() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}

func SentinelErr() {
	data := []byte("this is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{Status: NotFound, Message: fmt.Sprintf("file %s not found", file)}
	}
	return data, nil
}

func login(uid, pwd string) error {
	if rand.Intn(10) > 5 {
		return nil
	}
	return errors.New("error")
}

func getData(file string) ([]byte, error) {
	if rand.Intn(10) > 5 {
		return []byte{1, 2, 3}, nil
	}
	return nil, errors.New("error")
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func errorAs() {
	err := mockError()
	var myError MyError
	if errors.As(err, &myError) {
		fmt.Println(err.Error(), myError.Code)
	}
}

type MyError struct {
	Message string
	Code    string
}

func (m MyError) Error() string {
	return m.Message
}

func mockError() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("a new error [%w]", err)
		}
	}()
	return MyError{
		Code:    "102",
		Message: "hello, MyError",
	}
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Println(string(buf))
		}
	}()
	fmt.Println(60 / i)
}
