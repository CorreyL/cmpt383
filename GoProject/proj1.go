// package proj1
package main

import (
	"fmt"
	"crypto/sha256"
	"os"
	// "io/ioutil"
	"io"
	"encoding/hex"
	"path/filepath"
	"regexp"
	"strings"
)

// os.Args[0] is the name of the program
// os.Args[1] is the first argument passed into the program
	// os.Args[1] will be the directory of that the program will work in
	
// Use map to store current directory of files, then use a another data structure to store directories?

var currentDirFiles map[string]string = make(map[string]string)  // A map of the current directory's files, wherein keys are stored as Filename : SHA256-Code

func errorCheck(e error){
	if(e != nil){
		panic(e)
	}
}

func visit(path string, f os.FileInfo, err error) error {
	/*
	hasher := sha256.New()
	s, err2 := os.Open(path)
	errorCheck(err2)
	if _, err2 := io.Copy(hasher, s); err2 != nil{
		panic(err2)
	}
	currentDirFiles[path] = hex.EncodeToString(hasher.Sum(nil))
	*/
	r, _ := regexp.Compile(`\.[a-zA-Z0-9]+`) // Regular Expression for the string of a file
	fmt.Println("Testing: ", r.MatchString(path))
	if(r.MatchString(path)){
		fmt.Printf("Visited: %s\n", path)
		hasher := sha256.New()
		s, err2 := os.Open(path)
		errorCheck(err2)
		if _, err2 := io.Copy(hasher, s); err2 != nil{
			panic(err2)
		}
		split := strings.Split(path, "/")
		newPath := strings.TrimSpace(split[len(split)-1])
		fmt.Println("The newPath is: ", newPath)
		currentDirFiles[newPath] = hex.EncodeToString(hasher.Sum(nil))
	}
  return nil
}

func main(){
	hasher := sha256.New()
	root := os.Args[1]
	err := filepath.Walk(root, visit)
	s, err := os.Open("asn1.txt")
	errorCheck(err)
	defer s.Close()
	if _, err := io.Copy(hasher, s); err != nil{
		panic(err)
	}
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	fmt.Println("filepath.Walk() returned %v\n", err)
	fmt.Println(currentDirFiles)
}

