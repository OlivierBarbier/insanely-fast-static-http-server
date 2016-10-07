package main

import (
    "os"
    "net/http"
)

func main() {
  path := os.Getenv("STATIC_CONTENT_PATH")
  http.Handle("/", http.FileServer(http.Dir(path)))
  http.ListenAndServe(":3000", nil)
}
