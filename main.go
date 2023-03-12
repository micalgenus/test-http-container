package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
)

type response struct {
  URL     string      "json:\"url\""
  Method  string      "json:\"method\""
  Headers http.Header "json:\"headers\""
  Body    []byte      "json:\"body\""
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("OK"))
}

func notFoundHandler(w http.ResponseWriter, req *http.Request) {
  w.WriteHeader(404)
}

func errorHandler(w http.ResponseWriter, req *http.Request) {
  w.WriteHeader(500)
}

func echoHandler(w http.ResponseWriter, req *http.Request) {
  var err error
  res := &request{}
  res.Method = req.Method
  res.Headers = req.Header
  res.URL = req.URL.String()
  res.Body, err = ioutil.ReadAll(req.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  response, err := json.MarshalIndent(res, "", "  ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func main() {
  http.HandleFunc("/health", healthHandler)
  http.HandleFunc("/notfound", notFoundHandler)
  http.HandleFunc("/error", errorHandler)
  http.HandleFunc("/", echoHandler)

  http.ListenAndServe(":8080", nil)
}
