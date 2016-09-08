package main
import (
	"fmt"
	"io/ioutil"
)

func isPrime(x int) bool{
	for a:=x-1; a>1; a--{
		if(x%a == 0){
			return false
		}
	}
	return true
}

func countPrimes(x int) int{
	var total = 0
	if(x<0){
		return total
	}
	for a:=x-1;a>1;a--{
		if( isPrime(a) ){
			total++
		}
	}
	return total
}

func countStrings(filename string) map[string]int{
	var m map[string]int
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	var current string
	current = ""
	for a:=0;a<len(str);a++{
		if(string(str[a]) != string(" ") || string(str[a]) != string("\n")){
			current = current + string(str[a])
			fmt.Println(current)
		} else {
			fmt.Println(current)
			if(m[current] == 0){
				m[current] = 1
			} else{
				m[current] = m[current] + 1
			}
			current = ""
		}
	}
	return m
}

func main() {
	/*
	fmt.Println("Input a number for countPrimes()")
	var user_input int 
	fmt.Scan(&user_input)
  fmt.Println(countPrimes(user_input)) 
	*/
	fmt.Printf("{")
	for x, y := range countStrings("textfile.txt") { 
		fmt.Printf(x, ":",y)
	}
	fmt.Printf("}")
}