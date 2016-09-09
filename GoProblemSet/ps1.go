package ps1
import (
	"fmt"
	"strings"
	"errors"
	"io/ioutil"
	"strconv"
	"math"
)

// Question 1

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

// Question 2

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

// Question 3

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
	if ( !validTime24(a) || !validTime24(b) ){
		// Throw run_time error?
	}
	if( (a.hour == b.hour) && (a.minute == b.minute) && (a.second == b.second) ){
		return true
	}
	return false
}

func lessThanTime24(a Time24, b Time24) bool{
	if ( !validTime24(a) || !validTime24(b) ){
		// Throw run_time error?
	}
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
	var zeroTime Time24
	var err error
	
	zeroTime.hour = 0
	zeroTime.minute = 0
	zeroTime.second = 0
	if(len(times) == 0){ // If times has no arguments
		
		err = errors.New("The argument passed in is empty.")
		
		return zeroTime, err
	} else if(len(times) == 1){ // If times has only one argument, then it is the smallest time
		if ( !validTime24(times[0]) ){
			err = errors.New("There exists an invalid time in the slice passed in.")
			return zeroTime, err
		}
		return times[0], noError
	}
	track := 0 // Tracks which of the arguments has the smallest time thus far
	for x := 1; x<len(times);x++{
		if ( !validTime24(times[track]) || !validTime24(times[x]) ){
			err = errors.New("There exists an invalid time in the slice passed in.")
			return zeroTime, err
		}
		if( !lessThanTime24( times[track], times[x] ) ){ // Track only changes if the current tracked time is not strictly less than the one being compared to
			track = x
		}
	}
	return times[track], noError
}

// Question 4

func addZero(t uint8) string{ // Ensures that the method String() for Time24 adds a 0 leading to time values < 10
	if(int(t) < 10){
		return "0" + strconv.Itoa(int(t))
	}
	return strconv.Itoa(int(t))
}

func (t Time24) String() string{
	if ( !validTime24(t) ){
		// Throw run_time error?
	}
	return addZero(t.hour) + ":" + addZero(t.minute) + ":" + addZero(t.second)
}

func (t *Time24) AddOneHour(){
	if ( !validTime24(*t) ){
		// Throw run_time error?
	}
	if(t.hour == 23){
		t.hour = 0
	} else{
		t.hour++
	}
}

// Question 5

func binBitSeq(num int, size int) []int{
	var BitSeq []int
	BinaryVal := string(strconv.FormatInt(int64(num), 2)) // Determines the binary sequence of the value passed in; returned as a string
	numTrailZero := size - len(BinaryVal) // Need to determine how many trailing zeroes we need to add according to the size of the binary sequence
	for x := 0 ; x<numTrailZero ; x++{ // Adding the necessary number of trailing zeroes
		BitSeq = append(BitSeq, 0) 
	}
	for x := 0 ; x<len(BinaryVal) ; x++{ // Appends the string of BinaryVal to BitSeq []int
		var val int
		var err error
		val, err = strconv.Atoi(string(BinaryVal[x]))
		if( err == nil ){
			BitSeq = append( BitSeq, val )
		}
	}
	return BitSeq
}

func allBitSeqs(n int) [][]int{
	var bitSeqs [][]int
	if( n < 1 ){
		return bitSeqs
	}
	numOfBin := int(math.Pow(2,float64(n))) // The total number of binary sequences that will need to be appeneded on to bitSeqs
	for x:=0 ; x<numOfBin ; x++{
		bitSeqs = append(bitSeqs, binBitSeq(x, n))
	}
	return bitSeqs
}