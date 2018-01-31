package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "hello")
     err := r.ParseForm()
    if err != nil {
        // Handle error here via logging and then return    
        log.Println(" Error while parsing form")        
    }

    name := r.PostFormValue("name")
    fmt.Fprintf(w, "Hello, %s!", name)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    counter++
    fmt.Fprintf(w, strconv.Itoa(counter))
    mutex.Unlock()
}

func main() {
     http.HandleFunc("/dupa", echoString)
   /* http.HandleFunc("/increment", incrementCounter)
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })*/
     fs := http.FileServer(http.Dir("static"))
        http.Handle("/", fs)

        log.Println("Listening...")
        http.ListenAndServe(":3000", nil)

    log.Fatal(http.ListenAndServe(":8081", nil))
}