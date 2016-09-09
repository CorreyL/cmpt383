package ps1

import(
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
	var final string
	final = ""
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
	
	var expected string
	expected = "{\"The\":1, \"big\":3, \"dog\":1, \"ate\":1, \"the\":1, \"apple\":1}"
	
	if(final != expected){
		t.Error(
			"For textfile.txt",
			"Got ", final,
			"Expected ", expected,
		)
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
	if( !equalsTime24(times[0], times[1]) ){ // Equal times passed in
		t.Error(
			"For ", times[0], "and ", times[1],
			"Expected True",
			"Got False",
		)
	}
	if( equalsTime24(times[0], times[3]) ){ // Unequal times passed in
		t.Error(
			"For ", times[0], "and ", times[3],
			"Expected True",
			"Got False",
		)
	}
	
	// Testing lessThanTime24()
	
}

