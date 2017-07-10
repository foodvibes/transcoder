package main

import (
  "testing"
  "bytes"
  "path/filepath"
  "io/ioutil"
)

func TestMain(t *testing.T) {
  var (
    testPath = "test"
    contentPath = filepath.Join(testPath, "recipe.txt")
    templatePath = filepath.Join(testPath, "template.txt")
    expectedPath = filepath.Join(testPath, "expected.txt")
  )

  b := new(bytes.Buffer)
  Transcode(contentPath, templatePath, b, []string{filepath.Base(contentPath)})
  result := b.String()
  e, _ := ioutil.ReadFile(expectedPath)
  var expected = string(e)

  if result != expected {
    t.Error(
      "expected", expected,
      "result", result,
    )
  }
}
