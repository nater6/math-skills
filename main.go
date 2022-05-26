package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	//Get the filename from the terminal
	cmd := os.Args
	if len(cmd) != 2 {
		//If an invalid command is put in print a message to the terminal
		fmt.Println("Please enter a valid command: \"go run program-name.go filename.txt\"")
		return
	}
	stats := cmd[1]

	// Open the file
	file, err := os.Open(stats)
	if err != nil {
		fmt.Println("This file couldn't be found, please use a valid filename")
		log.Fatalf("ERROR: %s \n", err)
	}

	//scan through the file and add all the values to a slice
	numbers := []float64{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		currentNum, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid data type in the file")
			log.Fatalf("ERROR: %s \n", err)
		}
		numbers = append(numbers, float64(currentNum))
	}

	average := average(numbers)
	median := median(numbers)
	stdDev := stdDev(numbers)
	variance := variance(numbers)
	fmt.Println("Average: ", average)
	fmt.Println("Median: ", median)
	fmt.Println("Variance: ", variance)
	fmt.Println("Standard Deviation: ", stdDev)

}

func median(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)

	sort.Float64s(dataCopy)

	var median float64
	l := len(dataCopy)

	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[(l/2)-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}
	return median
}

func average(data []float64) float64 {
	//Calculate the average (Add up all numbers and divide by the amount of numbers
	var average float64

	for _, r := range data {
		average += r
	}

	average /= float64(len(data))

	return average
}

func stdDev(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	var sum, mean, sd float64

	for i, r := range data {
		dataCopy[i] = r
		sum += r
	}
	mean = sum / float64(len(data))

	for j := 0; j < len(dataCopy); j++ {
		sd += math.Pow(dataCopy[j]-mean, 2)
	}
	sd = math.Sqrt(sd / float64(len(dataCopy)))
	return sd
}

func variance(data []float64) float64 {
	//find the mean of the values
	var mean float64

	for _, r := range data {
		mean += r
	}
	mean /= float64(len(data))

	//fid the difference between each value and the mean
	sqrs := make([]float64, len(data))
	for i, r := range data {
		sqrs[i] = r - mean
	}

	//square the difference between the means
	for i, r := range sqrs {
		sqrs[i] = r * r
	}

	//find the sum of all the squares
	var sqrSums float64
	for _, r := range sqrs {
		sqrSums += r
	}
	fmt.Println("SUM OF THE SQR NUMBERS: ", sqrSums)

	//divide the sum of the sqrs by the amount of numbers
	sqrSums /= float64(len(data))
	fmt.Println("BEFORE CONVERTED TO INT: ", sqrSums)

	return sqrSums
}
