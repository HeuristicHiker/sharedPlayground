package main

/* I've got my key value pair
What about IV
	need to know if I have to subtract
Between 5 and 10 do I subtract?
	5 no - it's a key
	6 no - I add
	7 no - I add
	8 no - I add
	9 yes - I sub
Between 50 and 100?
	41 - XLI
	67 - LXVII
	71 - LXXI
	77 - LXXVII
	78 - LXXVIII
	79 - LXXIX
*/

func main() {
	// input := "MCMXCIV"
	// romanToInt(input)

}

func romanToInt(s string) int {
	total := 0
	largestValue := 0

	for _, numeral := range s {
		value := romanNumeralToInt(string(numeral))

		// is subtracting
		if value > largestValue {
			// subtract * 2 because it was already added
			total += value - 2*largestValue
		} else {
			total += value
		}

		largestValue = value

	}
	return total
}

func romanNumeralToInt(rn string) int {
	romanToInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	return romanToInt[rn]
}

type numInt struct {
	value   int
	numeral string
}

func intToRomanNumeral(i int) string {
	soln := ""
	k := 0
	numerals := []string{"I", "V", "X", "L", "C", "D", "M"}

	for i > 0 {
		d := i % 10
		v := ""
		if d < 4 {
			for j := 0; j < d; j++ {
				v += numerals[k]
			}
		} else if d == 4 {
			v += numerals[k] + numerals[k+1]
		} else if d < 9 {
			v += numerals[k+1]
			for j := 0; j < d-5; j++ {
				v += numerals[k]
			}
		} else if d == 9 {
			v += numerals[k] + numerals[k+2]
		}
		soln = v + soln
		k += 2
		i /= 10
	}

	return soln
}

func intToRomanNumeralsNotWorking(i int) string {
	numInts := []numInt{
		{1000, "M"},
		{500, "D"},
		{100, "C"},
		{50, "L"},
		{10, "X"},
		{5, "V"},
		{1, "I"},
	}

	soln := ""

	for idx, irn := range numInts {
		if i == 0 {
			break
		}

		if idx < len(numInts)-1 {
			next := numInts[idx+1]

			if i >= irn.value-next.value && i < irn.value && (irn.value-next.value) != next.value {
				soln += next.numeral + irn.numeral
				i -= irn.value - next.value
				continue
			}
		}

		for i >= irn.value {
			var lastNum numInt
			if idx > 0 {
				lastNum = numInts[idx-1]
			} else {
				lastNum = numInt{0, ""}
			}
			var nextNum numInt
			if idx < len(numInts)-1 {
				nextNum = numInts[idx+1]
			} else {
				nextNum = numInt{0, ""}
			}
			// fmt.Println(" irn ", irn.numeral, " lastVal: ", lastNum.numeral, "i is ", i, soln, " next: ", nextNum.numeral)
			if i == lastNum.value-nextNum.value {
				soln = soln + nextNum.numeral + lastNum.numeral
				i = 0
			} else {
				soln += irn.numeral
				i -= irn.value
			}
			// fmt.Println("--- soln: ", soln, " i: ", i)
		}
	}

	return soln
}

func digitToNumeral(i int) string {
	// numInts := []numInt{
	// 	{1, "I"},
	// 	{5, "V"},
	// 	{10, "X"},
	// 	{50, "L"},
	// 	{100, "C"},
	// 	{500, "D"},
	// 	{1000, "M"},
	// }

	// diff := 1

	// fmt.Println(i)

	// for _, ni := range numInts {
	// 	val := ni.value
	// 	s := ni.numeral

	// }

	return "Hi"

}
