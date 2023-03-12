package main

import (
    "net/http"
)

func HealthHandler(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("OK"))
}

func main() {
  http.HandleFunc("/health", HealthHandler)

  http.ListenAndServe(":8080", nil)
}
