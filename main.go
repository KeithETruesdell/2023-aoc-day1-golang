package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ex1  string = "input/ex1.txt"
	prod string = "input/prod.txt"
	ex2  string = "input/ex2.txt"
)

var (
	digitList []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func main() {
	var (
		ans int
		err error
	)

	f := getFlags()
	f2rStr := getFile2Read(f)
	fData, err := getFileData(f2rStr)
	if err != nil {
		fmt.Printf("ERROR reading file:\n    %v", err)
	}

	ans, err = processData(fData)
	if err != nil {
		fmt.Printf("ERROR processing file data:\n    %v", err)
	}

	fmt.Println("=====================")
	fmt.Println(ans)
}

// processData -
func processData(fData []string) (ans int, err error) {
	var (
		numList []int
	)
	for i := 0; i < len(fData); i++ {

		tmpLn := fData[i]
		fNum := findFirstNum(tmpLn)
		lNum := findLastNum(tmpLn)
		cNumStr := fNum + lNum

		comboNumInt, err := strconv.Atoi(cNumStr)
		if err != nil {
			return ans, err
		}

		numList = append(numList, comboNumInt)
	}

	ans = addNumList(numList)

	return ans, err
}

func convDigitStr(digitStr string) (dStr string) {
	if len(digitStr) == 1 {
		return digitStr
	}
	switch digitStr {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	fmt.Printf("  -- ERROR: Did not find dStr ( %v ) --  \n", digitStr)
	return dStr
}

func findFirstNum(ln string) (fNum string) {
	var (
		fNumIdx int = len(ln) + 1
	)

	for d := 0; d < len(digitList); d++ {
		tmpIdx := strings.Index(ln, digitList[d])
		if tmpIdx != -1 {
			if tmpIdx < fNumIdx {
				fNumIdx = tmpIdx
				fNum = digitList[d]
			}
		}
	}

	return convDigitStr(fNum)
}

func findLastNum(ln string) (lNum string) {
	var (
		lNumIdx int = -1
	)

	for d := 0; d < len(digitList); d++ {
		tmpIdx := strings.LastIndex(ln, digitList[d])
		if tmpIdx != -1 {
			if tmpIdx > lNumIdx {
				lNumIdx = tmpIdx
				lNum = digitList[d]
			}
		}
	}

	if len(lNum) < 1 {
		fmt.Println(ln)
		fmt.Println(lNumIdx)
		fmt.Println(lNum)
		fmt.Println("something wrong happend to lNum ...")
	}
	return convDigitStr(lNum)
}

func addNumList(numList []int) (ttl int) {
	for i := 0; i < len(numList); i++ {
		//fmt.Println(numList[i])
		ttl = ttl + numList[i]
	}

	return ttl
}

// getFlags - get the flags from the command line
func getFlags() (f int) {
	fTmp := flag.Int("f", 1, "Which example file to use?  1 is default, 0 is 'live'")
	flag.Parse()
	f = *fTmp
	return f
}

// getFile2Read - based on the flag get the string of the file to read
func getFile2Read(f int) (f2r string) {
	switch f {
	case 0:
		f2r = prod
	case 1:
		f2r = ex1
	case 2:
		f2r = ex2
	default:
		f2r = ex1
	}
	return f2r
}

// getFileData - read the data from the file and retun some sort of data structure...
func getFileData(f2r string) (fData []string, err error) {
	inputF, err := os.Open(f2r)
	if err != nil {
		return fData, err
	}
	defer inputF.Close()

	scn := *bufio.NewScanner(inputF)

	for scn.Scan() {
		ln := strings.TrimSpace(scn.Text())
		fData = append(fData, ln)
	}

	return fData, err
}
