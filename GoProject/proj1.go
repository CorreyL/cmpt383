// package proj1
package main

import (
	// "fmt"
	"crypto/sha256"
	"os"
	"io/ioutil"
	"encoding/hex"
)

// os.Args[0] is the name of the program
// os.Args[1] is the first argument passed into the program
	// os.Args[1] will be the directory of that the program will work in

func errorCheck(e error){
	if(e != nil){
		panic(e)
	}
}

func main(){
	hasher := sha256.New()
	s, err := ioutil.ReadFile(os.Args[1])
	hasher.Write(s)
	errorCheck(err);
	 os.Stdout.WriteString(hex.EncodeToString(hasher.Sum(nil)))
	// fmt.Println(os.Args[1])
}



