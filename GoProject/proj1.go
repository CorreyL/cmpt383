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
)

// os.Args[0] is the name of the program
// os.Args[1] is the first argument passed into the program
	// os.Args[1] will be the directory of that the program will work in

func errorCheck(e error){
	if(e != nil){
		panic(e)
	}
}

func visit(path string, f os.FileInfo, err error) error {
  fmt.Printf("Visited: %s\n", path)
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
	os.Stdout.WriteString(hex.EncodeToString(hasher.Sum(nil)))
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

/*
f, err := os.Open(os.Args[1])
if err != nil {
    log.Fatal(err)
}
defer f.Close()
if _, err := io.Copy(hasher, f); err != nil {
    log.Fatal(err)
}
*/ // Supposedly handles bigger files better?

