package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	ex1  string = "input/ex1.txt"
	prod string = "input/prod.txt"
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

func findFirstNum(ln string) (fNum string) {
	for i := 0; i < len(ln); i++ {
		if unicode.IsNumber(rune(ln[i])) {
			return string(ln[i])
		}
	}
	return fNum
}

func findLastNum(ln string) (lNum string) {
	for i := len(ln) - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(ln[i])) {
			return string(ln[i])
		}
	}
	return lNum
}

func addNumList(numList []int) (ttl int) {
	for i := 0; i < len(numList); i++ {
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
