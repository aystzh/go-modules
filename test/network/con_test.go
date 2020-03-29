package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handlerError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestCon(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handlerError(t, err)
	defer ln.Close()
	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handlerError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handlerError(t, err)
	if string(body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}

//推荐写法
func TestCon2(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	body, _ := ioutil.ReadAll(w.Result().Body)
	if string(body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}
