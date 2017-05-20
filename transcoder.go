package main

import (
  "os"
  "log"
  "io/ioutil"
  "strings"
  "html/template"
)

type Recipe struct {
  Title string
  Ingredients []string
  Yield string
  Directions []string
}

func main() {
  if (len(os.Args) < 2) {
    log.Fatal("No file specified.")
  }

  if (len(os.Args) < 3) {
    log.Fatal("No template specified.")
  }

  // Read file
  content, err := ioutil.ReadFile(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }

  // Split content by sequential newlines
  sections := strings.Split(string(content), "\n\n")
  if (len(sections) < 4) {
    log.Fatal("There should be at least 4 sections: title, ingredients, yield and directions.")
  }

  ingredients := strings.Split(sections[1], "\n")

  r := Recipe{sections[0], ingredients, sections[2], sections[3:]}

  t, err := template.ParseFiles(os.Args[2])
  if err != nil {
    log.Fatal(err)
  }

  err = t.Execute(os.Stdout, r)
  if err != nil {
    log.Fatal(err)
  }
}
