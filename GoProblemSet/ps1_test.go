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
	
}

