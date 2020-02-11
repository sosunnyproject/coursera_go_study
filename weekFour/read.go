package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "unicode/utf8"
)

type Names struct {
	fName string
	lName  string
}

func main() {
  // read file
  // represents in a slice of structs
  var filename string
  fmt.Printf("Please enter the name of file to test: ")
	fmt.Scan(&filename)
  file, err := os.Open(filename)
  if err != nil {
	  log.Fatal(err)
  } 
  defer file.Close()

  // text file: "first" + " " + "last"
  // read each line --> create struct of each name
  scanner := bufio.NewScanner(file)
  sliceNames := make([]Names, 0)

  for scanner.Scan() {
	  var name []string = strings.Split(scanner.Text(), " ")
	  var fname string = name[0]
	  var lname string = name[1]
	//   fmt.Println(fname)
	//   fmt.Println(lname)

  // add each struct to a slice
  //after all lines, slice containing one struct for each line of file
	  person := Names{fName: fname, lName: lname}
	//   fmt.Println(person)
	  sliceNames = append(sliceNames, person)
  }

  if err := scanner.Err(); err != nil {
	  log.Fatal(err)
  }

  // iterate through your slice of structs, print first and last names
  fmt.Println(sliceNames)
}