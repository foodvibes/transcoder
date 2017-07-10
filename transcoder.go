package main

import (
  "os"
  "log"
  "errors"
  "io"
  "io/ioutil"
  "strings"
  "text/template"
)

type Recipe struct {
  Title string
  Ingredients []string
  Yield string
  Directions []string
}

type TemplateContext struct {
  Recipe Recipe
  Args []string
}

func main() {
  if (len(os.Args) < 2) {
    log.Fatal("No file specified.")
  }

  if (len(os.Args) < 3) {
    log.Fatal("No template specified.")
  }

  err := Transcode(os.Args[1], os.Args[2], os.Stdout, os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

func Transcode(contentPath string, templatePath string, output io.Writer, extra []string) error {
  // Read file
  content, err := ioutil.ReadFile(contentPath)
  if err != nil {
    return err
  }

  // Split content by sequential newlines
  sections := strings.Split(string(content), "\n\n")
  if (len(sections) < 4) {
    return errors.New("There should be at least 4 sections: title, ingredients, yield and directions.")
  }

  // Second section is list of ingredients
  ingredients := strings.Split(sections[1], "\n")

  // First section is title; third is yield; fourth and beyond are directions
  r := Recipe{sections[0], ingredients, sections[2], sections[3:]}

  // Load the template
  t, err := template.ParseFiles(templatePath)
  if err != nil {
    return err
  }

  context := TemplateContext{r, extra}

  // Apply context to template and write to output
  err = t.Execute(output, context)
  if err != nil {
    return err
  }

  return nil
}
