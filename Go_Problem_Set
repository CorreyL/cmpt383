package main
import (
	"fmt"
	"strings"
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
	m := make(map[string]int)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	var words []string
	words = strings.Fields(str)
	for _,element := range words{
		if(m[element] == 0){
			m[element] = 1
		} else{
			m[element] = m[element] + 1
		}
	}
	return m
}

func main() {
	// Question 1
	/*
	fmt.Println("Input a number for countPrimes()")
	var user_input int 
	fmt.Scan(&user_input)
  fmt.Println(countPrimes(user_input)) 
	*/
	// Question 2
	fmt.Printf("{")
	m := countStrings("textfile.txt")
	len_check := 0 // Ensures that the string ", " is not printed for the final pair
	for x, y := range m { 
		fmt.Print("\"",x,"\"",":", y)
		if(len_check != len(m)-1){
			fmt.Print(", ")
		}
		len_check++
	}
	fmt.Printf("}")
}