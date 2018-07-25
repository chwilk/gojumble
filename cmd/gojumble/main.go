package main

import (
  "fmt"
  "os"
)

// Pull up dictionary and parse words
func readWords() {
  words, err := os.Open("words")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer words.Close()
}

func main() {
  readWords()
}
