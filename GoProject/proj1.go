// package proj1
package main

import (
	"fmt"
	"crypto/sha256"
	"os"
	"io/ioutil"
	//"io"
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

/*
// May need this later; it pulls just the filename out
split := strings.Split(path, "/")
newPath := strings.TrimSpace(split[len(split)-1])
fmt.Println("The newPath is: ", newPath)
*/

func visit(path string, fi os.FileInfo, err error) error {
	r, _ := regexp.Compile(`\.[a-zA-Z0-9]+`) // Regular Expression for the string of a file
	if(r.MatchString(path)){
		hasher := sha256.New()
		s, err := ioutil.ReadFile(path)
		hasher.Write(s)
		errorCheck(err);
		if _, exist := currentDirFiles[hex.EncodeToString(hasher.Sum(nil))]; exist{
			currentDirFiles[hex.EncodeToString(hasher.Sum(nil))] = currentDirFiles[hex.EncodeToString(hasher.Sum(nil))]+"|"+path
		} else{
			currentDirFiles[hex.EncodeToString(hasher.Sum(nil))] = path
		}
	}
  return nil
}

func main(){
	if( len(os.Args) < 2 ){ // If a directory was provided, then len(os.Args) >= 2 
		fmt.Println("ERROR: A directory was not supplied to proj1.go.")
		fmt.Println("Aborting program.")
		os.Exit(1)
	}
	err := filepath.Walk(os.Args[1], visit)
	errorCheck(err)
	if( len(currentDirFiles) == 0 ){ // Checks to see if anything was added to the map
		fmt.Println("The directory provided either does not exist OR contains no files.")
		os.Exit(0)
	}
	noDups := true;
	for x, y := range currentDirFiles {
		if( strings.IndexAny(y,"|") != -1 ){
			noDups = false;
			split := strings.Split(y, "|")
			fmt.Println("SHA-256 KEY:", x)
			fmt.Println("Number of files with this signature:", len(split) )
			for i := range split{
				fmt.Println("	", split[i])
			}
			fmt.Printf("\n\n")
		}
	}
	if( noDups ){
		fmt.Println("No duplicate files were found!")
	}
}

