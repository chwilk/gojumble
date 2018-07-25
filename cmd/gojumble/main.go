package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
)

// Pull up dictionary and parse words
func readWords() {
  wordList, err := os.Open("words")
  if err != nil {
    log.Fatal(err)
  }
  defer wordList.Close()
  
  scanner := bufio.NewScanner(wordList)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
  
}

func main() {
  readWords()
}
