package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type hijackHandler struct{}

func (hijackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	time.Sleep(1 * time.Second)
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("============")
	// Don't forget to close the connection:
	defer conn.Close()
	bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
	bufrw.Flush()
	s, err := bufrw.ReadString('\n')
	if err != nil {
		log.Printf("error reading string: %v", err)
		return
	}
	fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
	bufrw.Flush()
}

type barHandler struct{}

func (barHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// w.Write([]byte("hello world"))
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/hijack", hijackHandler{})
	mux.Handle("/bar", barHandler{})

	s := &http.Server{
		Addr:           ":8889",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
