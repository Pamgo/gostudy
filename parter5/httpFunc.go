package main

import "net/http"

type Handler interface {
	ServerHTTP(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter,*http.Request)

func (f HandlerFunc) SereverHTTP(w http.ResponseWriter, r *http.Request)  {
	f(w,r)
}

func HandleFunc(pattern string, handler func(http.ResponseWriter,*http.Request))  {
	http.DefaultServeMux.HandleFunc(pattern,handler)
}
