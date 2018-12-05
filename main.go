package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	numCases int
	Cases    []cakes
}
type cakes struct {
	pancakes []string
}

func main() {
	tcs, err := getData("trial.csv")
	if err != nil {
		log.Fatalf("error fethching data %v", err)
	}
	for _, tc := range tcs {
		fmt.Println("New Data set\n-----------------------------")
		for i := 0; i < tc.numCases; i++ {
			numFlips := flip(tc.Cases[i].pancakes)
			fmt.Printf("Case# %d: %d\n", i+1, numFlips)
		}
	}
}

func getData(fName string) (tcs []testCase, err error) {
	f, err := os.Open(fName)
	if err != nil {
		return tcs, err
	}
	defer f.Close()
	// each test case is stored in an individual line
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return tcs, err
	}
	for _, line := range lines {
		var tc testCase
		for i, data := range line {
			if i == 0 {
				tc.numCases, err = strconv.Atoi(data)
			} else {
				c := cakes{
					pancakes: strings.SplitAfterN(data, "", 100),
				}
				tc.Cases = append(tc.Cases, c)
			}
		}
		tcs = append(tcs, tc)
	}
	return tcs, nil
}

func flip(pCakes []string) (nFlips int) {
	n := len(pCakes)
	for i := 1; i <= n; i++ {
		if pCakes[n-i] != "+" {
			for j := n - i; j >= 0; j-- {
				if pCakes[j] == "+" {
					pCakes[j] = "-"
				} else {
					pCakes[j] = "+"
				}
			}
			nFlips++
		}
	}
	return nFlips
}
