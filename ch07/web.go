package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	l := LogAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.HandleGreeting)
	http.ListenAndServe(":8082", nil)
}

func LogOutput(message string) {
	fmt.Println(message)
}

type Logger interface {
	Log(message string)
}

// LogAdapter 定义 函数即值（类似函数指针）
type LogAdapter func(message string)

func (log LogAdapter) Log(message string) {
	log(message)
}

type SimpleDataStore struct {
	useData map[string]string
}

type DataStore interface {
	UserNameForId(userId string) (string, bool)
}

func (ds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	name, ok := ds.useData[userId]
	return name, ok
}

// NewSimpleDataStore 定义工厂函数
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		useData: map[string]string{"1": "Fred", "2": "Mary", "3": "pat"},
	}
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("in SayHello for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userId string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

// NewSimpleLogic 业务逻辑工厂类
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type Logic interface {
	SayHello(userId string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) HandleGreeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userId := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}
