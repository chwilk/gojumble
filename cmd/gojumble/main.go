package main

import (
  "fmt"
  "io/ioutil"
)

// Pull up dictionary and parse words
func readWords() {
  words, err := ioutil.ReadFile("words")
  if err != nil {
    panic(err)
  }
}

func main() {
  readWords()
}
