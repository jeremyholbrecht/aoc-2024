package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("error reading file... ", err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	var list1 = []int{}
	var list2 = []int{}
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			// split into substring based on whitespace
			fields := strings.Fields(string(line))

			// convert string data to int
			fieldToInt1, err := strconv.Atoi(fields[0])
			if err != nil {
				log.Fatal(err)
			}

			fieldToInt2, err := strconv.Atoi(fields[1])
			if err != nil {
				log.Fatal(err)
			}

			// add data to slice list1 & list2
			list1 = append(list1, fieldToInt1)
			list2 = append(list2, fieldToInt2)

		}
		if err != nil {
			break
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	// using math.Abs() which requires a float resulted in result displaying a not so usuable result (3.508942e+06)
	// using int and writing your own absolute function solves this
	/*
		var result float64
		for i := 0; i < len(list1); i++ {
			result += (math.Abs(float64(list1[i]) - float64(list2[i])))
		}

	*/

	var result int

	for i := 0; i < len(list1); i++ {
		result += absolute(list1[i] - list2[i])

	}

	fmt.Println(result)

}

// return a positive number if number is negative
func absolute(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
