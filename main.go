package main

import (

"fmt"
"net/http"
"log"
)


func HandleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}


func main () {


http.HandleFunc("/", HandleRequest)
err := http.ListenAndServe(":8081", nil)
    if err != nil {
        log.Fatal("ListenError:", err)
    }

}

