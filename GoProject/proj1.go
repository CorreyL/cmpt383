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
	// "strings"
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

/*
// May need this later; it pulls just the filename out
split := strings.Split(path, "/")
newPath := strings.TrimSpace(split[len(split)-1])
fmt.Println("The newPath is: ", newPath)
*/

func visit(path string, f os.FileInfo, err error) error {
	r, _ := regexp.Compile(`\.[a-zA-Z0-9]+`) // Regular Expression for the string of a file
	fmt.Println("Testing: ", path, " ", r.MatchString(path))
	if(r.MatchString(path)){
		fmt.Printf("Visited: %s\n", path)
		hasher := sha256.New()
		s, err2 := os.Open(path)
		errorCheck(err2)
		if _, err2 := io.Copy(hasher, s); err2 != nil{
			panic(err2)
		}
		if _, exist := currentDirFiles[hex.EncodeToString(hasher.Sum(nil))]; exist{
			currentDirFiles[hex.EncodeToString(hasher.Sum(nil))] = currentDirFiles[hex.EncodeToString(hasher.Sum(nil))]+"|"+path
		} else{
			currentDirFiles[hex.EncodeToString(hasher.Sum(nil))] = path
		}
	}
  return nil
}

func main(){
	// hasher := sha256.New()
	root := os.Args[1]
	err := filepath.Walk(root, visit)
	errorCheck(err)
	fmt.Println("filepath.Walk() returned %v\n", err)
	fmt.Println(currentDirFiles)
	// err = os.Remove("testDirectory/Asn 1.txt")
}

