/*
gojumble: golang implementation of simple de-jumbler program
*/
package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
  "strings"
  "sort"
)

// Function to hash words by alphabetizing their letters
func hash(s string) string {
    a := strings.Split(s, "")
    sort.Strings(a)
    return strings.Join(a, "")
}

// Pull up dictionary and parse words
func readWords() {
  wordList, err := os.Open("words")
  if err != nil {
    log.Fatal(err)
  }
  defer wordList.Close()
  scanner := bufio.NewScanner(wordList)
    for scanner.Scan() {
      myword := strings.Trim(scanner.Text()," ")
      myhash := hash(myword)
      fmt.Println(myhash, myword)
    }

    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }
}

func main() {
  readWords()
}
