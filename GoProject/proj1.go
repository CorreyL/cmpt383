package main

import (
	"fmt"
	"crypto/sha256"
	"os"
	"io/ioutil"
	"encoding/hex"
	"path/filepath"
	// "regexp"
	"strings"
)

var currentDirFiles map[string]string = make(map[string]string)  // A map of the current directory's files, wherein entries are stored as SHA256-Code : Filename|Filename|...

func errorCheck(e error){
	if(e != nil){
		fmt.Println("ERROR: The directory inputted does not exist.")
		fmt.Println("Aborting program.")
		os.Exit(2)
	}
}

func visit(path string, fi os.FileInfo, err error) error {
	checkDir, retErr := os.Stat(path)
	errorCheck(retErr)
	if( !checkDir.IsDir() ){
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
		os.Exit(2)
	}
	err := filepath.Walk(os.Args[1], visit)
	errorCheck(err)
	if( len(currentDirFiles) == 0 ){ // Checks to see if the directory had any files
		fmt.Println("The directory provided contains no files.")
		os.Exit(0)
	}
	noDups := true // If noDups is not changed, then the directory and all sub-directories have no duplicate files
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

