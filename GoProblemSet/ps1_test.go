package ps1

import(
	"fmt"
	"testing"
	"strconv"
)

var ToConsole bool = true

// Question 1
func TestCountprimes(t *testing.T){
	type testpair struct{
		value int
		expected int
	}
	var tests = []testpair{
		{4, 2},
		{10000, 1229},
		{-6, 0},
		// Own test cases
		{5, 3},
		{7, 4},
		{11, 5},
		{13, 6},
	}
	if(ToConsole){
		fmt.Println("Begin Testing: countPrimes(x,y)")
	}
	for _, pair := range tests {
		v := countPrimes(pair.value)
		if(ToConsole){
			fmt.Println("	Testing:", pair.value)
			fmt.Println("	Expecting:", pair.expected)
			fmt.Println("	countPrimes() returned:", v)
		}
		if v != pair.expected {
			t.Error(
				"For", pair.value,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
	if(ToConsole){
		fmt.Println("\n")
	}
}

// Question 2
func TestCountstrings(t *testing.T){
	if(ToConsole){
		fmt.Println("Begin Testing: countStrings()")
	}
	var final string = ""
	final = final + "{"
	len_check := 0 // Ensures that the string ", " is not printed for the final pair
	m, _ := countStrings("textfile.txt")
	for x, y := range m { 
		final = final + "\"" + x + "\"" + ":" + strconv.Itoa(y)
		if(len_check != len(m)-1){
			final = final + ", "
		}
		len_check++
	}
	final = final + "}"
	
	var expected string = "{\"The\":1, \"big\":3, \"dog\":1, \"ate\":1, \"the\":1, \"apple\":1}"
	if(ToConsole){
		fmt.Println("	Testing: textfile.txt")
		fmt.Println("	Expecting:", expected)
		fmt.Println("	countStrings() returned:", final)
	}
	if(final != expected){
		t.Error(
			"For textfile.txt",
			"Got ", final,
			"Expected ", expected,
		)
	}
	final = ""
	final = final + "{"
	len_check = 0 // Ensures that the string ", " is not printed for the final pair
	m, _ = countStrings("raspberryBeret.txt")
	for x, y := range m { 
		final = final + "\"" + x + "\"" + ":" + strconv.Itoa(y)
		if(len_check != len(m)-1){
			final = final + ", "
		}
		len_check++
	}
	final = final + "}"
	expected2 := map[string]int{
		"The":1,
		"second":1,
		"much":1,
		"think":1,
		"she":1,
		"She":1,
		"And":1,
		"warm":1,
		"her":1,
		"you":1,
		"wouldn't":1,
		"wear":1,
		"love":1,
		"wore":1,
		"in":1,
		"hand":1,
		"it":1,
		"was":1,
		"I":2,
		"beret":3,
		"if":1,
		"find":1,
		"more":1,
		"store":1,
		"kind":1,
		"a":2,
		"Raspberry":3,
	}
	fmt.Println("\nFor raspberryBeret.txt")
	for x, y := range m{
		if( ToConsole ){
			fmt.Println("	Testing: ", x)
			fmt.Println("	Expecting:", expected2[x])
			fmt.Println("	countStrings() returned:", y)
		}
		if(y != expected2[x]){
			t.Error(
				"For ", x,
				"Got ", y,
				"Expected ", expected2[x],
			)
		}
	}
	
	if( ToConsole ){
		fmt.Println("\n")
	}
}

// Question 3
func TestTime24(t *testing.T){
	var times = []Time24{
		{12, 30, 45},
		{12, 30, 45},
		{25, 61, 61},
		{1,1,1},
	}
	
	var times_valid = []Time24{ // Note that minTime24 requires that all times in the slice be valid times, thus a separate slice needed to be created
		{12, 30, 45},
		{12, 30, 45},
		{1,1,1},
	}
	
	// Testing validTime24()
	if(ToConsole){
		fmt.Println("Begin Testing: validTime24()")
	}
	var funcRetVal = validTime24(times[0]);
	if(ToConsole){
		fmt.Println("	Testing:", times[0])
		fmt.Println("	Expecting:", true)
		fmt.Println("	countStrings() returned:", funcRetVal)
	}
	if( !funcRetVal ) { // Valid time passed in
		t.Error(
			"For ", times[0],
			"Got False",
			"Expected True",
		)
	}
	funcRetVal = validTime24(times[2])
	if(ToConsole){
		fmt.Println("	Testing:", times[2])
		fmt.Println("	Expecting:", false)
		fmt.Println("	countStrings() returned:", funcRetVal)
	}
	if( funcRetVal ){ // Invalid time passed in
		t.Error(
			"For ", times[2],
			"Got True",
			"Expected False",
		)
	}
	
	// Testing equalsTime24()
	funcRetVal = equalsTime24(times[0], times[1])
	if(ToConsole){
		fmt.Println("Begin Testing: equalsTime24(a Time24, b Time24)")
		fmt.Println("	Testing:", times[0], "&", times[1])
		fmt.Println("	Expecting:", true)
		fmt.Println("	countStrings() returned:", funcRetVal)
	}
	if( !funcRetVal ){ // Equal times passed in
		t.Error(
			"For ", times[0], "and ", times[1],
			"Expected True",
			"Got False",
		)
	}
	funcRetVal = equalsTime24(times[0], times[3]) // Unequal times passed in
	if(ToConsole){
		fmt.Println("	Testing:", times[0], "&", times[3])
		fmt.Println("	Expecting:", false)
		fmt.Println("	countStrings() returned:", funcRetVal)
	}
	if( funcRetVal ){
		t.Error(
			"For ", times[0], "and ", times[3],
			"Expected True",
			"Got False",
		)
	}
	
	// Testing lessThanTime24()
	funcRetVal = lessThanTime24(times[3], times[0])
	if(ToConsole){
		fmt.Println("Begin Testing: lessThanTime24(a Time24, b Time24)")
		fmt.Println("	Testing:", times[3], "&", times[0])
		fmt.Println("	Expecting:", true)
		fmt.Println("	countStrings() returned:", funcRetVal)
	}
	if( !funcRetVal ){
		t.Error(
			"For lessThanTime24(times[0], times[3])",
			"Expected False",
			"Got True",
		)
	}
	funcRetVal = lessThanTime24(times[0], times[3])
	if(ToConsole){
		fmt.Println("	Testing:", times[0], "&", times[3])
		fmt.Println("	Expecting:", false)
		fmt.Println("	countStrings() returned:", funcRetVal)
	}
	if( funcRetVal ){
		t.Error(
			"For lessThanTime24(times[3], times[0])",
			"Expected True",
			"Got False",
		)
	}
	
	// Testing minTime24
	oneTime := Time24{hour: 1, minute: 1, second: 1}
	funcRetVal2, _ := minTime24(times_valid) // Expecting to return 01:01:01
	if(ToConsole){
		fmt.Println("Begin Testing: minTime24()")
		fmt.Println("	Testing:", times_valid)
		fmt.Println("	Expecting: 01:01:01", )
		fmt.Println("	countStrings() returned:", funcRetVal2)
	}
	// Note that minTime24 requires that all times in the slice be valid times, thus a separate slice needed to be created
	if ( funcRetVal2 != oneTime ){
		t.Error(
			"For ", times,
			"Expected ", times[3],
			"Got ", funcRetVal2,
		)
	}
	
	var emptyTimes []Time24 // Testing to see if an empty slice of Time24 returns 00:00:00
	var zeroTime = Time24{hour: 0, minute: 0, second: 0}
	funcRetVal2, _ = minTime24(emptyTimes)
	if(ToConsole){
		fmt.Println("	Testing:", emptyTimes, "(Ie. An empty array of Time24)")
		fmt.Println("	Expecting: 00:00:00", )
		fmt.Println("	countStrings() returned:", funcRetVal2)
	}
	if( funcRetVal2 != zeroTime ){
		t.Error(
			"For ", emptyTimes,
			"Expected ", zeroTime,
			"Got ", funcRetVal2,
		)
	}
	if( ToConsole ){
		fmt.Println("\n")
	}
}

	
// Question 4
func TestMethods(t *testing.T){
	// Testing String()
	var expected string = "05:39:08"
	time := Time24{hour: 5, minute: 39, second: 8}
	if(ToConsole){
		fmt.Println("Begin Testing: Time24.String()")
		fmt.Println("	Testing: time = 05:39:08; time.String()")
		fmt.Println("	Expecting: 05:39:08", )
		fmt.Println("	time.String() returned:", time.String())
	}
	if ( time.String() != expected ){
		t.Error(
			"For ", time,
			"Expected ", expected,
			"Got ", time.String(),
		)
	}
	
	expected = "23:59:59"
	time = Time24{hour: 23, minute: 59, second: 59}
	if(ToConsole){
		fmt.Println("	Testing: time = 23:59:59; time.String()")
		fmt.Println("	Expecting: 23:59:59", )
		fmt.Println("	time.String() returned:", time.String())
	}
	if ( time.String() != expected ){
		t.Error(
			"For ", time,
			"Expected ", expected,
			"Got ", time.String(),
		)
	}
	
	//Testing AddOneHour()
	AOH_time := Time24{hour: 20, minute: 15, second: 0}
	if( ToConsole ){
		fmt.Println("Begin Testing: Time24.AddOneHour()")
	}
	for x := 21; x < 25; x++ {
		AOH_time.AddOneHour()
		if(ToConsole){
			fmt.Println("	Testing:", AOH_time.String())
			if(x < 25){
				fmt.Println("	Expecting:", strconv.Itoa(x)+":"+strconv.Itoa(int(AOH_time.minute))+":0"+strconv.Itoa(int(AOH_time.second)) )
			}else{
				fmt.Println("	Expecting: 0:"+strconv.Itoa(int(AOH_time.minute))+":"+strconv.Itoa(int(AOH_time.second)))
			}
			fmt.Println("	AOH_time.AddOneHour returned:", AOH_time)
		}
		if( x < 24 && int(AOH_time.hour) != x ){
			t.Error(
				"For ", AOH_time,
				"Expected ", x,
				"Got ", AOH_time.hour,
			)
		} else if( x == 24 && int(AOH_time.hour) != 0 ) {
			t.Error(
				"For ", AOH_time,
				"Expected ", 0,
				"Got ", AOH_time.hour,
			)
		}
	}
	if( ToConsole ){
		fmt.Println("\n")
	}
}

// Question 5

func TestAllbitseqs(t *testing.T){
	var expected = [][][]int{
		{ {0},{1} },
		{ {0,0}, {0,1}, {1,0}, {1,1} },
		{ {0,0,0}, {0,0,1}, {0,1,0}, {0,1,1}, {1,0,0}, {1,0,1}, {1,1,0}, {1,1,1} },
		{ {0,0,0,0}, {0,0,0,1}, {0,0,1,0}, {0,0,1,1}, {0,1,0,0}, {0,1,0,1}, {0,1,1,0}, {0,1,1,1}, {1,0,0,0}, {1,0,0,1}, {1,0,1,0}, {1,0,1,1}, {1,1,0,0}, {1,1,0,1}, {1,1,1,0}, {1,1,1,1} },
	}
	if( ToConsole ){
		fmt.Println("Begin Testing: allBitSeqs()")
	}
	for x:=0; x<len(expected); x++{
		retVal := allBitSeqs(x+1)
		if( ToConsole ){
			fmt.Println("	Expecting:", expected[x])
			fmt.Println("	allBitSeqs("+strconv.Itoa(x+1)+") returned:", retVal)
		}
		for y:=0; y<len(retVal);y++ {
			for z:=0; z<len(retVal[0]); z++{
				if( retVal[y][z] != expected[x][y][z] ){
					t.Error(
						"For allBitSeqs(", x, ")","\n",
						"Expected \n", expected[x], "\n",
						"Got \n", retVal,
					)
				}
			}
		}
	}
}