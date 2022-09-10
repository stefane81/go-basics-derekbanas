package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	// sectionString := ""
	section := 1

	fmt.Println("Select your section")
	fmt.Println("(1): Input")
	fmt.Println("(2): Variables")
	fmt.Println("(3): Strings")
	fmt.Println("(4): Runes")
	fmt.Println("(5): Time master")
	fmt.Println("(6): Basic For loops")
	fmt.Println("(7): Slices")
	fmt.Println("(8): Function Errors")
	fmt.Println("(9): Basic pointers")

	fmt.Println("What section do you want to visit?")
	fmt.Scan(&section)
	// reader := bufio.NewReader(os.Stdin)
	// sectionString, err := reader.ReadString('\n')

	// section, err := strconv.Atoi(sectionString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	switch section {
	case 1:
		{
			Input()
		}
	case 2:
		{
			Variables()
		}
	case 3:
		{
			Strings()
		}
	case 4:
		{
			Runes()
		}
	case 5:
		{
			Time()
		}
	case 6:
		{
			BasicForLoops()
		}
	case 7:
		{
			Slices()
		}
	case 8:
		{
			FunctionErrors()
		}
	case 9:
		{
			BasicPointers()
		}
	case 10:
		{
			FileIO()
		}
	}

}

// case 10 FileIO
func FileIO() {

}

// case 9 BasicPointers
func changeVal(myPrt *int) {
	*myPrt = 12
}
func BasicPointers() {
	f3 := 10
	fmt.Println("f3 before func:", f3)
	changeVal(&f3)
	fmt.Println("f3 after func:", f3)
}

// case 8 FunctionErrors
func FunctionErrors() {
	fmt.Println(getQuo(5, 4))
}

func getQuo(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You cant divide by zero")
	} else {
		return x / y, nil
	}
}

// case 7 BasicForLoops
func BasicForLoops() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	aNums := []int{1, 2, 3, 4, 5}
	for _, num := range aNums {
		fmt.Println(num)
	}
}

// case 6 Slices
func Slices() {
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "simulated"
	sl1[4] = "Bergens"
	fmt.Println("Slice Size: ", len(sl1))
	for i := 0; i < len(sl1); i++ {
		fmt.Println(sl1[i])
	}

	// foreach
	for _, x := range sl1 {
		fmt.Println(x)
	}
}

// case1 input
func Input() {
	fmt.Println("***Input section***")
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		fmt.Println("Hello", name)
	} else {
		log.Fatal(err)
	}
}

// case 2 Variables
func Variables() {
	fmt.Println("***Variable section***")
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	fmt.Println(cV4, err, reflect.TypeOf(cV4))
}

// case 3 Strings
func Strings() {
	fmt.Println("***String section***")
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	fmt.Println(sV2)
	fmt.Println("Length:", len(sV2))
	fmt.Println("Contains Another:", strings.Contains(sV2, "Another"))
	fmt.Println("o index:", strings.Index(sV2, "o"))
	fmt.Println("Replace:", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3)
	fmt.Println("Split:", strings.Split("a-b-c-d", "-"))
}

// case 4 Runes
func Runes() {
	rStr := "abcdefg"
	fmt.Println("Rune Count:", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}

// case 5 Time
func Time() {
	tNow := time.Now()
	fmt.Println(tNow.Year(), tNow.Month(), tNow.Day())
	fmt.Printf("%d:%d:%d\n", tNow.Hour(), tNow.Minute(), tNow.Second())
}
