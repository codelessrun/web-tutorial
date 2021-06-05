package main

import "net/http"

type MyHanlder struct{}

func (m *MyHanlder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("HELLO WORLD"))
}

type Abouthandler struct{}

func (m *Abouthandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("About"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func main() {

	mh := MyHanlder{}
	about := Abouthandler{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, // defaultServeMux
	}

	http.Handle("/hello", &mh) // 注册到 defaultServeMux
	http.Handle("/about", &about)
	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Home"))
	})
	// http.HandleFunc("/welcome", welcome)
	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Write([]byte("Hello World"))
	// })

	// http.ListenAndServe("localhost:8080", nil)
}
