package main
import (
	"fmt"
	"strings"
	"errors"
	"io/ioutil"
	"strconv"
	"math"
)

func isPrime(x int) bool{ // Used in countPrimes; Returns true if the given number is a Prime Number
	for a:=x-1; a>1; a--{
		if(x%a == 0){
			return false
		}
	}
	return true
}

func countPrimes(x int) int{ // Counts the number of Prime Numbers from all values 1:x (not-inclusive)
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

func countStrings(filename string) map[string]int{ // Returns a map where each pair is a word and the number of time it occurs in the text file
	m := make(map[string]int)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	var words []string
	words = strings.Fields(str) // Store the contents of the file in a slice
	for _,element := range words{
		if(m[element] == 0){
			m[element] = 1
		} else{
			m[element] = m[element] + 1
		}
	}
	return m
}

type Time24 struct {
    hour, minute, second uint8
}
// 0 <= hour < 24
// 0 <= minute < 60
// 0 <= second < 60

func validTime24(time Time24) bool{
	if( (0 <= time.hour && time.hour < 24) && (0 <= time.minute && time.minute < 60) && (0 <= time.second && time.second < 60) ){
		return true
	}
	return false
}

func equalsTime24(a Time24, b Time24) bool{
	if( (a.hour == b.hour) && (a.minute == b.minute) && (a.second == b.second) ){
		return true
	}
	return false
}

func lessThanTime24(a Time24, b Time24) bool{
	if( (a.hour <= b.hour) ){
		if( a.hour != b.hour ){
			return true
		} else {
			if ( a.minute <= b.minute ){
				if( a.minute != b.minute ){
					return true
				}	else {
					if( a.second < b.second ){
						return true
					}
				}
			}
		}
	}
	return false
}

func minTime24(times []Time24) (Time24, error){
	var noError error
	noError = errors.New("nil")
	if(len(times) == 0){ // If times has no arguments
		var zeroTime Time24
		var err error
		
		zeroTime.hour = 0
		zeroTime.minute = 0
		zeroTime.second = 0
		
		err = errors.New("The argument passed in is empty.")
		
		return zeroTime, err
	} else if(len(times) == 1){ // If times has only one argument, then it is the smallest time
		return times[0], noError
	}
	track := 0 // Tracks which of the arguments has the smallest time thus far
	for x := 1; x<len(times);x++{
		if( !lessThanTime24( times[track], times[x] ) ){ // Track only changes if the current tracked time is not strictly less than the one being compared to
			track = x
		}
	}
	return times[track], noError
}

func addZero(t uint8) string{ // Ensures that the method String() for Time24 adds a 0 leading to time values < 10
	if(int(t) < 10){
		return "0" + strconv.Itoa(int(t))
	}
	return strconv.Itoa(int(t))
}

func (t Time24) String() string{
	return addZero(t.hour) + ":" + addZero(t.minute) + ":" + addZero(t.second)
}

func (t *Time24) AddOneHour(){
	if(t.hour == 23){
		t.hour = 0
	} else{
		t.hour++
	}
}

func allBitSeqs(n int) [][]int{
	var bitSeqs [][]int
	numOfBin := int(math.Pow(2,float64(n))) // The total number of binary sequences that will need to be appeneded on to bitSeqs
	fmt.Println(numOfBin)
	for x:=0 ; x<numOfBin ; x++{
		var row []int
		for y:=0 ; y<n ; y++{
			row = append(row, 0)
		}
		bitSeqs = append(bitSeqs, row)
	}
	return bitSeqs
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
	/*
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
	*/
	var valid_time Time24
	var valid_time2 Time24
	var invalid_time Time24
	var smallest_time Time24
	
	valid_time.hour = 12 // All values within the range to be a valid time
	valid_time.minute = 30
	valid_time.second = 45
	
	invalid_time.hour = 25 // All values exceed the allowed values to be a valid time
	invalid_time.minute = 61
	invalid_time.second = 61
	
	valid_time2.hour = 12 // To test equalsTime24()
	valid_time2.minute = 30
	valid_time2.second = 45
	
	smallest_time.hour = 1
	smallest_time.minute = 1
	smallest_time.second = 1
	
	/*
	fmt.Println( equalsTime24(valid_time, valid_time2) )
	fmt.Println( equalsTime24(valid_time, invalid_time) )
	
	fmt.Println( lessThanTime24(valid_time, invalid_time) )
	fmt.Println( lessThanTime24(invalid_time, valid_time) )
	*/
	
	/*
	times := []Time24{valid_time, invalid_time, smallest_time}
		
	var retSmallestTime Time24
	var errMsg error
	
	retSmallestTime, errMsg = minTime24(times)
	
	fmt.Println(retSmallestTime.hour, retSmallestTime.minute, retSmallestTime.second)
	fmt.Println(errMsg)
	fmt.Println(valid_time.String())
	valid_time.AddOneHour()
	fmt.Println(valid_time.String())
	*/
	
	test := [][]int{}
	row1 := []int{1,2,3}
	row2 := []int{4,5,6}
	test = append(test,row1)
	test = append(test,row2)
	
	fmt.Println(allBitSeqs(2))
}