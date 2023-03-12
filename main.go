package main

import (
    "net/http"
)

func main() {
  http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("OK"))
  })

  http.ListenAndServe(":8080", nil)
}
