package ps1

import(
	// "fmt"
	"testing"
	"strconv"
)

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
	for _, pair := range tests {
		v := countPrimes(pair.value)
		if v != pair.expected {
			t.Error(
				"For", pair.value,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

// Question 2
func TestCountstrings(t *testing.T){
	var final string = ""
	final = final + "{"
	len_check := 0 // Ensures that the string ", " is not printed for the final pair
	m := countStrings("textfile.txt")
	for x, y := range m { 
		final = final + "\"" + x + "\"" + ":" + strconv.Itoa(y)
		if(len_check != len(m)-1){
			final = final + ", "
		}
		len_check++
	}
	final = final + "}"
	
	var expected string = "{\"The\":1, \"big\":3, \"dog\":1, \"ate\":1, \"the\":1, \"apple\":1}"
	
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
	m = countStrings("raspberryBeret.txt")
	for x, y := range m { 
		final = final + "\"" + x + "\"" + ":" + strconv.Itoa(y)
		if(len_check != len(m)-1){
			final = final + ", "
		}
		len_check++
	}
	final = final + "}"
	expected = "{\"The\":1, \"second\":1, \"much\":1, \"think\":1, \"she\":1, \"She\":1, \"And\":1, \"warm\":1, \"her\":1, \"you\":1, \"wouldn't\":1, \"wear\":1, \"love\":1, \"wore\":1, \"in\":1, \"hand\":1, \"beret\":3, \"find\":1, \"more\":1, \"kind\":1, \"store\":1, \"if\":1, \"a\":2, \"Raspberry\":3, \"it\":1, \"was\":1, \"I\":2}" //'
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
	if( !validTime24(times[0]) ) { // Valid time passed in
		t.Error(
			"For ", times[0],
			"Got False",
			"Expected True",
		)
	}
	
	if( validTime24(times[2]) ){ // Invalid time passed in
		t.Error(
			"For ", times[2],
			"Got True",
			"Expected False",
		)
	}
	
	// Testing equalsTime24()
	if result, _ := equalsTime24(times[0], times[1]); !result { // Equal times passed in
		t.Error(
			"For ", times[0], "and ", times[1],
			"Expected True",
			"Got False",
		)
	}
	if result, _ := equalsTime24(times[0], times[3]); result { // Unequal times passed in
		t.Error(
			"For ", times[0], "and ", times[3],
			"Expected True",
			"Got False",
		)
	}
	
	// Testing lessThanTime24()
	if result, _ := lessThanTime24(times[3], times[0]); !result{
		t.Error(
			"For lessThanTime24(times[0], times[3])",
			"Expected False",
			"Got True",
		)
	}
	
	if result, _ := lessThanTime24(times[0], times[3]); result {
		t.Error(
			"For lessThanTime24(times[3], times[0])",
			"Expected True",
			"Got False",
		)
	}
	
	// Testing minTime24
	oneTime := Time24{hour: 1, minute: 1, second: 1}
	var retVal Time24
	retVal, _ = minTime24(times_valid) // Expecting to return 01:01:01
	// Note that minTime24 requires that all times in the slice be valid times, thus a separate slice needed to be created
	if ( retVal != oneTime ){
		t.Error(
			"For ", times,
			"Expected ", times[3],
			"Got ", retVal,
		)
	}
	
	var emptyTimes []Time24 // Testing to see if an empty slice of Time24 returns 00:00:00
	var zeroTime = Time24{hour: 0, minute: 0, second: 0}
	retVal, _ = minTime24(emptyTimes)
	if( retVal != zeroTime ){
		t.Error(
			"For ", emptyTimes,
			"Expected ", zeroTime,
			"Got ", retVal,
		)
	}
}

	
// Question 4
func TestMethods(t *testing.T){
	// Testing String()
	var expected string = "05:39:08"
	time := Time24{hour: 5, minute: 39, second: 8}
	if ( time.String() != expected ){
		t.Error(
			"For ", time,
			"Expected ", expected,
			"Got ", time.String(),
		)
	}
	
	expected = "23:59:59"
	time = Time24{hour: 23, minute: 59, second: 59}
		if ( time.String() != expected ){
		t.Error(
			"For ", time,
			"Expected ", expected,
			"Got ", time.String(),
		)
	}
	
	//Testing AddOneHour()
	AOH_time := Time24{hour: 20, minute: 15, second: 0}
	for x := 21; x < 25; x++ {
		AOH_time.AddOneHour()
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
}

// Question 5

func TestAllbitseqs(t *testing.T){
	var expected = [][][]int{
		{ {0},{1} },
		{ {0,0}, {0,1}, {1,0}, {1,1} },
		{ {0,0,0}, {0,0,1}, {0,1,0}, {0,1,1}, {1,0,0}, {1,0,1}, {1,1,0}, {1,1,1} },
		{ {0,0,0,0}, {0,0,0,1}, {0,0,1,0}, {0,0,1,1}, {0,1,0,0}, {0,1,0,1}, {0,1,1,0}, {0,1,1,1}, {1,0,0,0}, {1,0,0,1}, {1,0,1,0}, {1,0,1,1}, {1,1,0,0}, {1,1,0,1}, {1,1,1,0}, {1,1,1,1} },
	}
	for x:=0; x<len(expected); x++{
		retVal := allBitSeqs(x+1)
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