package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

var fileName string = "/tmp/dump.html"
var sites = []string {
  "www.google.com",
  "www.safetychanger.com"}


func main() {
  fmt.Println("Start processing...")

  if len(os.Args) > 1 {
    // get filename from 1st arg
    fileName = os.Args[1]

    // sites from args
    if len(os.Args) > 2 {
      sites = os.Args[2:]
    }
  }

  for i := 0; i < len(sites); i++ {
    write(getPage(sites[i]))
  }

  defer fmt.Println("Ready")
}

func write(content []byte) {
  var f *os.File
  var err error

  if _, err := os.Stat(fileName); err == nil {
    f, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
  } else {
    f, err = os.Create(fileName)
  }
  check(err)

  defer f.Close()

  no, err := f.Write(content)
  fmt.Printf("wrote %d bytes\n", no)
  check(err)
  f.Sync()
}

func getPage(url string) []byte {
  if strings.Contains(url, "http") == false {
    url = "http://" + url
  }

  resp, err := http.Get(url)
  check(err)
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  check(err)

  fmt.Println("read " + url)
  return body
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
